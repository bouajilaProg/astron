package sprites

var MeteorSmall = SpriteDef{
	Art: []string{
		" .-. ",
		"( o )",
		" `-' ",
	},
	DefaultVX: -0.2,
	DefaultVY: 0.15,
	Layer:     9,
}

var MeteorMedium = SpriteDef{
	Art: []string{
		"  _..---.._  ",
		" /  _  _   \\ ",
		"|  (O)  (o) |",
		" \\__'_'___ / ",
	},
	DefaultVX: -0.3,
	DefaultVY: 0.2,
	Layer:     9,
}

var Meteors = []SpriteDef{MeteorSmall, MeteorMedium}
