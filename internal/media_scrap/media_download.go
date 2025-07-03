package mediascrap

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ranon7/1500_webscraping/internal/commons"
)

func downloadFile(ctx context.Context, url string, path string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	tryNumber := 1
	delay := initialDelay

	for {
		commons.VerboseLogger.Printf("attempt %d to download file", tryNumber)
		resp, err := http_client.Do(req)
		if err == nil {
			out, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("error creating file %s: %w", path, err)
			}
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil {
				return fmt.Errorf("error writing to file %s: %w", path, err)
			}
			break
		} else {
			if tryNumber > maxRetries {
				commons.VerboseLogger.Println("failed to download but retries exceded")
				return fmt.Errorf("failed to download %s: %w", url, err)
			}
			commons.VerboseLogger.Println("failed to download, will retry")

			time.Sleep(initialDelay)
			delay *= 2
			tryNumber++

			if resp != nil {
				resp.Body.Close()
			}
		}
	}

	return nil
}
