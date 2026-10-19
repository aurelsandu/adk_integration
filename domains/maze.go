package domains

import (
	"fmt"
	"math"
	"strings"

	"github.com/aurelsandu/adk-integration/search"
)

type Point struct {
	X, Y int
}

// MazeState implements the search.State interface
type MazeState struct {
	Pos      Point
	Grid     [][]string
	Goal     Point
	stepCost float64
}

func NewMazeState(grid []string, start, goal Point) *MazeState {
	parsedGrid := make([][]string, len(grid))
	for r, row := range grid {
		parsedGrid[r] = strings.Fields(row)
	}
	return &MazeState{
		Pos:      start,
		Grid:     parsedGrid,
		Goal:     goal,
		stepCost: 0.0,
	}
}

func (m *MazeState) ID() string {
	return fmt.Sprintf("%d,%d", m.Pos.X, m.Pos.Y)
}

func (m *MazeState) Cost() float64 {
	return 1.0
}

func (m *MazeState) Heuristic(goal search.State) float64 {
	g := goal.(*MazeState)
	dx := math.Abs(float64(m.Pos.X - g.Pos.X))
	dy := math.Abs(float64(m.Pos.Y - g.Pos.Y))
	return dx + dy
}

func (m *MazeState) IsGoal(goal search.State) bool {
	g := goal.(*MazeState)
	return m.Pos == g.Pos
}

func (m *MazeState) Neighbors() []search.State {
	rows := len(m.Grid)
	cols := len(m.Grid[0])
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var neighbors []search.State
	for _, d := range directions {
		nextPos := Point{X: m.Pos.X + d.X, Y: m.Pos.Y + d.Y}

		// Check boundaries
		if nextPos.X < 0 || nextPos.X >= rows || nextPos.Y < 0 || nextPos.Y >= cols {
			continue
		}

		// Check walls
		if m.Grid[nextPos.X][nextPos.Y] == "#" {
			continue
		}

		nextState := &MazeState{
			Pos:      nextPos,
			Grid:     m.Grid,
			Goal:     m.Goal,
			stepCost: 1.0,
		}
		neighbors = append(neighbors, nextState)
	}
	return neighbors
}
