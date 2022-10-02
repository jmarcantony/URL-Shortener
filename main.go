package main

import (
	"log"
	"net/http"
	"url_shortner/routes"

	"github.com/boltdb/bolt"
)

const port = "8080"

func main() {
	db, err := bolt.Open("urls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := new(routes.Router)
	r.Init(db)
	log.Println("Server listening on port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
