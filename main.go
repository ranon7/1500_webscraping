package main

import (
	"log"
	"os"

	mediascrap "github.com/ranon7/1500_webscraping/internal/media_scrap"
)

func main() {
	_ = os.Args[0]

	if len(os.Args) < 2 {
		log.Fatal("expected at least one argument to be passed (the subcommand)")
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "media_scrap":
		if err := mediascrap.Run(args); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", subcommand)
	}
}
