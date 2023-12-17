package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")

	if !isValidFilename(filename) {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("./storage", filename)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file information:", err)
		http.Error(w, "Error getting file information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", http.DetectContentType([]byte(filename)))

	http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
}

func isValidFilename(filename string) bool {
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_."
	for _, char := range filename {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return false
	}

	return true
}
