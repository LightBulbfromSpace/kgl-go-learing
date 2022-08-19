package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(post Post, renderer *PostRenderer) postViewModel {
	vm := postViewModel{Post: post}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(post.Body), renderer.mdParser, nil))
	return vm
}

func (pr *PostRenderer) Render(w io.Writer, post Post) error {
	return pr.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(post, pr))
	/*if err := pr.templ.Execute(w, post); err != nil {
		return err
	}
	return nil*/
}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return pr.templ.ExecuteTemplate(w, "index.gohtml", posts)
	/*indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}*/

	/*templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil*/
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
