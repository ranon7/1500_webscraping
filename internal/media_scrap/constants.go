package mediascrap

import (
	"fmt"
	"net/http"
	"time"
)

const (
	caravelaDomain = "1500chan.org"
	caravelaScheme = "https://"
	maxRetries     = 3
	timeout        = 2 * time.Minute
	initialDelay   = 3 * time.Second
)

var (
	caravelaBaseUrl = fmt.Sprintf("%s%s", caravelaScheme, caravelaDomain)
	cookies         = []*http.Cookie{
		{
			Name:  "mc",
			Value: "1",
		},
	}
	http_client = &http.Client{
		Timeout: time.Second * 30,
	}
)
