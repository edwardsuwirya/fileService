package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./tmp"

type fileUploadHandler struct {
}

func NewFileUploadHandler() IHandler {
	return &fileUploadHandler{}
}

func (h *fileUploadHandler) Handler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
	}
	defer file.Close()

	fileSize := handler.Size

	if fileSize > maxUploadSize {
		http.Error(w, "File To Big", http.StatusBadRequest)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		http.Error(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	newPath := filepath.Join(uploadPath, handler.Filename)

	newFile, err := os.Create(newPath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Can not create file", http.StatusBadRequest)
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes); err != nil {
		http.Error(w, "Can not write file", http.StatusBadRequest)
		return
	}
	w.Write([]byte(newPath))
}
