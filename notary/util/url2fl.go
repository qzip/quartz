package util

import (
	"context"
	"io"
	"net/http"
	"os"
)

//URLwriter URL to io.Writer
func URLwriter(ctx context.Context, url string, out io.Writer) (written int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	written, err = io.Copy(out, resp.Body)
	return

}

//DownloadFile https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
