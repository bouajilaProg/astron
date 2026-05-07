package sprites

import "math"

var Satellite = SpriteDef{
	Art: []string{
		"  _|_  ",
		"-[___]-",
		"  _|_  ",
		" /   \\ ",
	},
	DefaultVX: -0.05,
	DefaultVY: 0.0,
	Layer:     4,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.X += e.VX
		// Slight hover up and down
		e.Y += math.Sin(e.Tick*0.05) * 0.05
	},
}

var Probe = SpriteDef{
	Art: []string{
		"  .-.  ",
		"=[_I_]-(",
		"  'V'  ",
	},
	DefaultVX: 0.06,
	DefaultVY: 0.0,
	Layer:     4,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.X += e.VX
		e.Y += math.Cos(e.Tick*0.03) * 0.03
	},
}

var Satellites = []SpriteDef{Satellite, Probe}
