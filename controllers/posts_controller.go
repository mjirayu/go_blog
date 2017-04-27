package controllers

import (
	"blog/helpers"
	"blog/models"
	"fmt"
	"net/http"
)

// PostsHandler : Get All Posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// session, err := store.Get(r, "session")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	fmt.Println(r.Context().Value("currentUser"))

	posts, err := models.GetPosts()
	if err != nil {
		return
	}

	helpers.RenderTemplate(w, posts, "app", "posts/index")
}
