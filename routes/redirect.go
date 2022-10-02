package routes

import (
	"net/http"

	"github.com/boltdb/bolt"
)

func (R *Router) Redircet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := []byte(r.URL.Query().Get("slug"))
		var url string
		R.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucketName)
			urlB := b.Get(slug)
			if urlB != nil {
				url = string(urlB)
			} else {
				w.Write([]byte("No shorteneed url for this slug"))
			}
			return nil
		})
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	} else {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}
}
