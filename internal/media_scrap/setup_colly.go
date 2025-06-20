package mediascrap

import "github.com/gocolly/colly/v2"

func setupColly(hrefCh chan<- string) *colly.Collector {
	c := colly.NewCollector(colly.AllowedDomains(caravelaDomain))
	c.SetCookies(caravelaBaseUrl, cookies)

	c.OnHTML(".fileinfo > a", func(e *colly.HTMLElement) {
		hrefCh <- e.Attr("href")
	})

	c.OnScraped(func(r *colly.Response) {
		close(hrefCh)
	})

	return c
}
