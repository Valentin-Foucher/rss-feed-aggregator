package rss

type Direction uint

type IItemsIterator interface {
	HasNext() bool
	HasPrevious() bool
	Next() []IItem
	Previous() []IItem
}

const (
	Right Direction = 0
	Left            = 1
)

type ItemsIterator struct {
	index  int
	window int

	direction Direction
	items     []IItem
}

func (i *ItemsIterator) HasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false

}
func (i *ItemsIterator) Next() []IItem {
	if i.HasNext() {
		if i.direction == Left {
			i.index = i.index + i.window
		}

		last := i.index + i.window
		if last > len(i.items)-1 {
			last = len(i.items) - 1
		}
		result := i.items[i.index:last]
		i.index = last
		i.direction = Right
		return result
	}
	return nil
}

func (i *ItemsIterator) HasPrevious() bool {
	if i.index > 0 {
		return true
	}
	return false

}
func (i *ItemsIterator) Previous() []IItem {
	if i.HasPrevious() {
		if i.direction == Right {
			i.index = i.index - i.window
		}

		first := i.index - i.window
		if first < 0 {
			first = 0
		}
		result := i.items[first:i.index]
		i.index = first
		i.direction = Left
		return result
	}
	return nil
}

func (coll *ItemsCollection) CreateIterator() IItemsIterator {
	return &ItemsIterator{
		index:  0,
		window: 10,
		items:  coll.Items,
	}
}
