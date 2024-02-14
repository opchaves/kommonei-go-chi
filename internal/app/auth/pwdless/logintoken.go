package pwdless

import (
	"crypto/rand"
	"errors"
	"sync"
	"time"
)

var (
	errTokenNotFound = errors.New("login token not found")
)

// LoginToken is an in-memory saved token referencing an account ID and an expiry date.
type LoginToken struct {
	Token  string
	UserID string
	Expiry time.Time
}

// LoginTokenAuth implements passwordless login authentication flow using temporary in-memory stored tokens.
type LoginTokenAuth struct {
	token            map[string]LoginToken
	mux              sync.RWMutex
	loginURL         string
	loginTokenLength int
	loginTokenExpiry time.Duration
}

// NewLoginTokenAuth configures and returns a LoginToken authentication instance.
func NewLoginTokenAuth() (*LoginTokenAuth, error) {
	a := &LoginTokenAuth{
		token:            make(map[string]LoginToken),
		loginURL:         "http://localhost:3000/login",
		loginTokenLength: 8,
		loginTokenExpiry: 11 * time.Minute,
	}
	return a, nil
}

// CreateToken creates an in-memory login token referencing account ID. It returns a token containing a random tokenstring and expiry date.
func (a *LoginTokenAuth) CreateToken(id string) LoginToken {
	lt := LoginToken{
		Token:  randStringBytes(a.loginTokenLength),
		UserID: id,
		Expiry: time.Now().Add(a.loginTokenExpiry),
	}
	a.add(lt)
	a.purgeExpired()
	return lt
}

// GetAccountID looks up the token by tokenstring and returns the account ID or error if token not found or expired.
func (a *LoginTokenAuth) GetAccountID(token string) (string, error) {
	lt, exists := a.get(token)
	if !exists || time.Now().After(lt.Expiry) {
		return "", errTokenNotFound
	}
	a.delete(lt.Token)
	return lt.UserID, nil
}

func (a *LoginTokenAuth) get(token string) (LoginToken, bool) {
	a.mux.RLock()
	lt, ok := a.token[token]
	a.mux.RUnlock()
	return lt, ok
}

func (a *LoginTokenAuth) add(lt LoginToken) {
	a.mux.Lock()
	a.token[lt.Token] = lt
	a.mux.Unlock()
}

func (a *LoginTokenAuth) delete(token string) {
	a.mux.Lock()
	delete(a.token, token)
	a.mux.Unlock()
}

func (a *LoginTokenAuth) purgeExpired() {
	for t, v := range a.token {
		if time.Now().After(v.Expiry) {
			a.delete(t)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randStringBytes(n int) string {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	for k, v := range buf {
		buf[k] = letterBytes[v%byte(len(letterBytes))]
	}
	return string(buf)
}
