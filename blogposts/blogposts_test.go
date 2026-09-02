package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Kishorsenthilkumar/blogposts"
)

func TestNewline(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)
	fs := fstest.MapFS{
		"sample1.md": {Data: []byte(firstBody)},
		"sample2.md": {Data: []byte(secondBody)},
	}

	post, err := blogposts.NewPostsFromFs(fs)

	got := post[0]
	want := blogposts.Post{"Post 1", "Description 1"}

	assertPost(t, got, want)

	if err != nil {
		t.Fatal(err)
	}

}

func assertPost(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
