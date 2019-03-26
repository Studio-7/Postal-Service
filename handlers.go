package main

import (
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cvhariharan/Utils/utils"
)

// Guess image format from gif/jpeg/png/webp
func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}

// Guess image mime types from gif/jpeg/png/webp
func guessImageMimeTypes(r io.Reader) string {
	format, _ := guessImageFormat(r)
	if format == "" {
		return ""
	}
	return mime.TypeByExtension("." + format)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseMultipartForm(10 << 20)
	username := r.FormValue("username")
	title := r.FormValue("title")
	message := r.FormValue("message")
	tags := r.FormValue("hashtags")

	hashtags := strings.Split(tags, ",")

	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	format := guessImageMimeTypes(file)
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	tempFile, err := os.OpenFile("./temp-images/upload-"+timestamp+format, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	imgloc := tempFile.Name() + format
	fmt.Println(imgloc)
	// return that we have successfully uploaded our file!
	if utils.CreatePost(title, message, imgloc, hashtags, username, Session) {
		jsonString = `{ "result": "successfully uploaded", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	} else {
		jsonString = `{ "error": "could not create post", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	}

	w.Write([]byte(jsonString))

}
