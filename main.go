package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		p := filepath.Join(".", filepath.Clean(r.URL.Path))

		if f, err := os.Stat(p); errors.Is(err, os.ErrNotExist) || f.IsDir() {
			http.ServeFile(w, r, "index.html")
		} else {
			http.ServeFile(w, r, p)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
