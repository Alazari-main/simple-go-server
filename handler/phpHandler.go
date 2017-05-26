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

// PhpHandler handles requests for PHP files.
type PhpHandler struct{}

func (handler *PhpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if !checkPhpExists() {
		log.Fatal("This handler requires PHP to work.")
	}

	filePath := filepath.ToSlash(string(config.Instance.WorkingDirectory) + request.URL.Path)

	if !checkFileSystemItemExists(filePath) {
		http.NotFound(writer, request)
		return
	}

	var output bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("php", filePath)
	cmd.Stdout = &output
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {

		fmt.Fprintln(writer, filePath + " responded:")
		
		writer.Write(output.Bytes())
		writer.Write([]byte(err.Error()))

		return
	}

	fmt.Fprintln(writer, filePath + " responded:")

	writer.Write(output.Bytes())
}

// checkNodeExists returns true when PHP is installed.
func checkPhpExists() bool {
	var output bytes.Buffer

	cmd := exec.Command("php", "-v")
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	result, _ := regexp.MatchString(`^(PHP)\s[0-9]+\.[0-9]+\.[0-9]+`, output.String())

	return result
}
