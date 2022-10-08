package files

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetFile(path string) (string, error) {
	if strings.HasPrefix(path, "http") {
		return path, nil
	} else {
		return GetFileEncoding(path)
	}
}

func GetFileEncoding(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	contentType := http.DetectContentType(content)

	encoded = "data:" + contentType + ";base64," + encoded

	return encoded, nil
}
