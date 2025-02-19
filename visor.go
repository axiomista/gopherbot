package gopherbot

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

// VisorDevice controls the Gopherbot Visor Neopixel LED.
type VisorDevice struct {
	ws2812.Device
	LED     []color.RGBA
	rg      bool
	forward bool
	pos     int
}

// Visor returns a new VisorDevice to control Gopherbot Visor.
func Visor() *VisorDevice {
	neo := machine.A3
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	v := ws2812.New(neo)

	return &VisorDevice{
		Device: v,
		LED:    make([]color.RGBA, VisorLEDCount),
	}
}

// Show sets the visor to display the current LED array state.
func (v *VisorDevice) Show() {
	v.WriteColors(v.LED)
}

// Off turns off all the LEDs.
func (v *VisorDevice) Off() {
	v.Clear()
}

// Clear clears the visor.
func (v *VisorDevice) Clear() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	v.Show()
}

// Red turns all of the Visor LEDs red.
func (v *VisorDevice) Red() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	}

	v.Show()
}

// Green turns all of the Visor LEDs green.
func (v *VisorDevice) Green() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	}

	v.Show()
}

// Blue turns all of the Visor LEDs blue.
func (v *VisorDevice) Blue() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
	}

	v.Show()
}

// Purple turns all of the Visor LEDs purple.
func (v *VisorDevice) Purple() {
	for i := range v.LED {
		v.LED[i] = color.RGBA{R: 0x80, G: 0x00, B: 0x80}
	}

	v.Show()
}

// Xmas light style
func (v *VisorDevice) Xmas() {
	v.rg = !v.rg
	for i := range v.LED {
		v.rg = !v.rg
		if v.rg {
			v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		} else {
			v.LED[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
		}
	}

	v.Show()
}

// Cylon visor mode.
func (v *VisorDevice) Cylon() {
	if v.forward {
		v.pos += 2
		if v.pos >= VisorLEDCount {
			v.pos = VisorLEDCount - 2
			v.forward = false
		}
	} else {
		v.pos -= 2
		if v.pos < 0 {
			v.pos = 0
			v.forward = true
		}
	}

	for i := 0; i < VisorLEDCount; i += 2 {
		if i == v.pos {
			v.LED[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			v.LED[i+1] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		} else {
			v.LED[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			v.LED[i+1] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}
	}

	v.Show()
}


// Rainbow visor mode. 
func (v *VisorDevice) Rainbow() {
	var r map[int]color.RGBA
	r = make(map[int]color.RGBA)

	// Pick some rainbowish colors for the map, as many as you want
	// (If you have more colors than LEDs, the remaining colors won't appear)
	r[0] = color.RGBA{R: 0xff, G: 0xff, B: 0xff} // White
	r[1] = color.RGBA{R: 0xff, G: 0xff, B: 0x00} // Yellow
	r[2] = color.RGBA{R: 0xff, G: 0x7f, B: 0x00} // Orange
	r[3] = color.RGBA{R: 0xff, G: 0x00, B: 0x00} // Red
	r[4] = color.RGBA{R: 0xff, G: 0x00, B: 0x7f} // Pink
	r[5] = color.RGBA{R: 0x94, G: 0x00, B: 0xd3} // Violet
	r[6] = color.RGBA{R: 0x4b, G: 0x00, B: 0x82} // Purple
	r[7] = color.RGBA{R: 0x00, G: 0x00, B: 0xff} // Blue
	r[8] = color.RGBA{R: 0x46, G: 0xf0, B: 0xf0} // Cyan
	r[9] = color.RGBA{R: 0x00, G: 0x98, B: 0x88} // Teal
	r[10] = color.RGBA{R: 0x00, G: 0xff, B: 0x00} // Green
	r[11] = color.RGBA{R: 0xbc, G: 0xf6, B: 0x0c} // Lime
	m := len(r)

	// For each position, choose a color in the map
	for i := 0; i < VisorLEDCount; i ++ {
		v.LED[i] = r[i%m]
	}
	v.Show()
}
