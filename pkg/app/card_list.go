package main

import (
	"fmt"
	"image"

	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"

	"gioui.org/layout"
	"gioui.org/unit"
)

func CardListLayout(gtx layout.Context, size image.Point, iterator rss.IItemsIterator) layout.Dimensions {
	items := iterator.Next()

	list := layout.List{Axis: layout.Vertical, Alignment: layout.Middle}
	list.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
		item := items[i]
		return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return Card(gtx, item.Title(), item.Description())
		})
	})
	return layout.Dimensions{Size: size}
}

func x() {
	iter, _ := rss.GetItemsIteratorFromFeeds([]string{"https://feeds.simplecast.com/54nAGcIl", "https://www.dailymail.co.uk/sciencetech/index.rss"})
	fmt.Print(iter)
}
