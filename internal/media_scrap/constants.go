package mediascrap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	logger        = log.New(os.Stdout, "info  ", log.LstdFlags)
	verboseLogger = log.New(io.Discard, "debug ", log.LstdFlags)
)
