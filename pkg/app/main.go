package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/utils"
)

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

	var cfg utils.Configuration
	utils.ReadConfiguration(&cfg)

	iter, err := rss.GetItemsIteratorFromFeeds(cfg.RssFeeds)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	paginator := newPaginator(iter)
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			MainLayout(gtx, paginator)
			e.Frame(gtx.Ops)
		}
	}
}
