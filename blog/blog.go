package blog

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}
	defer postFile.Close()

	return newPost(postFile)
}

// func newPost(postFile io.Reader) (Post, error) {
// 	postData, err := io.ReadAll(postFile)
// 	if err != nil {
// 		return Post{}, nil
// 	}

// 	post := Post{Title: string(postData[7:])}
// 	return post, nil
// }

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readTagMetaLine := func(tagName string) []string {
		var result []string

		for _, tag := range strings.Split(readMetaLine(tagName), ",") {
			result = append(result, strings.Trim(tag, " "))
		}

		return result
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		// Tags:        strings.Split(readTagMetaLine(tagsSeparator), ", "),
		Tags: readTagMetaLine(tagsSeparator),
		Body: readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	var body strings.Builder
	for scanner.Scan() {
		fmt.Fprintln(&body, scanner.Text())
	}

	return strings.TrimSuffix(body.String(), "\n")
}
