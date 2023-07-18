package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("failed to start ftp server error=%v\n", err)
	}
}
