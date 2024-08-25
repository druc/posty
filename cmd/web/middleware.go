package main

import (
	"context"
	"net/http"

	"github.com/druc/posty/internal/models"
)

type contextKey string

const contextKeyUser = contextKey("user")

func (app *app) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.session.Get(r, "posty")
		userId, ok := session.Values["userId"].(int)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		user, err := app.users.Find(userId)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *app) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(contextKeyUser).(models.User)
		if !ok {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
