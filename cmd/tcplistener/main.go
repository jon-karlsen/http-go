package main

import (
        "fmt"
        "log"
        "bytes"
        "net"
)

func handleConn(conn net.Conn, bytesToRead int) <- chan string {
        out := make(chan string, 1)

        go func() {
                defer conn.Close()

                line := ""
                for {
                        buf := make([]byte, bytesToRead)

                        n, err := conn.Read(buf)
                        if err != nil {
                                //out <- fmt.Sprintf("Connection closed from %s: %v\n", conn.RemoteAddr().String(), err)
                                if line != "" {
                                        out <- line
                                }
                                close(out)
                                break
                        }

                        buf = buf[:n]
                        if i := bytes.IndexByte(buf, '\n'); i != -1 {
                                line += string(buf[:i])
                                out <- fmt.Sprintf("%s\n", line)
                                buf = buf[i+1:]
                                line = ""
                        }

                        line += string(buf)
                }
        }()

        return out
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

                out := handleConn(conn, bytesToRead)
                for line := range(out) {
                        fmt.Print(line)
                }
        }
}
