package main

import (
	"image/color"
	"time"

	"github.com/axiomista/gopherbot"
)

var (
	accel   *gopherbot.AccelerometerDevice
	visor   *gopherbot.VisorDevice
	speaker *gopherbot.SpeakerDevice
)

func main() {
	antenna := gopherbot.Antenna()
	backpack := gopherbot.Backpack()
	visor = gopherbot.Visor()

	go antenna.Blink()

	visor.Rainbow()
	backpack.Alternate(color.RGBA{R: 0x80, G: 0x00, B: 0x80}, color.RGBA{R: 0xff, G: 0xa5, B: 0x00})
	time.Sleep(200 * time.Millisecond)

}
