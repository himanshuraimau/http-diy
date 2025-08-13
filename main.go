package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 8)
	currentLine := ""

	for {
		n, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				if currentLine != "" {
					fmt.Printf("read: %s\n", currentLine)
				}
				break
			}
			log.Fatal(err)
		}

		parts := strings.Split(string(data[:n]), "\n")

		for i, part := range parts {
			if i < len(parts)-1 {
				// Not the last part, print complete line
				fmt.Printf("read: %s%s\n", currentLine, part)
				currentLine = ""
			} else {
				// Last part, add to current line
				currentLine += part
			}
		}
	}
	f.Close()
}
