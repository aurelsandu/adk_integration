package search

import (
	"container/heap"
	"fmt"
)

// State defines the interface that any problem state must implement
type State interface {
	ID() string
	Cost() float64
	Heuristic(goal State) float64
	Neighbors() []State
	IsGoal(goal State) bool
}

// Node represents a search node in the A* tree
type Node struct {
	State   State
	G       float64 // Cost from start
	H       float64 // Estimated cost to goal
	F       float64 // G + H
	Parent  *Node
	index   int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].F < pq[j].F }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// FindPath executes the generic A* search algorithm
func FindPath(start, goal State) ([]State, error) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	startNode := &Node{
		State: start,
		G:     0,
		H:     start.Heuristic(goal),
		F:     start.Heuristic(goal),
	}
	heap.Push(pq, startNode)

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		if current.State.IsGoal(goal) {
			var path []State
			curr := current
			for curr != nil {
				path = append([]State{curr.State}, path...)
				curr = curr.Parent
			}
			return path, nil
		}

		visited[current.State.ID()] = true

		for _, neighbor := range current.State.Neighbors() {
			if visited[neighbor.ID()] {
				continue
			}

			gCost := current.G + neighbor.Cost()
			hCost := neighbor.Heuristic(goal)
			
			neighborNode := &Node{
				State:   neighbor,
				G:       gCost,
				H:       hCost,
				F:       gCost + hCost,
				Parent:  current,
			}

			heap.Push(pq, neighborNode)
		}
	}

	return nil, fmt.Errorf("no path found to target state")
}
