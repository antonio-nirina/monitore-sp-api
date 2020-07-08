package main
import (
	"fmt"
	"log"
	//"math"
	"strings"

	"github.com/jroimartin/gocui"
)
const noteText = `
Monotoring Log Sp-Api
CHECK STATUS SERVICE
^C = Exit
`

type result struct {
	st []string
	view *gocui.View
}

type helpStatus struct {
	x int
	y int 
}
var res result
var hl helpStatus
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, setListStatus(1)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, setListStatus(-1)); err != nil {
			log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// maxX, maxY := g.Size()
	w,h := sizeTitle() 
	list := []string{"Item1","Item2","Item3","Item4","Item5"}
	res.st = list
	hl.x = 1
	hl.y = 0
	if v, err := g.SetView("Monitore Sp-api", 1, 0, w, h); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, noteText)
	}
	if v, err := g.SetView("st1", 1, 6, w, 7+len(list)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Status"
		v.Highlight = true
		setState(v)
	}
	if v, err := g.SetView("Log", w+2, 0, 50, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Log"
		fmt.Fprintln(v, "Data")
	}
	return nil
}

func sizeTitle() (int,int) {
	lines := strings.Split(noteText, "\n")
	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}
	h := len(lines) 
	w = w + 2

	return w,h
}

func setListStatus(d int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		count := 0
		if hl.y < len(res.st) && d == 1 {
			count++
			// hl.y = hl.y+1
		}
		if hl.y > len(res.st) && d == -1{
			hl.y = hl.y-1
		}
		hl.y = count
		fmt.Println(count)
		// setState(res.view)
		return nil
	}
}

func setState(vs *gocui.View) {
	vs.SelBgColor = gocui.ColorGreen
	vs.SelFgColor = gocui.ColorRed
	vs.SetCursor(hl.x,hl.y)
	res.view = vs
	for _,val := range res.st {
		fmt.Fprintln(vs, val)
	}
}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

