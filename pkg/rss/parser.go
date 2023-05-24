package rss

import "github.com/mmcdole/gofeed"

func GetItemsFromFeeds(feed_urls []string) ([]*gofeed.Item, error) {
	var items []*gofeed.Item
	fp := gofeed.NewParser()

	for _, url := range feed_urls {
		feed, err := fp.ParseURL(url)
		items = append(items, feed.Items...)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}
