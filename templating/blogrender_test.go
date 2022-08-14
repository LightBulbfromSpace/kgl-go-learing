package blogrenderer_test

import (
	"bytes"
	blogrenderer "github.com/LightBulbfromSpace/kld-go-learning/templating"
	approvals "github.com/approvals/go-approval-tests"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "This is title",
			Tags:        []string{"tag1", "tag2"},
			Description: "This is description",
			Body:        "This is body",
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buff := bytes.Buffer{}

		err = postRenderer.Render(&buff, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buff.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buff := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "My post"}, {Title: "MY POST 2"}}

		if err := postRenderer.RenderIndex(&buff, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buff.String())
	})
}
