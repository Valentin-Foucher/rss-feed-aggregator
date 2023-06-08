package rss

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type ItemsCollection struct {
	Items []IItem
}

func GetItemsFromFeeds(rssFeeds map[string][]string) (*ItemsCollection, error) {
	var items []IItem
	fp := gofeed.NewParser()
	result := new(ItemsCollection)

	for source, urls := range rssFeeds {
		for _, url := range urls {
			feed, err := fp.ParseURL(url)
			for _, data := range feed.Items {
				item := getItem(data, source)
				items = append(items, item)
			}
			if err != nil {
				return nil, err
			}
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
	Link() string
	Source() string
}

type Item struct {
	data   *gofeed.Item
	source string
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

func (item *Item) Link() string {
	return item.data.Link
}

func (item *Item) Source() string {
	return item.source
}

func getItem(data *gofeed.Item, source string) IItem {
	item := new(Item)
	item.data = data
	item.source = source
	return item
}
