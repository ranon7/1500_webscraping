package mediascrap

import (
	"fmt"
	"net/http"
)

var (
	caravelaDomain  = "1500chan.org"
	caravelaScheme  = "https://"
	caravelaBaseUrl = fmt.Sprintf("%s%s", caravelaScheme, caravelaDomain)
	cookies         = []*http.Cookie{
		{
			Name:  "mc",
			Value: "1",
		},
	}
)
