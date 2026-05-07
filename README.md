# astron

[![Release](https://github.com/bouajilaProg/astron/actions/workflows/release.yml/badge.svg)](https://github.com/bouajilaProg/astron/actions)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bouajilaProg/astron)](https://github.com/bouajilaProg/astron)
[![License](https://img.shields.io/github/license/bouajilaProg/astron)](LICENSE)

Terminal space scene with animated sprites, built with Go, Bubble Tea, and Lip Gloss.

![astron preview](https://placehold.co/800x400/080811/7EF9FF?text=astron+preview)

## Features
- Layered sprite rendering with per-entity z-depth
- Individual movement behaviors per sprite
- Animated frames for select entities
- Opaque fills to avoid transparency bleed-through
- Responsive to terminal resize

## Run
```bash
go run .
```

## Build
```bash
go build -ldflags="-s -w" -o astron .
```

## Controls
- `q` or `ctrl+c`: quit

## Structure
```
astron/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── .gitignore
├── sprites/
│   ├── sprites.go
│   ├── comet.go
│   ├── meteor.go
│   ├── planet.go
│   ├── satellite.go
│   ├── station.go
│   ├── ufo.go
│   └── star.go
└── ALL_SPRITES.txt
```

## Releases
Binaries for Linux, macOS, and Windows are available on the [Releases](https://github.com/bouajilaProg/astron/releases) page.