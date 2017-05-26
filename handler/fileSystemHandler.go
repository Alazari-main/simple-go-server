package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"server/config"
)

// FileSystemHandler handles requests for files and folders
type FileSystemHandler struct{}

func (handler *FileSystemHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	
	var innerHandler http.Handler
	
	fsItemPath := filepath.ToSlash(string(config.Instance.WorkingDirectory) + request.URL.Path)

	if !checkFileSystemItemExists(fsItemPath) {
		http.NotFound(writer, request)
		return
	}

	isFolder, err := fsItemIsFolder(fsItemPath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	if isFolder {
		innerHandler = &FolderHandler {}
	} else {
		innerHandler = &FileHandler {}
	}

	innerHandler.ServeHTTP(writer, request)
}

func fsItemIsFolder(fsItemPath string) (bool, error) {
	item, err := os.Stat(fsItemPath)

	return item.IsDir(), err
}
