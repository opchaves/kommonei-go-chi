package pwdless

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mssola/useragent"
	"github.com/opchaves/go-kom/config"
	"github.com/opchaves/go-kom/model"
	"github.com/opchaves/go-kom/server/auth/jwt"
)

type tokenRequest struct {
	Token string `json:"token"`
}

type tokenResponse struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (body *tokenRequest) Bind(r *http.Request) error {
	body.Token = strings.TrimSpace(body.Token)

	return validation.ValidateStruct(body,
		validation.Field(&body.Token, validation.Required, is.Alphanumeric),
	)
}

func (rs *Resource) token(w http.ResponseWriter, r *http.Request) {
	oplog := log(r)

	// TODO: regex for token in path
	token := chi.URLParam(r, "token")

	if token == "" {
		render.Render(w, r, ErrUnauthorized(ErrLoginToken))
		return
	}

	sId, err := rs.LoginAuth.GetAccountID(token)
	if err != nil {
		render.Render(w, r, ErrUnauthorized(ErrLoginToken))
		return
	}

	id, err := uuid.Parse(sId)

	// user, err := rs.Q.GetUserById(r.Context(), id)
	user, err := rs.Q.GetDefaultUserWorkspace(r.Context(), id)
	if err != nil {
		// account deleted before login token expired
		render.Render(w, r, ErrUnauthorized(ErrUnknownLogin))
		return
	}

	// TODO: check if user is active
	// if !user.CanLogin() {
	// 	render.Render(w, r, ErrUnauthorized(ErrLoginDisabled))
	// 	return
	// }

	ua := useragent.New(r.UserAgent())
	browser, _ := ua.Browser()

	var expiresAt pgtype.Timestamp
	expiresAt.Scan(time.Now().Add(config.JwtRefreshExpiry))
	identifier := fmt.Sprintf("%s on %s", browser, ua.OS())

	aToken := model.CreateTokenParams{
		Token:      uuid.Must(uuid.NewRandom()).String(),
		ExpiresAt:  expiresAt,
		Mobile:     ua.Mobile(),
		UserID:     user.ID,
		Identifier: &identifier,
	}

	newToken, err := rs.Q.CreateToken(r.Context(), aToken)
	if err != nil {
		oplog.With("user", user.ID).Error(err.Error())
		render.Render(w, r, ErrInternalServerError)
		return
	}

	claims := jwt.AppClaims{
		ID:    user.ID.String(),
		OrgID: user.WorkspaceID.String(),
		Sub:   user.FirstName,
		Roles: []string{"user"},
	}

	refreshClaims := jwt.RefreshClaims{
		ID:    int(newToken.ID),
		Token: newToken.Token,
	}

	access, refresh, err := rs.TokenAuth.GenTokenPair(claims, refreshClaims)
	if err != nil {
		oplog.With("user", user.ID).Error(err.Error())
		render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Respond(w, r, &tokenResponse{
		Access:  access,
		Refresh: refresh,
	})
}
