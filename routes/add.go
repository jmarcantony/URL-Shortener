package routes

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/boltdb/bolt"
)

const slugSize = 5

var chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateSlug() []byte {
	slug := make([]byte, slugSize)
	for i := range slug {
		slug[i] = chars[rand.Intn(len(chars))]
	}
	return slug
}

// BpLnf
func (R *Router) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			w.Write([]byte("An Error occured when parsing your request"))
		}
		url := r.FormValue("url")
		var slug []byte
		R.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucketName)
			for true {
				pug := generateSlug()
				if b.Get(slug) == nil {
					slug = pug
					break
				}
			}
			err := b.Put(slug, []byte(url))
			http.Redirect(w, r, "/?slug="+string(slug), http.StatusFound)
			return err
		})
	} else {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}
}
