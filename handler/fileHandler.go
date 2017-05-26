package handler

import (
	"net/http"
	"path/filepath"
	"server/config"
)

type FileHandler struct{}

func (handler *FileHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	filePath := filepath.ToSlash(string(config.Instance.WorkingDirectory) + request.URL.Path)

	if !checkFileSystemItemExists(filePath) {
		http.NotFound(writer, request)
		return
	}

	http.ServeFile(writer, request, filePath)
}
