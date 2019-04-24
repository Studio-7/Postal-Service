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
	"encoding/json"
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

func getpost(w http.ResponseWriter, r *http.Request) {
	var resp interface{}
	var resps []interface{}
	username := r.Form.Get("username")
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	postids := r.Form.Get("ids")
	simplified := r.Form.Get("simplified")
	ids := strings.Split(postids, ",")
	fmt.Println(ids)
	for _, id := range ids {
		id = strings.TrimSpace(id)
		fmt.Println(id)
		if simplified == "true" {
			resp = utils.GetPost(id, true, Session)
		} else {
			resp = utils.GetPost(id, false, Session)
		}
		resps = append(resps, resp)
	}
	j, _ := json.Marshal(resps)
	jsonString := `{ "result": ` + string(j) + `, "token": "` + utils.GenerateJWT(username, Session) + "\"}"
	w.Write([]byte(jsonString))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	var imgloc string
	var success string
	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(32 << 20)
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
	}

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

func addComment(w http.ResponseWriter, r *http.Request) {
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	message := r.Form.Get("message")
	postId := r.Form.Get("postid")

	if utils.AddComment(postId, username, message, Session) {
		jsonString = `{ "result": "comment added successfully", "token": "` + utils.GenerateJWT(username, Session) + "\"}"
	} else {
		jsonString = `{ "error": "could not add comment", "token": "` + utils.GenerateJWT(username, Session) + "\"}"
	}

	w.Write([]byte(jsonString))
}

// func getComments(w http.ResponseWriter, r *http.Request) {
// 	var comments interface{}
// 	r.ParseForm()
// 	w.Header().Set("Content-Type", "application/json")
// 	postId := r.Form.Get("postid")

// 	comments = utils.GetComments(postId, Session)
// 	resp, _ := json.Marshal(comments)
// 	w.Write(resp)
// }