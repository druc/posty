package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.getHome)
	mux.Handle("GET /posts/create", app.requireAuth(http.HandlerFunc(app.createPost)))
	mux.Handle("POST /posts/create", app.requireAuth(http.HandlerFunc(app.storePost)))
	mux.HandleFunc("GET /register", app.getRegister)
	mux.HandleFunc("POST /register", app.storeRegister)
	mux.HandleFunc("GET /login", app.getLogin)
	mux.HandleFunc("POST /login", app.storeLogin)
	mux.HandleFunc("POST /logout", app.storeLogout)

	fs := http.FileServer(http.Dir("./assets/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fs))

	return app.authenticate(mux)
}
