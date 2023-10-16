package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	staticRoot, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(staticRoot)))
	http.ListenAndServe(":8080", nil)
}
