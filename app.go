package main

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"net/http"
)

//go:embed public/*
var public embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("listening on", port)

	var files = fs.FS(public)
	htmlContent, err := fs.Sub(files, "public")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(htmlContent)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
