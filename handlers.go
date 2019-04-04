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
	var imgloc string
	var success string
	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(10 << 20)
	username := r.FormValue("username")
	title := r.FormValue("title")
	message := r.FormValue("message")
	tags := r.FormValue("hashtags")
	travelcapsule := r.FormValue("travelcapsule")

	hashtags := strings.Split(tags, ",")

	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		imgloc = ""
	} else {
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

		tempFile.Write(fileBytes)
		imgloc = tempFile.Name() + format
		// fmt.Println("Non empty file")
	}

	// fmt.Println("IMGLOC: "+imgloc)

	success = utils.CreatePost(travelcapsule, title, message, imgloc, hashtags, username, Session)

	if success != "" {
		jsonString = `{ "result": "successfully uploaded", "token": "` + utils.GenerateJWT(username, Session) + "\", \"travelcapsule\" : \"" + success + "\" }"
	} else {
		jsonString = `{ "error": "could not create post", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	}

	w.Write([]byte(jsonString))

}

func likePost(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	postId := r.Form.Get("postid")

	if utils.LikePost(postId, username, Session) {
		jsonString = `{ "result": "liked", "token": "` + utils.GenerateJWT(username, Session) + "\"}"
	} else {
		jsonString = `{ "error": "could not like", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	}

	w.Write([]byte(jsonString))
}

func unlikePost(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	postId := r.Form.Get("postid")

	if utils.UnlikePost(postId, username, Session) {
		jsonString = `{ "result": "unliked", "token": "` + utils.GenerateJWT(username, Session) + "\"}"
	} else {
		jsonString = `{ "error": "could not unlike", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	}

	w.Write([]byte(jsonString))
}

func createTC(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	title := r.Form.Get("title")

	success := utils.CreateTC(title, username, Session)

	if success != "" {
		jsonString = `{ "result": "successfully created", "token": "` + utils.GenerateJWT(username, Session) + "\", \"travelcapsule\" : \"" + success + "\" }"
	} else {
		jsonString = `{ "error": "could not create tc", "token": "` + utils.GenerateJWT(username, Session) + "\" }"
	}

	w.Write([]byte(jsonString))
}