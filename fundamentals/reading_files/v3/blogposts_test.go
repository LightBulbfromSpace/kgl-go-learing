package blogposts_test

import (
	"errors"
	blogposts "github.com/LightBulbfromSpace/kld-go-learning/reading_files/v3"
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
	t.Run("error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("Expected to get error but didn't get one")
		}
	})
	/*t.Run("title", func(t *testing.T) {
			fs := fstest.MapFS{
				"hello_world1.md": {Data: []byte("Title: Post 1")},
				"hello_world2.md": {Data: []byte("Title: Post 2")},
			}

			posts, _ := blogposts.NewPostsFromFS(fs)

			assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})
		})
		t.Run("title and description", func(t *testing.T) {
			const (
				firstBody = `Title: Post 1
	Description: Description 1`
				secondBody = `Title: Post 2
	Description: Description 2`
			)
			fs := fstest.MapFS{
				"Number1.md": {Data: []byte(firstBody)},
				"Number2.md": {Data: []byte(secondBody)},
			}

			posts, _ := blogposts.NewPostsFromFS(fs)

			assertPost(t, posts[0], blogposts.Post{
				Title:       "Post 1",
				Description: "Description 1",
			})
		})*/
	t.Run("title, tags and description", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: go, err`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
		)

		fs := fstest.MapFS{
			"Number1.md": {Data: []byte(firstBody)},
			"Number2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"go", "err"},
		})
	})
	t.Run("head and body", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: go, err
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
My
Awesome
Post`
		)
		fs := fstest.MapFS{
			"Number1.md": {Data: []byte(firstBody)},
			"Number2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostsFromFS(fs)

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"go", "err"},
			Body: `Hello
World`,
		})
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v but want %+v", got, want)
	}
}
