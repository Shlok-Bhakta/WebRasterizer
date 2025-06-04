package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// read the file from the header and find it in the public folder
	log.Default().Printf("Method: %s, URL: %s, Header: %v\n", r.Method, r.URL, r.Header)
	if r.URL.Path == "/" {
		// Since there is no path look for index.html in the public folder
		filedata, err := os.ReadFile("./public/index.html")
		if err != nil {
			// file not found, return 404
			http.Error(w, "File not found", http.StatusNotFound)
			log.Default().Println("File not found:", err)
			return
		} else {
			// file found, return the file
			w.Write(filedata)
			return
		}
	} else {
		// find the file in the public folder
		filedata, err := os.ReadFile("./public" + r.URL.Path)
		if err != nil {
			// file not found, return 404
			http.Error(w, "File not found", http.StatusNotFound)
			log.Default().Println("File not found:", err)
			return
		}
		// file found, return the file
		w.Write(filedata)
		return

	}
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
