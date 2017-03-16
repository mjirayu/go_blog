package models

import "time"

// Post : Post Model
type Post struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
}

// GetPost : Get Post By Id
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select * from posts where id = $1", id).Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
	return
}

// GetPosts : All Posts
func GetPosts() (posts []Post, err error) {
	rows, err := Db.Query("select * from posts")
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// GetPostsWithLimit : Test
func GetPostsWithLimit(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select * from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Create Post
func (post *Post) Create() (err error) {
	statement := "insert into posts (title, body, created_at) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Title, post.Body, time.Now()).Scan(&post.ID)
	return
}

// Update Post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set title = $2, body = $3 where id = $1", post.ID, post.Title, post.Body)
	return
}
