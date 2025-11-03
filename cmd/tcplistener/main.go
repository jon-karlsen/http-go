package main

import (
        "fmt"
        "log"
        "bytes"
        "net"
)

func handleConn(conn net.Conn, bytesToRead int) {
        defer conn.Close()

        line := ""
        for {
                buf := make([]byte, bytesToRead)

                n, err := conn.Read(buf)
                if err != nil {
                        fmt.Printf("Connection closed from %s: %v\n", conn.RemoteAddr().String(), err)
                        return
                }

                buf = buf[:n]
                if i := bytes.IndexByte(buf, '\n'); i != -1 {
                        line += string(buf[:i])
                        fmt.Printf("%s\n", line)
                        buf = buf[i+1:]
                        line = ""
                }

                line += string(buf)
        }
}

func main() {
        bytesToRead := 8
        port := ":42069"

        listener, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("Failed to create TCP listener: %v", err)
        }
        defer listener.Close()

        for {
                conn, err := listener.Accept()
                if err != nil {
                        fmt.Printf("Failed to create accept connection: %v\n", err)
                        continue
                }

                go handleConn(conn, bytesToRead)
        }
}
