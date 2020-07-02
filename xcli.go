package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

const noteText = `KEYBINDINGS
Monotoring Log Sp-Api
TAB: CHECK STATUS SERVICE
^C: Exit
`

type xWidget struct {
	name string
	x, y int
	w, h int
	body string
}

type xDataList struct {
	name string
	x, y int
	w, h int
	body string
}

var xd xDataList
var xW xWidget

func newNotice(title string, x, y int, body string) *xWidget {
	lines := strings.Split(body, "\n")
	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}
	h := len(lines) + 1
	w = w + 1

	xW.name = title
	xW.x = x
	xW.h = h
	xW.w = w
	xW.body = body
	xW.y = y

	return &xW
}

func listDataLog(title string, x, y int, body string) *xDataList {
	xd.body = body
	xd.name = title
	xd.h = 10
	xd.x = x
	xd.y = y
	xd.w = 100

	return &xd
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true

	notice := newNotice("Monitore Sp-api", 1, 0, noteText)
	layout := listDataLog("Logrus", 30, 0, "List Api")
	g.SetManager(notice, layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (w *xWidget) Layout(g *gocui.Gui) error {
	//_, maxY := g.Size()
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, w.body)
	}

	if v, err := g.SetView("side", 1, 8, w.x+w.w, w.y+w.h); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Error"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorRed
		fmt.Fprintln(v, "Item 1")
		fmt.Fprintln(v, "Item 2")
		fmt.Fprintln(v, "Item 3")
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}
	return nil
}

func (d *xDataList) Layout(g *gocui.Gui) error {
	v, err := g.SetView(d.name, d.x, d.y, d.x+d.w, d.y+d.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, d.body)
	}
	v.Title = "Success"
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
