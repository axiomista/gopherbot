package main

import (
	"time"

	"github.com/axiomista/gopherbot"
)

const (
	greenVisor = iota
	redVisor
	cylonVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

func main() {
	visor := gopherbot.Visor()

	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	mode := redVisor

	for {
		if right.Pushed() {
			mode++
			if mode > xmasVisor {
				mode = greenVisor
			}
		}

		if left.Pushed() {
			mode--
			if mode < greenVisor {
				mode = xmasVisor
			}
		}

		switch mode {
		case greenVisor:
			visor.Green()
		case redVisor:
			visor.Red()
		case cylonVisor:
			visor.Cylon()
		case xmasVisor:
			visor.Xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
