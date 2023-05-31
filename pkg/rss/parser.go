package rss

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type ItemsCollection struct {
	Items []IItem
}

func GetItemsFromFeeds(feed_urls []string) (*ItemsCollection, error) {
	var items []IItem
	fp := gofeed.NewParser()
	result := new(ItemsCollection)

	for _, url := range feed_urls {
		feed, err := fp.ParseURL(url)
		for _, data := range feed.Items {
			item := getItem(data)
			items = append(items, item)
		}
		if err != nil {
			return nil, err
		}
	}
	result.Items = items
	return result, nil
}

type IItem interface {
	Title() string
	Description() string
	PublishedDate() string
	ParsedPublishedDate() *time.Time
}

type Item struct {
	data *gofeed.Item
}

func (item *Item) Title() string {
	return item.data.Title
}

func (item *Item) Description() string {
	return item.data.Description
}

func (item *Item) PublishedDate() string {
	return item.data.Published
}

func (item *Item) ParsedPublishedDate() *time.Time {
	return item.data.PublishedParsed
}

func getItem(data *gofeed.Item) IItem {
	item := new(Item)
	item.data = data
	return item
}
