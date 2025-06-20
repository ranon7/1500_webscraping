package mediascrap

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func enableVerboseLogging() {
	verboseLogger.SetOutput(os.Stdout)
}

func Run(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer cancel()

	var board string
	var thread string
	var formats FileFormats
	var location string
	var verbose bool
	var m int

	fs := flag.NewFlagSet("mediascrap", flag.ExitOnError)
	fs.StringVar(&board, "board", "", "Valid board name from 1500chan.org. Required")
	fs.StringVar(&thread, "thread", "", "Thread number must exist in the selected board. Required")
	fs.Var(&formats, "formats", "A list of file formats as a comma separated string. E.g.: jpg,mp4,webm. Required")
	fs.StringVar(&location, "location", "", "The location where the media will be saved into your file system. If it doesn't exist it'll be created. Required")
	fs.BoolVar(&verbose, "verbose", false, "Enable detailed logs. Disabled by default")
	fs.IntVar(&m, "m", 30, "Controls the parallism - how many files being downloaded at a time. Defaults to 30")

	if err := fs.Parse(args); err != nil {
		return err
	}

	reqArgs := []string{"board", "thread", "formats", "location"}
	if err := validateArgs(reqArgs, fs); err != nil {
		return err
	}

	if verbose {
		enableVerboseLogging()
	}

	verboseLogger.Println("running media_scrap")

	threadUrl := buildThreadUrl(board, thread)
	downloadFolder := buildDownloadLocation(location, board, thread)
	total := 0
	hrefCh := make(chan string)
	c := setupColly(hrefCh)
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(m)

	verboseLogger.Printf("thread url %s", threadUrl)
	verboseLogger.Printf("download location %s", downloadFolder)
	if err := ensureDir(downloadFolder); err != nil {
		return err
	}

	go func() {
		for href := range hrefCh {
			if validateHref(board, href, formats) {
				url := buildUrl(href)
				verboseLogger.Printf("selected for download %s", url)
				group.Go(func() error {
					path := buildPathFromUrl(url, downloadFolder)
					err, exists := fileExists(path)
					if err != nil {
						return err
					}
					if exists {
						verboseLogger.Printf("%s already exists, skipped", path)
					} else {
						if err := downloadFile(ctx, url, path); err != nil {
							return err
						}
					}
					return nil
				})
				total++
			}
		}
	}()

	if err := c.Visit(threadUrl); err != nil {
		return err
	}

	if err := group.Wait(); err != nil {
		return err
	}

	logger.Printf("finished, files saved to %s", downloadFolder)

	return nil
}
