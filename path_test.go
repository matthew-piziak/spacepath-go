package spacepath

import (
	"math"
	"testing"
)

type GridNode struct {
	x int8
	y int8
}

const (
	UP    Action = "up"
	DOWN  Action = "down"
	LEFT  Action = "left"
	RIGHT Action = "right"
)

// orthogonal movement
func (node GridNode) Neighbors() []Edge {
	return []Edge{
		Edge{
			dest:   GridNode{x: node.x + 1, y: node.y},
			action: "right",
		},
		Edge{
			dest:   GridNode{x: node.x - 1, y: node.y},
			action: "left",
		},
		Edge{
			dest:   GridNode{x: node.x, y: node.y + 1},
			action: "up",
		},
		Edge{
			dest:   GridNode{x: node.x, y: node.y - 1},
			action: "down",
		}}
}

// euclidean norm
func (node GridNode) Heuristic(goal Node) float64 {
	gridGoal := goal.(GridNode)
	return math.Hypot(float64(gridGoal.x-node.x), float64(gridGoal.y-node.y))
}

func (node GridNode) Success(goal Node) bool {
	return node == goal
}

func TestAStar(t *testing.T) {
	start := GridNode{x: 0, y: 0}
	goal := GridNode{x: 4, y: 4}
	path := AStar(start, goal)
	expectedLength := 8
	if len(path) != expectedLength {
		t.Errorf(
			"Expected length %d, was length %d.",
			expectedLength,
			len(path))
	}
}
