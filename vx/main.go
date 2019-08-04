package main

import (
	"image/color"
	"time"

	"github.com/axiomista/gopherbot"
)


func main() {
	antenna := gopherbot.Antenna()
	backpack := gopherbot.Backpack()
	visor := gopherbot.Visor()
	speaker := gopherbot.Speaker()
	left := gopherbot.LeftButton()
	right := gopherbot.RightButton()

	color1 := color.RGBA{R: 0x80, G: 0x00, B: 0x80}
	color2 := color.RGBA{R: 0xff, G: 0xa5, B: 0x00}

	go antenna.Blink()
	visor.Rainbow()

	for {
		if left.Pushed() {
			speaker.Blip()
			color1 = color.RGBA{R: 0x80, G: 0x00, B: 0x80}
			color2 = color.RGBA{R: 0xff, G: 0xa5, B: 0x00}

		}
		if right.Pushed() {
			speaker.Blip()
			color1 = color.RGBA{R: 0xff, G: 0x00, B: 0x7f}
			color2 = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
		}
		backpack.Alternate(color1, color2)

		// Iterate through the map and move the colors over forever, waiting 200ms
		for j := range visor.LED {
			visor.LED[j] = visor.LED[(j+1)%gopherbot.VisorLEDCount]
		}
		visor.Show()
		time.Sleep(200 * time.Millisecond)
	}
}
