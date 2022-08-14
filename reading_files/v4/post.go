package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

type Separator string

const (
	titleSeparator       Separator = "Title: "
	descriptionSeparator Separator = "Description: "
	tagsSeparator        Separator = "Tags: "
	bodySeparator        Separator = "---"
)

var MetaSeparators = []Separator{
	titleSeparator,
	descriptionSeparator,
	tagsSeparator,
}

//NewPost parses file contents code
func NewPost(postFile io.Reader) (Post, error) {
	post := Post{}

	scanner := bufio.NewScanner(postFile)

	for i := 0; i < len(MetaSeparators); i++ {
		post.readMetaLine(scanner)
	}
	post.readBody(scanner, bodySeparator)

	return post, nil
}

func (p *Post) readMetaLine(scanner *bufio.Scanner) {
	scanner.Scan()
	line := scanner.Text()

	if strings.Contains(line, string(titleSeparator)) {
		p.Title = strings.TrimPrefix(scanner.Text(), string(titleSeparator))
	} else if strings.Contains(line, string(descriptionSeparator)) {
		p.Description = strings.TrimPrefix(scanner.Text(), string(descriptionSeparator))
	} else if strings.Contains(line, string(tagsSeparator)) {
		p.Tags = strings.Split(strings.TrimPrefix(scanner.Text(), string(tagsSeparator)), ", ")
	}
}

func (p *Post) readBody(scanner *bufio.Scanner, separator Separator) {
	scanner.Scan()
	for scanner.Text() != string(separator) {
		if !scanner.Scan() {
			return
		}
	}
	buff := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buff, scanner.Text())
	}
	p.Body = strings.TrimSuffix(buff.String(), "\n")
}

/*
func (p *Post) readMetaLine(scanner *bufio.Scanner) {
	scanner.Scan()
	line := scanner.Text()
	if strings.Contains(line, string(titleSeparator)) {
		p.Title = strings.TrimPrefix(scanner.Text(), string(titleSeparator))
	} else if strings.Contains(line, string(descriptionSeparator)) {
		p.Description = strings.TrimPrefix(scanner.Text(), string(descriptionSeparator))
	} else if strings.Contains(line, string(tagsSeparator)) {
		p.Tags = strings.Split(strings.TrimPrefix(scanner.Text(), string(tagsSeparator)), ", ")
	}
}
*/

/*
func (p *Post) readMetaLine(scanner *bufio.Scanner, separator Separator) { //доработать функцию так,
	scanner.Scan() // чтобы она не зависела от внешнего "раздилителя"
	switch separator {
	case titleSeparator:
		p.Title = strings.TrimPrefix(scanner.Text(), string(titleSeparator))
	case descriptionSeparator:
		p.Description = strings.TrimPrefix(scanner.Text(), string(descriptionSeparator))
	case tagsSeparator:
		p.Tags = strings.Split(strings.TrimPrefix(scanner.Text(), string(tagsSeparator)), ", ")
	}
}
*/
