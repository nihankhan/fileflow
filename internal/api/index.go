package api

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		http.Error(w, "Error parsing index html template!", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
