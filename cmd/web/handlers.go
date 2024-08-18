package main

import (
	"net/http"

	"github.com/druc/posty/internal/forms"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	posts, err := app.posts.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	render(w, r, "home.page.html", pageData{"Posts": posts})
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {
	render(w, r, "post.create.page.html", nil)
}

func (app *app) storePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content")
	form.MaxLength("title", 255)

	if !form.Valid() {
		render(w, r, "post.create.page.html", pageData{"Form": form})
		return
	}

	err = app.posts.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
