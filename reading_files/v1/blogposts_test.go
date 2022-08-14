package blogposts_test

import (
	"errors"
	blogposts "github.com/LightBulbfromSpace/kld-go-learning/reading_files/v1"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("failed to open")
}

func TestNewBlogpost(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello_world.md": {Data: []byte("hi")},
			"hello_word.md":  {Data: []byte("word")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
	t.Run("error check", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("Expected to get error but didn't get one")
		}
	})
}
