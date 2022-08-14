package blogposts_test

import (
	"errors"
	blogposts "github.com/LightBulbfromSpace/kld-go-learning/reading_files/v2"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("failed to open")
}

func TestNewBlogpost(t *testing.T) {
	t.Run("error check", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("Expected to get error but didn't get one")
		}
	})
	t.Run("title check", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world1.md": {Data: []byte("Title: Post 1")},
			"hello_world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)
		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
