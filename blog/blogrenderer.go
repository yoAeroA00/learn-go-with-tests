package blog

import (
	"html/template"
	"io"
)

const postTemplate = `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

// var postTemplates embed.FS

// func Render(w io.Writer, p Post) error {
// 	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\nTags: <ul>", p.Title, p.Description)
// 	if err != nil {
// 		return err
// 	}

// 	for i, tag := range p.Tags {
// 		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
// 		if err != nil {
// 			return err
// 		}
// 		if i == len(p.Tags)-1 {
// 			_, err := fmt.Fprintf(w, "</ul>")
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

func Render(w io.Writer, p Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	// templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	return templ.Execute(w, p)
}
