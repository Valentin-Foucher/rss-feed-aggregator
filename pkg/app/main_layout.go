package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
)

var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
	white      = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	grey       = color.NRGBA{R: 0xDE, G: 0xDE, B: 0xDE, A: 0xFF}
)

func MainLayout(gtx layout.Context) layout.Dimensions {
	iter, err := rss.GetItemsIteratorFromFeeds([]string{"https://feeds.simplecast.com/54nAGcIl", "https://www.dailymail.co.uk/sciencetech/index.rss"})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return Background(gtx, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return CardListLayout(gtx, iter)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			b := new(Button)
			return b.Layout(gtx)
		}),
	)
}

func main() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			MainLayout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
