package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/evgzor/go1fl-sprint6-final-tpl/internal/service"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	curDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(curDir, "index.html")
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := service.Convert(string(data))

	ext := filepath.Ext(handler.Filename)

	filename := time.Now().UTC().String() + ext

	dst, err := os.Create(filename)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.WriteString(dst, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

}
