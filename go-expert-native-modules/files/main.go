package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(filename string) *os.File {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	return f
}

func WriteFile(file *os.File, content string) {
	_, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Reading file %s...\n", filename)
	fmt.Println(string(content))
}

func ReadFileBufferMode(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return
		}

		fmt.Println(string(buffer[:n]))
	}
}

func main() {
	f := CreateFile("new-file.txt")
	WriteFile(f, "Hello, World!")
	defer f.Close()

	ReadFile("new-file.txt")
	ReadFileBufferMode("file-14MB.txt")
}
