package controllers

import (
	"blog/helpers"
	"blog/models"
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var ctx context.Context

// SignUpHandler : describe what this function does
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, nil, "app", "users/signup")
}

// CreateUserHandler : describe what this function does
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: bs,
	}

	if err := user.Create(); err != nil {
		return
	}

	http.Redirect(w, r, "/login", 302)
}

// LoginHandler : describe what this function does
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, nil, "app", "users/login")
}

// CreateSessionHandler : describe
func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.GetUser(email)
	if err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sID := uuid.NewV4()
	session.Values[user.Email] = sID.String()
	session.Save(r, w)

	ctx = context.WithValue(r.Context(), "currentUser", user)

	http.Redirect(w, r, "/posts", 302)
}
