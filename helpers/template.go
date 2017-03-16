package helpers

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate : Test
func RenderTemplate(w http.ResponseWriter, data interface{}, layout string, filenames ...string) {
	files := []string{
		"templates/application/footer.gohtml",
		"templates/application/header.gohtml",
		fmt.Sprintf("templates/layout/%s.gohtml", layout),
	}

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.gohtml", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, fmt.Sprintf("%s.gohtml", layout), data)
}
