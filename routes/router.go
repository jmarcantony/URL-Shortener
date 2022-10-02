package routes

import (
	"fmt"
	"net/http"
	"url_shortner/middleware/logger"

	"github.com/boltdb/bolt"
)

var bucketName = []byte("urls")

type Router struct {
	http.ServeMux
	db *bolt.DB
}

func (r *Router) Init(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket: %s", bucketName)
		}
		return nil
	})
	r.db = db
	r.Handle("/", logger.LogHandler(http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/add", logger.Log(r.Add))
	r.HandleFunc("/u", logger.Log(r.Redircet))
}
