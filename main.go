package main

import (
	"fmt"
	"os"
	"log"
	"bytes"
	"io"
)

func getLinesChannel(f io.ReadCloser, bytesToRead int) <- chan string {
	ch := make(chan string)

	go func() {
		line := ""
		for {
			buf := make([]byte, bytesToRead)

			n, err := f.Read(buf)
			if err != nil {
				break
			}

			buf = buf[:n]
			if i := bytes.IndexByte(buf, '\n'); i != -1 {
				line += string(buf[:i])
				ch <- line
				buf = buf[i+1:]
				line = ""
			}

			line += string(buf)
		}

		if len(line) > 0 {
			ch <- line
		}

		close(ch)
	}()

	return ch
}

func main() {
	filePath := "messages.txt"
	bytesToRead := 8

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	
	stream := getLinesChannel(file, bytesToRead)
	for line := range stream {
		fmt.Printf("read: %s\n", line)
	}
}
//
// Create a string variable to hold the contents of the "current line" of the file. It needs to persist between reads (loop iterations).
// After reading 8 bytes, split the data on newlines (\n) to create a slice of strings - let's call these split sections "parts". There will typically only be one or two "parts" because we're only reading 8 bytes at a time.
// For each part except the last one, print a line to the console in this format:
// read: LINE
//
// Where LINE is the "current line" we've aggregated so far plus the current "part". Then reset the "current line" variable to an empty string. Note that if we only have one "part", we don't need to print, as we have not reached a new line yet.
//
// Add the last "part" to the "current line" variable. Repeat until you reach the end of the file.
// Once you're done reading the file, if there's anything left in the "current line" variable, print it in the same read: LINE format.
