package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: 5 * time.Second}
	req, err := c.Get("http://golang.org")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	println(string(body))
}
