package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/druc/posty/internal/models/sqlite"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	posts   *sqlite.PostModel
	users   *sqlite.UserModel
	session *sessions.CookieStore
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.NewCookieStore([]byte("vlD0fatfJFfGt9FmaZnUKcC2nn0GeSYH"))
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteLaxMode

	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},
		users: &sqlite.UserModel{
			DB: db,
		},
		session: session,
	}

	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listing on :8000")
	srv.ListenAndServe()
}
