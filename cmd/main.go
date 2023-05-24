package main

import (
	"fmt"
	"sort"

	"github.com/mmcdole/gofeed"
)

type ItemsIterator struct {
	index  int
	window int

	items []*gofeed.Item
}

func (i *ItemsIterator) hasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false

}
func (i *ItemsIterator) next() []*gofeed.Item {
	if i.hasNext() {
		items := i.items[i.index : i.index+i.window]
		i.index += i.window
		return items
	}
	return nil
}

func createItemsIterator(items []*gofeed.Item) IItemsIterator {
	return &ItemsIterator{
		index:  0,
		window: 5,
		items:  items,
	}
}

func fromFeeds(feed_urls []string) error {
	var items []*gofeed.Item
	fp := gofeed.NewParser()

	for _, url := range feed_urls {
		feed, err := fp.ParseURL(url)
		items = append(items, feed.Items...)
		if err != nil {
			return err
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PublishedParsed.After(*items[j].PublishedParsed)
	})

	iter := createItemsIterator(items)
	for iter.hasNext() {
		itemsWindow := iter.next()
		for _, item := range itemsWindow {
			if item != nil {
				fmt.Println(item.Title)
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
