package main

import "github.com/mmcdole/gofeed"

type IItemsIterator interface {
	hasNext() bool
	next() []*gofeed.Item
}
