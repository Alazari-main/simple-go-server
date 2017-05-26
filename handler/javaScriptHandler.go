package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"server/config"
	"path/filepath"
)

// JavaScriptHandler handles requests for JS files.
type JavaScriptHandler struct{}

func (handler *JavaScriptHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if !checkNodeExists() {
		log.Fatal("This handler requires NODE.JS to work.")
	}

	filePath := filepath.ToSlash(string(config.Instance.WorkingDirectory) + request.URL.Path)

	if !checkFileSystemItemExists(filePath) {
		http.NotFound(writer, request)
		return
	}

	var output bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("node", filePath)
	cmd.Stdout = &output
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(writer, filePath + " responded:")

		writer.Write(stderr.Bytes())
		writer.Write([]byte(err.Error()))
		
		return
	}

	fmt.Fprintln(writer, filePath + " responded:")

	writer.Write(output.Bytes())
}

// checkNodeExists returns true when NODE.JS is installed.
func checkNodeExists() bool {
	var output bytes.Buffer

	cmd := exec.Command("node", "-v")
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.String())
	result, _ := regexp.MatchString(`^v[0-9]+\.[0-9]+\.[0-9]+`, output.String())

	return result
}
