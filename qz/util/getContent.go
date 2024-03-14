package util

import (
	"io"
	"net/http"
)

// GetContent gets the content of the url

func GetContent(url string) ([]byte, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

//to detect content type see https://golang.org/pkg/net/http/#DetectContentType
// http.DetectContentType(buf)      https://mimesniff.spec.whatwg.org/
