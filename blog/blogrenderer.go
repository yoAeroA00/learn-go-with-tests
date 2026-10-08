package blog

import (
	"fmt"
	"io"
)

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\nTags: <ul>", p.Title, p.Description)
	if err != nil {
		return err
	}

	for i, tag := range p.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
		if i == len(p.Tags)-1 {
			_, err := fmt.Fprintf(w, "</ul>")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
