package main

import (
	"fmt"
	"log"
	"net/http"
	"youtube-get-video/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/validate", handlers.ValidateHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
