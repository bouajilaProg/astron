package sprites

import "math"

var UFOClassic = SpriteDef{
	Frames: [][]string{
		{
			"     .---.     ",
			"  _/__~*~__\\_  ",
			" (###########) ",
			"   .  .  .  .  ",
		},
		{
			"     .---.     ",
			"  _/__~*~__\\_  ",
			" (###########) ",
			"   '  '  '  '  ",
		},
	},
	DefaultVX: 0.35,
	DefaultVY: 0.0,
	Layer:     6,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.X += e.VX
		e.Y += math.Sin(e.Tick*0.1) * 0.1
	},
}

var UFORetro = SpriteDef{
	Frames: [][]string{
		{
			"     _     ",
			"   /[ ]\\   ",
			" _(##o##)_ ",
			"  '-_-_-'  ",
		},
		{
			"     _     ",
			"   /[ ]\\   ",
			" _(##o##)_ ",
			"  '-\"-\"-'  ",
		},
	},
	DefaultVX: -0.25,
	DefaultVY: 0.0,
	Layer:     6,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.X += e.VX
		e.Y += math.Cos(e.Tick*0.08) * 0.15
	},
}

var Rocket = SpriteDef{
	Frames: [][]string{
		{
			"    ^    ",
			"   /#\\   ",
			"  |#A#|  ",
			"  |#R#|  ",
			" /|#C#|\\ ",
			"| |_|_| |",
			" '==^==' ",
			"  ( v )  ",
		},
		{
			"    ^    ",
			"   /#\\   ",
			"  |#A#|  ",
			"  |#R#|  ",
			" /|#C#|\\ ",
			"| |_|_| |",
			" '==^==' ",
			"   (V)   ",
		},
		{
			"    ^    ",
			"   /#\\   ",
			"  |#A#|  ",
			"  |#R#|  ",
			" /|#C#|\\ ",
			"| |_|_| |",
			" '==^==' ",
			"    v    ",
		},
	},
	DefaultVX:  0.0,
	DefaultVY:  -0.2, // Moves UP
	IsVertical: true,
	Layer:      5,
	UpdateFunc: func(e *Entity, screenWidth, screenHeight int) {
		e.Y += e.VY // Only moves up
	},
}

var UFOs = []SpriteDef{UFOClassic, UFORetro, Rocket}
