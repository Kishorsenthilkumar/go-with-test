package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFs(FileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(FileSystem, ".")
	var post []Post
	for range dir {
		post = append(post, Post{})
	}
	return post
}
