package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

const (
	stateCookieName  = "oauthstate"
	userIDCookieName = "userID"
)

func setOauthStateCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: stateCookieName, Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func setUserIDCookie(w http.ResponseWriter, userID string) {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     userIDCookieName,
		Value:    userID,
		Expires:  expiration,
		HttpOnly: false, // needs access from main.js
		Path:     "/",
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(w, &cookie)
}
