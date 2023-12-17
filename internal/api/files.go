package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Files(w http.ResponseWriter, r *http.Request) {
	files, err := listAllFiles("./storage")

	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := template.ParseFiles("./templates/files.html")

	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Files []string
	}{
		Files: files,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		fmt.Println(err)
	}
}

func listAllFiles(dir string) ([]string, error) {
	var fileList []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileList = append(fileList, info.Name())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return fileList, nil
}
