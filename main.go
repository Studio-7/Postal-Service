package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cvhariharan/Utils/utils"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var Session *r.Session

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
	endpoints := os.Getenv("DBURL")
	url := strings.Split(endpoints, ",")
	dbpass := os.Getenv("DBPASS")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url,
		Password: dbpass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	Session = s
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/post/createpost", utils.AuthMiddleware(createPost, Session))
	http.HandleFunc("/post/like", utils.AuthMiddleware(likePost, Session))
	http.HandleFunc("/post/unlike", utils.AuthMiddleware(unlikePost, Session))
	http.HandleFunc("/post/createtravelcapsule", utils.AuthMiddleware(createTC, Session))
	log.Fatal(http.ListenAndServe(port, nil))
}
