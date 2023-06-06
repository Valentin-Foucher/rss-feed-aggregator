package main

// import (
// 	"fmt"

// 	"gioui.org/io/pointer"
// 	"gioui.org/layout"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// 	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
// )

// func Button(gtx layout.Context, iterator rss.IItemsIterator) layout.Dimensions {
// 	var startButton widget.Clickable
// 	btn := material.Button(th, &startButton, "Start")

// 	pointer.InputOp{
// 		Tag:   btn,
// 		Types: pointer.Press | pointer.Release,
// 	}.Add(gtx.Ops)

// 	for _, e := range gtx.Events(btn) {
// 		fmt.Print(e)
// 		if e, ok := e.(pointer.Event); ok {
// 			if e.Type == pointer.Press {
// 				iterator.Next()
// 			} else {
// 				fmt.Print(e.Type)
// 			}
// 		} else {
// 			fmt.Print(ok)
// 		}
// 	}

// 	if startButton.Clicked() {
// 		fmt.Print("ok")
// 		iterator.Next()
// 	} else {
// 		fmt.Print(btn.Button.Clicks())
// 	}
// 	return btn.Layout(gtx)
// }
