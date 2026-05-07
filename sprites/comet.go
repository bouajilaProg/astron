package sprites

var Comet1 = SpriteDef{
	Art: []string{
		"      .      ",
		"  .  :  .    ",
		" ::::::::(@)",
		"  '  :  '    ",
		"      '      ",
	},
	DefaultVX: 0.3,
	DefaultVY: 0.1,
	Layer:     8,
}

var Comet2 = SpriteDef{
	Art: []string{
		"     .---.          ",
		"====(     )  .      ",
		"     `---'          ",
	},
	DefaultVX: 0.4,
	DefaultVY: 0.0,
	Layer:     8,
}

var Comets = []SpriteDef{Comet1, Comet2}
