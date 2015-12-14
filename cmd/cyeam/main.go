package main

import (
	"os"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/http"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.Serve(":" + port)
}

func init() {
	go HaixiuzuStart()
}
