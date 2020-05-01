package main

import (
	"math/rand"
	"time"

	"github.com/jiripetrlik/evo-art/gui"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	gui.GuiServer()
}
