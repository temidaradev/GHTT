package handlers

import (
	"main/views/blog"
	"net/http"
)

func HandleBlog(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, blog.Blog())
}
