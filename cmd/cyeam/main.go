package main

import (
	"os"

	"github.com/mnhkahn/http"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.Serve(":" + port)
}
