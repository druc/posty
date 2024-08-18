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

func (app *app) getLogin(w http.ResponseWriter, r *http.Request) {
	render(w, r, "login.page.html", nil)
}

func (app *app) storeLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id, err := app.users.Authenticate(
		r.PostForm.Get("email"),
		r.PostForm.Get("password"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	session, _ := app.session.Get(r, "posty")
	session.Values["userId"] = id
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func (app *app) storeLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := app.session.Get(r, "posty")
	delete(session.Values, "userId")
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func (app *app) getRegister(w http.ResponseWriter, r *http.Request) {
	render(w, r, "register.page.html", nil)
}

func (app *app) storeRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = app.users.Insert(
		r.PostForm.Get("name"),
		r.PostForm.Get("email"),
		r.PostForm.Get("password"),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/login", 302)
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
