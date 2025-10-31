package main

import (
	"fmt"
	"os"
	"io"
	"log"
)

func main() {
	filePath := "messages.txt"
	bytesToRead := 8

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	
	buffer := make([]byte, bytesToRead)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				if bytesRead > 0 {
					fmt.Printf("EOF reached; read %d bytes: %s\n", bytesRead, buffer[:bytesRead])
				}
				break
			}
			log.Fatalf("Failed to read from file: %v", err)
		}
		fmt.Printf("read: %s\n", buffer[:bytesRead])
	}
}
