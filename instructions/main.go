package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	mainDir = "attacks"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", mainDir, path))
	if err != nil {
		log.Printf("failed to find file address=%s\n", path)

		_, _ = fmt.Fprint(w, err)
	}

	http.ServeContent(w, r, path, time.Now(), bytes.NewReader(data))
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))

	http.HandleFunc("/health", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(writer, "OK")
	})
	http.HandleFunc("/download", downloadHandler)

	log.Printf("ftp server started on :%d ...", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("failed to start ftp server error=%v\n", err)
	}
}
