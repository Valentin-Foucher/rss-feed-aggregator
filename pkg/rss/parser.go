package rss

import (
	"time"

	"github.com/mmcdole/gofeed"
)

func GetItemsFromFeeds(feed_urls []string) ([]IItem, error) {
	var items []IItem
	fp := gofeed.NewParser()

	for _, url := range feed_urls {
		feed, err := fp.ParseURL(url)
		for _, data := range feed.Items {
			item := GetItem(data)
			items = append(items, item)
		}
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

type IItem interface {
	Title() string
	Content() string
	PublishedDate() string
	ParsedPublishedDate() *time.Time
}

type Item struct {
	data *gofeed.Item
}

func (item *Item) Title() string {
	return item.data.Title
}

func (item *Item) Content() string {
	return item.data.Content
}

func (item *Item) PublishedDate() string {
	return item.data.Published
}

func (item *Item) ParsedPublishedDate() *time.Time {
	return item.data.PublishedParsed
}

func GetItem(data *gofeed.Item) IItem {
	item := new(Item)
	item.data = data
	return item
}
