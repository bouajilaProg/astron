# astron

Fast-loading terminal space scene built with Go, Bubble Tea, and Lip Gloss.

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

## Controls
- `q` or `ctrl+c`: quit

## Structure
```
astron/
├── main.go
├── go.mod
├── sprites/
│   ├── sprites.go
│   ├── planet.go
│   ├── ufo.go
│   ├── comet.go
│   ├── meteor.go
│   ├── satellite.go
│   ├── station.go
│   └── star.go
└── ALL_SPRITES.txt
```
