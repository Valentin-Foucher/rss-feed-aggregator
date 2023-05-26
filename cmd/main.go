package main

import (
	"fmt"

	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
)

func main() {
	if iter, err := rss.FromFeeds([]string{"https://feeds.simplecast.com/54nAGcIl", "https://www.dailymail.co.uk/sciencetech/index.rss"}); err == nil {
		for iter.HasNext() {
			itemsWindow := iter.Next()
			for _, item := range itemsWindow {
				if item != nil {
					fmt.Println(item.Title())
				}
			}
		}
	} else {
		fmt.Println(err)
	}

	// reader := strings.NewReader(feed.Items[0].Content)
	// browser.OpenReader(reader)
}
