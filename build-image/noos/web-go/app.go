package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed ui
var UI embed.FS

func main() {
	uiFS, err := fs.Sub(UI, "ui")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(uiFS)))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
