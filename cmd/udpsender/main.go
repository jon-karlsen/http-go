package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
        port := 42069

        serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", port))
        if err != nil {
                log.Fatalf("Failed to resolve UDP address: %v\n", err)
        }

        conn, err := net.DialUDP("udp", nil, serverAddr)
        if err != nil {
                log.Fatalf("Failed to dial UDP: %v\n", err)
        }
        defer conn.Close()

        reader := bufio.NewReader(os.Stdin)
        for {
                fmt.Println(">")
                line, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Printf("Error reading from stdin: %v\n", err)
                }

                n, err := conn.Write([]byte(line))
                if err != nil {
                        fmt.Printf("Error writing to UDP connection: %v\n", err)
                }

                fmt.Printf("Wrote %d bytes to UDP connection\n", n)
        }
}
