package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Kishorsenthilkumar/blogposts"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello.md":  {Data: []byte("hi")},
		"hello1.md": {Data: []byte("how r u")},
	}
	posts, err := blogposts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(fs) != len(posts) {
		t.Errorf("got %d wanted %d", len(posts), len(fs))
	}
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello.md":  {Data: []byte("Title: Post1")},
		"hello1.md": {Data: []byte("Title: Post2")},
	}
	posts, err := blogposts.NewPostsFromFs(fs)

	got := posts[0]
	want := blogposts.Post{Title: "Post1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	if err != nil {
		t.Fatal(err)
	}

}
