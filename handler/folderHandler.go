package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"server/config"
	"strings"
	"strconv"
)

type file struct {
	Name string
	Path string
	Size string
}

type folder struct {
	Name string
	Path string
}

type fileSystem struct {
	Parent  *folder
	Files   []file
	Folders []folder
}

// FolderHandler handles requests when no file specified and returns content of specified folder.
type FolderHandler struct{}

func (handler *FolderHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	folderPath := filepath.ToSlash(string(config.Instance.WorkingDirectory) + request.URL.Path)

	if !checkFileSystemItemExists(folderPath) {
		http.NotFound(writer, request)
		return
	}

	fs, err := createFileSystemRepresentation(folderPath, request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := loadTemplate()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(writer, "fileSystem", &fs)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func createFileSystemRepresentation(folderPath string, request *http.Request) (fileSystem, error) {
	fs := fileSystem{Files: make([]file, 0, 10), Folders: make([]folder, 0, 10)}

	fs.Parent = getParentFromPath(request.URL.Path, request)

	files, err := ioutil.ReadDir(folderPath)

	if err != nil {
		return fs, err
	}

	for _, fsItem := range files {
		if fsItem.IsDir() {
			fs.Folders = append(fs.Folders, folder{Name: fsItem.Name(), Path: getURLFromName(fsItem.Name(), request)})
			continue
		}

		fs.Files = append(fs.Files, file{Name: fsItem.Name(), Path: getURLFromName(fsItem.Name(), request), Size: strconv.FormatInt(fsItem.Size(), 10) + " bytes"})
	}

	return fs, nil
}

func getParentFromPath(path string, request *http.Request) *folder {
	if path == "/" {
		return nil
	}

	folderPath := filepath.ToSlash(string(config.Instance.WorkingDirectory))
	folderPath = filepath.Join(folderPath, path, "..")

	_, err := os.Stat(folderPath)
	if err != nil {
		return nil
	}

	return &folder{Name: "..", Path: getURLFromName("..", request)}
}

func getURLFromName(path string, request *http.Request) string {
	requestPath := request.URL.Path
	if !strings.HasSuffix(requestPath, "/") {
		requestPath = requestPath + "/"
	}

	return requestPath + path
}

func loadTemplate() (*template.Template, error) {
	templatePath, _ := filepath.Abs("./templates")

	template := template.New("fileSystem")
	var err error

	templates, _ := ioutil.ReadDir(templatePath)

	for _, tmpl := range templates {
		if !strings.HasSuffix(tmpl.Name(), ".html") || tmpl.IsDir() {
			fmt.Println("Unable to parse: " + tmpl.Name())
			continue
		}

		filePath := templatePath + string(filepath.Separator) + tmpl.Name()

		content, _ := ioutil.ReadFile(filePath)

		if strings.HasPrefix(tmpl.Name(), template.Name()) {
			_, err = template.Parse(string(content))
			if err != nil {
				break
			}

			continue
		}

		_, err = template.New(tmpl.Name()).Parse(string(content))
		if err != nil {
			break
		}
	}

	return template, err
}
