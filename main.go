package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var shouldReload = true

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if shouldReload {
		shouldReload = false
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("reload"))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

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
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[1;32m"
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handler)
	port := 8080
	fmt.Println(colorGreen + "Starting server! port: " + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	for err != nil {
		fmt.Println(colorRed + "Failed to bind port: " + strconv.Itoa(port) + colorReset)
		port += 1
		fmt.Println(colorGreen + "Starting server! port: " + strconv.Itoa(port))
		err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}
}
