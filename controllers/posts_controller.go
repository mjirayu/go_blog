package controllers

import (
	"blog/helpers"
	"blog/models"
	"net/http"
)

// PostsHandler : Get All Posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetPosts()
	if err != nil {
		return
	}

	helpers.RenderTemplate(w, posts, "app", "posts/index")
}
