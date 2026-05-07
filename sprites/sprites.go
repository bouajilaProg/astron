package sprites

import (
	"github.com/charmbracelet/lipgloss"
)

type SpriteDef struct {
	Art        []string
	Frames     [][]string
	DefaultVX  float64
	DefaultVY  float64
	IsVertical bool
	Layer      int
	UpdateFunc func(e *Entity, screenWidth, screenHeight int)
}

type Entity struct {
	Art        []string
	Frames     [][]string
	X, Y       float64
	VX, VY     float64
	Tick       float64
	Layer      int
	Style      lipgloss.Style
	UpdateFunc func(e *Entity, screenWidth, screenHeight int)
}

func (e *Entity) Width() int {
	maxWidth := 0
	for _, line := range e.Art {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

func (e *Entity) Height() int {
	return len(e.Art)
}

func (e *Entity) Update(screenWidth, screenHeight int) {
	e.Tick++

	if len(e.Frames) > 0 {
		frameIdx := (int(e.Tick) / 4) % len(e.Frames)
		e.Art = e.Frames[frameIdx]
	}

	if e.UpdateFunc != nil {
		e.UpdateFunc(e, screenWidth, screenHeight)
	} else {
		e.X += e.VX
		e.Y += e.VY
	}

	w := float64(e.Width())
	h := float64(e.Height())

	if e.X > float64(screenWidth) {
		e.X = -w
	} else if e.X < -w {
		e.X = float64(screenWidth)
	}

	if e.Y > float64(screenHeight) {
		e.Y = -h
	} else if e.Y < -h {
		e.Y = float64(screenHeight)
	}
}
