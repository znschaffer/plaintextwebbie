package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

func main() {
	staticRoot, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.FS(staticRoot)))
	http.ListenAndServe(":"+port, nil)
}
