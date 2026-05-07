package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"astron/sprites"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

const (
	tickInterval      = 50 * time.Millisecond
	starTwinkleChance = 0.08
)

type model struct {
	width, height int
	canvas        []string
	rng           *rand.Rand
	initialized   bool

	items []sprites.Entity
}

func newModel() model {
	return model{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.ensureCanvas()
		if m.width > 0 && m.height > 0 {
			m.seedEntities() // Reseed on resize to fix growing issues and maintain proportions
			m.initialized = true
		}
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case tickMsg:
		if !m.initialized {
			return m, m.Init()
		}

		for i := range m.items {
			m.items[i].Update(m.width, m.height)
			if len(m.items[i].Art) == 1 && (m.items[i].Art[0] == "." || m.items[i].Art[0] == "*") {
				if m.rng.Float64() < starTwinkleChance {
					if m.rng.Float64() < 0.5 {
						m.items[i].Art = sprites.StarBright
						m.items[i].Style = starBrightStyle()
					} else {
						m.items[i].Art = sprites.Star
						m.items[i].Style = starDimStyle()
					}
				}
			}
		}
		return m, m.Init()
	}

	return m, nil
}

func (m model) View() string {
	if m.width <= 0 || m.height <= 0 || len(m.canvas) != m.width*m.height {
		return ""
	}

	bgStyle := lipgloss.NewStyle().Background(lipgloss.Color("#080811")).Foreground(lipgloss.Color("#080811"))
	emptyChar := bgStyle.Render(" ")

	for i := range m.canvas {
		m.canvas[i] = emptyChar
	}

	draw := func(e sprites.Entity) {
		for y, line := range e.Art {
			for x, char := range line {
				if char == ' ' {
					continue
				}

				// '#' acts as a solid fill that gets rendered as an opaque space
				strChar := string(char)
				if char == '#' {
					strChar = " "
				}

				posX, posY := int(e.X)+x, int(e.Y)+y
				if posX >= 0 && posX < m.width && posY >= 0 && posY < m.height {
					idx := posY*m.width + posX
					m.canvas[idx] = e.Style.Render(strChar)
				}
			}
		}
	}

	items := make([]sprites.Entity, len(m.items))
	copy(items, m.items)
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Layer < items[j].Layer
	})
	for _, item := range items {
		draw(item)
	}

	var out strings.Builder
	for y := 0; y < m.height; y++ {
		rowStart := y * m.width
		for x := 0; x < m.width; x++ {
			out.WriteString(m.canvas[rowStart+x])
		}
		out.WriteByte('\n')
	}

	return out.String()
}

func (m *model) ensureCanvas() {
	if m.width <= 0 || m.height <= 0 {
		m.canvas = nil
		return
	}

	if len(m.canvas) != m.width*m.height {
		m.canvas = make([]string, m.width*m.height)
	}
}

func (m *model) seedEntities() {
	starCount := clampInt((m.width*m.height)/35, 60, 300)
	m.items = make([]sprites.Entity, 0, starCount+12)
	for i := 0; i < starCount; i++ {
		vx := m.randRange(-0.01, 0.0) // very slow drift left
		vy := 0.0
		star := m.makeStar(sprites.Star, starDimStyle(), vx, vy)
		star.Layer = 0
		if m.rng.Float64() < 0.35 {
			star.Art = sprites.StarBright
			star.Style = starBrightStyle()
		}
		m.items = append(m.items, star)
	}

	planetStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3F2B96")).Background(lipgloss.Color("#080811"))    // deep purple
	planetStyleAlt := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF007A")).Background(lipgloss.Color("#080811")) // neon pink
	stationStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB800")).Background(lipgloss.Color("#080811"))   // gold
	satelliteStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00E5FF")).Background(lipgloss.Color("#080811")) // cyan
	meteorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4500")).Background(lipgloss.Color("#080811"))    // fiery red
	cometStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#B026FF")).Background(lipgloss.Color("#080811"))     // neon purple
	ufoColors := []string{"#39FF14", "#FF00FF", "#FFFF00", "#00FFFF"}
	ufoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(ufoColors[m.rng.Intn(len(ufoColors))])).Background(lipgloss.Color("#080811"))

	// Foreground and Background Planets - Make them static
	bgPlanet := m.makeEntityRange(sprites.PlanetsBg[m.rng.Intn(len(sprites.PlanetsBg))], planetStyle, m.height/6, m.height/3)
	m.items = append(m.items, bgPlanet)

	fgPlanet := m.makeEntityRange(sprites.PlanetsFg[m.rng.Intn(len(sprites.PlanetsFg))], planetStyleAlt, m.height/2, m.height-10)
	m.items = append(m.items, fgPlanet)

	station := m.makeEntityRange(sprites.Stations[m.rng.Intn(len(sprites.Stations))], stationStyle, m.height/6, m.height/2)
	m.items = append(m.items, station)

	sat := m.makeEntityRange(sprites.Satellites[m.rng.Intn(len(sprites.Satellites))], satelliteStyle, 1, m.height/3)
	m.items = append(m.items, sat)

	meteorCount := clampInt(m.width/40, 2, 5)
	for i := 0; i < meteorCount; i++ {
		met := m.makeEntity(sprites.Meteors[m.rng.Intn(len(sprites.Meteors))], meteorStyle)
		m.items = append(m.items, met)
	}

	com := m.makeEntity(sprites.Comets[m.rng.Intn(len(sprites.Comets))], cometStyle)
	m.items = append(m.items, com)

	ufo := m.makeEntityRange(sprites.UFOs[m.rng.Intn(len(sprites.UFOs))], ufoStyle, m.height/3, m.height/2)
	m.items = append(m.items, ufo)
}

func (m *model) makeStar(art []string, style lipgloss.Style, vx, vy float64) sprites.Entity {
	e := sprites.Entity{Art: art, VX: vx, VY: vy, Style: style}
	e.X, e.Y = m.randPos(e.Width(), e.Height())
	return e
}

func (m *model) makeEntity(def sprites.SpriteDef, style lipgloss.Style) sprites.Entity {
	art := def.Art
	if len(def.Frames) > 0 {
		art = def.Frames[0]
	}
	e := sprites.Entity{Art: art, Frames: def.Frames, VX: def.DefaultVX, VY: def.DefaultVY, Layer: def.Layer, Style: style, UpdateFunc: def.UpdateFunc}
	e.X, e.Y = m.randPos(e.Width(), e.Height())
	return e
}

func (m *model) makeEntityRange(def sprites.SpriteDef, style lipgloss.Style, minY, maxY int) sprites.Entity {
	art := def.Art
	if len(def.Frames) > 0 {
		art = def.Frames[0]
	}
	e := sprites.Entity{Art: art, Frames: def.Frames, VX: def.DefaultVX, VY: def.DefaultVY, Layer: def.Layer, Style: style, UpdateFunc: def.UpdateFunc}
	if def.IsVertical {
		w := e.Width()
		maxX := m.width - w
		if maxX < 0 {
			maxX = 0
		}
		e.X = m.randRange(0, float64(maxX))
		e.Y = float64(m.height)
	} else {
		e.X, e.Y = m.randPosRange(e.Width(), e.Height(), minY, maxY)
	}
	return e
}

func (m *model) randPos(w, h int) (float64, float64) {
	maxX := m.width - w
	if maxX < 0 {
		maxX = 0
	}
	maxY := m.height - h
	if maxY < 0 {
		maxY = 0
	}
	return m.randRange(0, float64(maxX)), m.randRange(0, float64(maxY))
}

func (m *model) randPosRange(w, h int, minY, maxY int) (float64, float64) {
	maxX := m.width - w
	if maxX < 0 {
		maxX = 0
	}

	if minY < 0 {
		minY = 0
	}
	if maxY < minY {
		maxY = minY
	}

	maxYAdjusted := maxY - h
	if maxYAdjusted < minY {
		maxYAdjusted = minY
	}

	return m.randRange(0, float64(maxX)), m.randRange(float64(minY), float64(maxYAdjusted))
}

func (m *model) randRange(min, max float64) float64 {
	if max <= min {
		return min
	}
	return min + m.rng.Float64()*(max-min)
}

func clampInt(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func starDimStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#454552")).Background(lipgloss.Color("#080811"))
}

func starBrightStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#8A8A9E")).Background(lipgloss.Color("#080811"))
}

func main() {
	p := tea.NewProgram(newModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("astron error: %v\n", err)
	}
}
