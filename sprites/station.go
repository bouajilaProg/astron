package sprites

import "math"

var SpaceStation = SpriteDef{
	Art: []string{
		"      _|_          ",
		" [I]--[H]--[I]     ",
		"  |    |    |      ",
		"==[=]--o--[=]==    ",
		"  |    |    |      ",
		" [I]--[H]--[I]     ",
		"      _|_          ",
	},
	DefaultVX: 0.02,
	DefaultVY: 0.0,
	Layer:     3,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.X += e.VX
		e.Y += math.Sin(e.Tick*0.02) * 0.02
	},
}

var Stations = []SpriteDef{SpaceStation}
