package screenshot

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func GetScreenshot(url, elementID string) []byte {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
	}

	if err := chromedp.Run(ctx, chromedp.WaitVisible("#"+elementID, chromedp.ByID)); err != nil {
		log.Fatal(err)
	}

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Screenshot("#"+elementID, &buf, chromedp.ByID)); err != nil {
		log.Fatal(err)
	}

	return buf
}
