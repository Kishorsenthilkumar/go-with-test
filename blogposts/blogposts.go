package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFs(FileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(FileSystem, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post

	for _, fname := range dir {
		post, err := getPost(FileSystem, fname.Name())

		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	postfile, err := filesystem.Open(filename)

	if err != nil {
		return Post{}, err
	}
	defer postfile.Close()

	return newPost(postfile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner:=bufio.NewScanner(postFile)

	scanner.Scan()
	TitleLine:=scanner.Text()

	scanner.Scan()
	DescriptionLine:=scanner.Text()

	post:=Post{Title: string(TitleLine)[7:],Description: string(DescriptionLine)[13:]}

	return post,nil
}
