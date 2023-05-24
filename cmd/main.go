package main

import (
	"fmt"
	"sort"

	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
)

func sortItems(items []rss.IItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].ParsedPublishedDate().After(*items[j].ParsedPublishedDate())
	})
}

func fromFeeds(feedUrls []string) error {
	items, err := rss.GetItemsFromFeeds(feedUrls)
	if err != nil {
		return err
	}

	sortItems(items)

	iter := rss.GetItemsIterator(items)
	for iter.HasNext() {
		itemsWindow := iter.Next()
		for _, item := range itemsWindow {
			if item != nil {
				fmt.Println(item.Title())
			}
		}
	}

	return nil
}

func main() {
	fromFeeds([]string{"https://feeds.simplecast.com/54nAGcIl", "https://www.dailymail.co.uk/sciencetech/index.rss"})

	// reader := strings.NewReader(feed.Items[0].Content)
	// browser.OpenReader(reader)
}
