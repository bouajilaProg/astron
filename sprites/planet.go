package sprites

var PlanetHuge = SpriteDef{
	Art: []string{
		"          _..-===========--.._          ",
		"      _.-'####################'-._      ",
		"    .'############################'.    ",
		"   /################################\\   ",
		"  |##################################|  ",
		"  |##==============================##|  ",
		"  |##################################|  ",
		"   \\################################/   ",
		"    '.############################.'    ",
		"      '-._####################_.-'      ",
		"          `--============--'            ",
	},
	DefaultVX: 0.0,
	DefaultVY: 0.0,
	Layer:     2,
}

var PlanetRinged = SpriteDef{
	Art: []string{
		"          _..._          ",
		"        .'#####'.        ",
		"  _..---|#######|---.._  ",
		" `---.._|#######|_..---' ",
		"        '._____.'        ",
	},
	DefaultVX: 0.0,
	DefaultVY: 0.0,
	Layer:     2,
}

var PlanetSmall = SpriteDef{
	Art: []string{
		"   .---.   ",
		" _/#####\\_ ",
		"(#########)",
		" \\_#####_/ ",
		"   '---'   ",
	},
	DefaultVX: 0.0,
	DefaultVY: 0.0,
	Layer:     7,
}

var PlanetsBg = []SpriteDef{PlanetHuge, PlanetRinged}
var PlanetsFg = []SpriteDef{PlanetSmall}
