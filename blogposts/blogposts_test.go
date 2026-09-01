package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/Kishorsenthilkumar/blogposts"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello.md":  {Data: []byte("hi")},
		"hello1.md": {Data: []byte("how r u")},
	}
	posts := blogposts.NewPostsFromFs(fs)

	if len(fs) != len(posts) {
		t.Errorf("got %d wanted %d", len(posts), len(fs))
	}
}
