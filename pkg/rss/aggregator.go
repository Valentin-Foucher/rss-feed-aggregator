package rss

import (
	"sort"
)

func sortItems(itemsCollection *ItemsCollection) {
	sort.Slice(itemsCollection.Items, func(i, j int) bool {
		return itemsCollection.Items[i].ParsedPublishedDate().After(*itemsCollection.Items[j].ParsedPublishedDate())
	})
}

func GetItemsIteratorFromFeeds(rssFeeds map[string][]string) (IItemsIterator, error) {
	// aggregate
	itemsCollection, err := GetItemsFromFeeds(rssFeeds)
	if err != nil {
		return nil, err
	}

	// sort historically
	sortItems(itemsCollection)

	iter := itemsCollection.CreateIterator()
	return iter, nil
}
