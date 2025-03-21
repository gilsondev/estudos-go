package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: 5 * time.Second}
	jsonVar := bytes.NewBuffer([]byte(`{"foo": "bar"}`))

	req, err := c.Post("https://httpbin.org/post", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	println(string(body))
}
