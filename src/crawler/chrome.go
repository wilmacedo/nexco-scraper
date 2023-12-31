package crawler

import (
	"context"
	"log"
	"time"

	"go-scraper/src/config"

	"github.com/chromedp/chromedp"
)

func Run(domain config.Domain) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	timeout := time.Duration(config.GetEnv().Timeout) * time.Second
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	start := time.Now()
	var value string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(domain.URL),
		chromedp.Text(domain.Tag, &value, chromedp.NodeVisible),
	)

	log.Printf("Finished crawler in %s", time.Since(start))

	return value, err
}
