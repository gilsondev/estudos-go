package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", "https://google.com.br", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "text/html")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	println(string(body))
}
