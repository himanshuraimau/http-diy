package main

import (
	"io"
	"log"
	"os"
	"strings"
	"fmt"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		data := make([]byte, 8)
		currentLine := ""

		for {
			n, err := f.Read(data)
			if err != nil {
				if err == io.EOF {
					if currentLine != "" {
						lines <- currentLine
					}
					return
				}
				log.Fatal(err)
			}

			parts := strings.Split(string(data[:n]), "\n")
			for i, part := range parts {
				if i < len(parts)-1 {
					lines <- currentLine + part
					currentLine = ""
				} else {
					currentLine += part
				}
			}
		}
	}()

	return lines
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
