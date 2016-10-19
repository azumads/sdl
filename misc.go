package sdl

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Creates a new file upload http request with optional extra params
func NewfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func DownloadFromUrl(fileUrl string) (fileName string, err error) {
	response, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Error while downloading", fileUrl, "-", err)
		return
	}
	defer response.Body.Close()

	if arr := strings.Split(response.Header.Get("content-disposition"), `"`); len(arr) > 2 {
		fileName, _ = url.QueryUnescape(arr[1])
	} else {
		fileName = strconv.Itoa(int(time.Now().Unix()))
	}

	fmt.Println("Downloading", fileUrl, "to", fileName)
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}

	defer output.Close()
	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", fileUrl, "-", err)
		return
	}
	return
}
