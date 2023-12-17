package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nihankhan/fileflow/internal/api"
)

func Routers() *mux.Router {
	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", api.Index)
	route.HandleFunc("/upload", api.Upload)
	route.HandleFunc("/files", api.Files)
	route.HandleFunc("/download", api.Download)

	route.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage"))))
	route.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	return route
}
