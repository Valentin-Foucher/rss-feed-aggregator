package rss

import "github.com/mmcdole/gofeed"

type IItemsIterator interface {
	HasNext() bool
	Next() []*gofeed.Item
}

type ItemsIterator struct {
	index  int
	window int

	items []*gofeed.Item
}

func (i *ItemsIterator) HasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false

}
func (i *ItemsIterator) Next() []*gofeed.Item {
	if i.HasNext() {
		items := i.items[i.index : i.index+i.window]
		i.index += i.window
		return items
	}
	return nil
}

func GetItemsIterator(items []*gofeed.Item) IItemsIterator {
	return &ItemsIterator{
		index:  0,
		window: 5,
		items:  items,
	}
}
