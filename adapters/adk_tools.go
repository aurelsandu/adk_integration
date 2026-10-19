package adapters

import (
	"fmt"
	"strings"

	"github.com/aurelsandu/adk-integration/domains"
	"github.com/aurelsandu/adk-integration/search"
)

// SolveMazeTool acts as the bridge between the LLM agent and the A* deterministic solver.
func SolveMazeTool(mazeLines []string) (string, error) {
	var start, end domains.Point
	foundStart, foundEnd := false, false

	for r, row := range mazeLines {
		fields := strings.Fields(row)
		for c, val := range fields {
			if val == "S" {
				start = domains.Point{X: r, Y: c}
				foundStart = true
			} else if val == "E" {
				end = domains.Point{X: r, Y: c}
				foundEnd = true
			}
		}
	}

	if !foundStart || !foundEnd {
		return "", fmt.Errorf("invalid maze: missing Start (S) or Exit (E)")
	}

	startState := domains.NewMazeState(mazeLines, start, end)
	goalState := domains.NewMazeState(mazeLines, end, end)

	path, err := search.FindPath(startState, goalState)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Path found successfully with %d steps:\n", len(path)-1))
	for i, state := range path {
		s := state.(*domains.MazeState)
		sb.WriteString(fmt.Sprintf("Step %d: (%d, %d)\n", i, s.Pos.X, s.Pos.Y))
	}

	return sb.String(), nil
}
