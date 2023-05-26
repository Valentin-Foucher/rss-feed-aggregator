package rss

type IItemsIterator interface {
	HasNext() bool
	Next() []IItem
}

type ItemsIterator struct {
	index  int
	window int

	items []IItem
}

func (i *ItemsIterator) HasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false

}
func (i *ItemsIterator) Next() []IItem {
	if i.HasNext() {
		items := i.items[i.index : i.index+i.window]
		i.index += i.window
		return items
	}
	return nil
}

func (coll *ItemsCollection) CreateIterator() IItemsIterator {
	return &ItemsIterator{
		index:  0,
		window: 5,
		items:  coll.Items,
	}
}
