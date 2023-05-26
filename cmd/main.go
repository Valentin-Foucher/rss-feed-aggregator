package main

import (
	"fmt"
	"sort"

	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
)

func sortItems(itemsCollection *rss.ItemsCollection) {
	sort.Slice(itemsCollection.Items, func(i, j int) bool {
		return itemsCollection.Items[i].ParsedPublishedDate().After(*itemsCollection.Items[j].ParsedPublishedDate())
	})
}

func fromFeeds(feedUrls []string) error {
	// aggregate
	itemsCollection, err := rss.GetItemsFromFeeds(feedUrls)
	if err != nil {
		return err
	}

	// sort historically
	sortItems(itemsCollection)

	iter := itemsCollection.CreateIterator()
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
