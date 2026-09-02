package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFs(FileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(FileSystem, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, fname := range dir {
		post, err := getPost(FileSystem, fname.Name())
	}
	if err != nil {
		return nil, err
	}

}
