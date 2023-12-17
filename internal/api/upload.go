package api

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// In your Upload function
func Upload(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/upload.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	// Check if it's a POST request
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadedFileName := generateUniqueFilename(handler.Filename)
		uploadedFilePath := filepath.Join("./storage", uploadedFileName)

		uploadedFile, err := os.Create(uploadedFilePath)
		if err != nil {
			http.Error(w, "Error creating the uploaded file", http.StatusBadRequest)
			return
		}
		defer uploadedFile.Close()

		_, err = io.Copy(uploadedFile, file)
		if err != nil {
			http.Error(w, "Error copying file content", http.StatusInternalServerError)
			return
		}

		// Display success message
		successMessage := fmt.Sprintf("File '%s' uploaded successfully.", handler.Filename)

		// Pass the success message to the template
		tmpl.Execute(w, struct{ SuccessMessage string }{SuccessMessage: successMessage})
		return
	}

	// For other methods or GET requests, serve the HTML template
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
	}
}

func generateUniqueFilename(filename string) string {
	ext := filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]

	return fmt.Sprintf("%s%s", base, ext)
}
