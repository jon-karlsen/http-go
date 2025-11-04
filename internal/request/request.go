package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
        HttpVersion     string
        RequestTarget   string
        Method          string
}

type Request struct {
        RequestLine RequestLine
}

var ERR_BAD_NO_PARTS = fmt.Errorf("Error parsing request line; invalid number of parts")
var SEPARATOR = "\r\n"

func parseRequestMethod(input *string) (*string, error) {
        for _, ch := range(*input) {
                if ch != ' ' {
                        if int(ch) < int('A') || int(ch) > int('Z') {
                                return nil, fmt.Errorf("Invalid HTTP method: '%s'", *input)
                        }
                }
        }

        return input, nil
}

func parseRequestVersion(input *string) (*string, error) {
        parts := strings.Split(*input, "/")
        if parts[len(parts)-1] != "1.1" {
                return nil, fmt.Errorf("Invalid version '%s'; only 'HTTP/1.1' is supported", *input)
        }

        return &parts[len(parts)-1], nil
}

func parseToArrOfStrings(input []byte) ([]string, []byte, error) {
        const expectedParts = 3

        i, j, n := 0, 0, len(input)
        output := make([]string, expectedParts, expectedParts)

        for i < n {
                if i+1 < n && string(input[i:i+len(SEPARATOR)]) == SEPARATOR {
                        break
                }

                if input[i] == ' ' {
                        j++
                        if j >= len(output) {
                                return nil, nil, ERR_BAD_NO_PARTS
                        }
                } else {
                        output[j] += string(input[i])
                }
                i++
        }

        if j < expectedParts-1 || output[j] == "" {
                return nil, nil, ERR_BAD_NO_PARTS
        }

        return output, input[i+len(SEPARATOR):], nil
}

func parseRequestLine(input []byte) (*RequestLine, []byte, error) {
        parsed, rem, err := parseToArrOfStrings(input)
        if err != nil {
                return nil, nil, err
        }

        method, err := parseRequestMethod(&parsed[0])
        if err != nil {
                return nil, nil, err
        }

        ver, err := parseRequestVersion(&parsed[2])
        if err != nil {
                return nil, nil, err
        }

        output := RequestLine{
                HttpVersion: *ver,
                RequestTarget: parsed[1],
                Method: *method,
        }

        return &output, rem, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
        input, err := io.ReadAll(reader)
        if err != nil {
                return nil, err
        }

        reqLine, _, err := parseRequestLine(input)
        if err != nil {
                return nil, err
        }

        output := Request{
                RequestLine: *reqLine,
        }

        return &output, nil
}
