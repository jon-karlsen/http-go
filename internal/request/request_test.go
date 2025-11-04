package request

import (
        "strings"
        "testing"

        "github.com/stretchr/testify/assert"
        "github.com/stretchr/testify/require"
)

func TestRequestLineParse(t *testing.T) {
        // TEST: Good GET request line
        test_input := "GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"
        r, err := RequestFromReader(strings.NewReader(test_input))

        require.NoError(t, err)
        require.NotNil(t, r)

        assert.Equal(t, "GET", r.RequestLine.Method)
        assert.Equal(t, "/", r.RequestLine.RequestTarget)
        assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

        // TEST: Good GET Request line with path
        test_input = "GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"
        r, err = RequestFromReader(strings.NewReader(test_input))

        require.NoError(t, err)
        require.NotNil(t, r)

        assert.Equal(t, "GET", r.RequestLine.Method)
        assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
        assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

        // TEST: Invalid number of parts in request line (too few)
        test_input = "/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"
        _, err = RequestFromReader(strings.NewReader(test_input))
        require.Error(t, err)
        assert.EqualError(t, err, "Error parsing request line; invalid number of parts")

        // TEST: Invalid number of parts in request line (too many)
        test_input = "GET /coffee latte HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"
        _, err = RequestFromReader(strings.NewReader(test_input))
        require.Error(t, err)
        assert.EqualError(t, err, "Error parsing request line; invalid number of parts")

        // TEST: Invalid version in Request line
        test_input = "GET / HTTP/1.2\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"
        r, err = RequestFromReader(strings.NewReader(test_input))
        require.Error(t, err)
        assert.EqualError(t, err, "Invalid version 'HTTP/1.2'; only 'HTTP/1.1' is supported")
}
