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
		cwd, err := os.Getwd()
		if err != nil {
			return
		}
		p := filepath.Join(cwd, filepath.Clean(r.URL.Path))

		if f, err := os.Stat(p); errors.Is(err, os.ErrNotExist) || f.IsDir() {
			os.Stderr.WriteString("File not found: " + p + ", serving index.html instead\n")
			http.ServeFile(w, r, "index.html")
		} else {
			http.ServeFile(w, r, p)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
