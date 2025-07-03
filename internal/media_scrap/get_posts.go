package mediascrap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ranon7/1500_webscraping/internal/commons"
)

func getThreadPosts(ctx context.Context, url string) ([]map[string]any, error) {
	var data map[string][]map[string]any

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	tryNumber := 1
	delay := initialDelay

	for {
		resp, err := http_client.Do(req)
		if err == nil {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				commons.VerboseLogger.Printf("error reading response body: %v", err)
				continue
			}

			if err := json.Unmarshal(bodyBytes, &data); err != nil {
				commons.VerboseLogger.Printf("failure unmarshalling json %v", err)
				continue
			}

			posts := data["posts"]
			return posts, nil
		} else {
			if tryNumber > maxRetries {
				commons.VerboseLogger.Println("failed to download but retries exceded")
				return nil, fmt.Errorf("failed to download %s: %w", url, err)
			}
			commons.VerboseLogger.Println("failed to get posts, will retry")

			time.Sleep(initialDelay)
			delay *= 2
			tryNumber++

			if resp != nil {
				resp.Body.Close()
			}
		}

	}
}
