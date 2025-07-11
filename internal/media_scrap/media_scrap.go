package mediascrap

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/ranon7/1500_webscraping/internal/commons"
	"golang.org/x/sync/errgroup"
)

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
	var forceDownload bool

	fs := flag.NewFlagSet("mediascrap", flag.ExitOnError)
	fs.StringVar(&board, "board", "", "Valid board name from 1500chan.org. Required")
	fs.StringVar(&thread, "thread", "", "Thread number must exist in the selected board. Required")
	fs.Var(&formats, "formats", "A list of file formats as a comma separated string. E.g.: jpg,mp4,webm. Required")
	fs.StringVar(&location, "location", "", "The location where the media will be saved into your file system. If it doesn't exist it'll be created. Required")
	fs.BoolVar(&verbose, "verbose", false, "Enable detailed logs. Disabled by default")
	fs.IntVar(&m, "m", 30, "Controls the parallism - how many files being downloaded at a time. Defaults to 30")
	fs.BoolVar(&forceDownload, "force", false, "Whether to force download existing files in the destination folder")

	if err := fs.Parse(args); err != nil {
		return err
	}

	reqArgs := []string{"board", "thread", "formats", "location"}
	if err := commons.ValidateArgs(reqArgs, fs); err != nil {
		return err
	}

	if verbose {
		commons.EnableVerboseLogging()
	}

	commons.VerboseLogger.Println("running media_scrap")

	threadUrl := buildThreadUrl(board, thread)
	downloadFolder := buildDownloadLocation(location, board, thread)
	total := 0
	urlCh := make(chan string)
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(m)

	commons.VerboseLogger.Printf("thread url %s", threadUrl)
	commons.VerboseLogger.Printf("download location %s", downloadFolder)
	if err := ensureDir(downloadFolder); err != nil {
		return err
	}

	go func() {
		for url := range urlCh {
			if validateHref(url, formats) {
				commons.VerboseLogger.Printf("selected for download %s", url)
				group.Go(func() error {
					path := buildPathFromUrl(url, downloadFolder)
					err, exists := fileExists(path)
					if err != nil {
						return err
					}
					if exists && !forceDownload {
						commons.VerboseLogger.Printf("%s already exists, not downloading again", path)
					} else {
						if err := downloadFile(ctx, url, path); err != nil {
							return err
						}
					}
					return nil
				})
				total++
			} else {
				commons.VerboseLogger.Printf("download skipped for file %s", url)
			}
		}
	}()

	posts, err := getThreadPosts(ctx, threadUrl)
	if err != nil {
		return err
	}

	for _, post := range posts {
		if _, exists := post["filename"]; exists {
			tim := post["tim"].(string)
			ext := post["ext"].(string)

			url := buildFileUrl(board, tim, ext)
			urlCh <- url
		}
	}

	if err := group.Wait(); err != nil {
		return err
	}

	commons.Logger.Printf("finished, files saved to %s", downloadFolder)

	return nil
}
