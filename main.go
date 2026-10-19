package main

import (
	"fmt"
	"strings"
	"time"

//	"github.com/aurelsandu/adk-integration/adapters"
	"github.com/aurelsandu/adk-integration/domains"
	"github.com/aurelsandu/adk-integration/search"
)

// Clear terminal screen using ANSI escape codes
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// AnimatePath prints the maze and animates the agent moving along the path
func AnimatePath(mazeLines []string, path []search.State) {
	// Parse grid
	grid := make([][]string, len(mazeLines))
	for r, row := range mazeLines {
		grid[r] = strings.Fields(row)
	}

	// Extract coordinate steps
	var steps []domains.Point
	for _, s := range path {
		ms := s.(*domains.MazeState)
		steps = append(steps, ms.Pos)
	}

	// Animate step by step
	for i, step := range steps {
		clearScreen()
		fmt.Printf("=== Symbolic AI: A* Maze Animation (Step %d/%d) ===\n\n", i+1, len(steps))

		// Render current frame
		for r := range grid {
			for c := range grid[r] {
				currentPos := domains.Point{X: r, Y: c}
				
				// Check if current step is the agent's position
				if currentPos == step {
					fmt.Print("🤖 ") // Robot/Agent icon
				} else {
					// Check if it's part of past path
					isVisitedPath := false
					for j := 0; j < i; j++ {
						if steps[j] == currentPos {
							isVisitedPath = true
							break
						}
					}
					
					if isVisitedPath {
						fmt.Print("*  ") // Traveled path
					} else {
						fmt.Print(grid[r][c] + "  ")
					}
				}
			}
			fmt.Println()
		}

		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("\n✨ Goal reached successfully!")
}

func main() {
	maze := []string{
		"S . . # . . .",
		"# # . # . # .",
		". . . . . # .",
		". # # # . # .",
		". . . # . . E",
	}

	startState := domains.NewMazeState(maze, domains.Point{X: 0, Y: 0}, domains.Point{X: 4, Y: 6})
	goalState := domains.NewMazeState(maze, domains.Point{X: 4, Y: 6}, domains.Point{X: 4, Y: 6})

	path, err := search.FindPath(startState, goalState)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	AnimatePath(maze, path)
}
