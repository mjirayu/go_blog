package controllers

import (
	"blog/helpers"
	"net/http"
)

// Home : describe what this function does
func Home(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, nil, "app", "pages/home")
}
