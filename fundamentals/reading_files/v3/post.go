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

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	bodySeparator        = "---"
)

//NewPost parses file contents code
func NewPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(separator string) string {
		notEnd := scanner.Scan()
		if notEnd != false {
			return strings.TrimPrefix(scanner.Text(), separator)
		}
		return ""
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buff := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buff, scanner.Text())
	}
	return strings.TrimSuffix(buff.String(), "\n")
}

/*
	var title, description = titleLine, descriptionLine
	scanner := bufio.NewScanner(postFile)

	notEnd := scanner.Scan()
	if notEnd != false {
		title = scanner.Text()
	}

	notEnd = scanner.Scan()
	if notEnd != false {
		description = scanner.Text()
	}

	return Post{Title: title[7:], Description: description[13:]}, nil
*/

/*const (
	titleLine       = "Title: "
	descriptionLine = "Description: "
)

//NewPost parses file contents code
func NewPost(postFile io.Reader) (Post, error) {
	var title, description = titleLine, descriptionLine
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title = readLine()[len(titleLine):]
	description = readLine()[len(descriptionLine):]
	return Post{Title: title, Description: description}, nil
}*/
