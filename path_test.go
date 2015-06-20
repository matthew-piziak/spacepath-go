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

func (node GridNode) Neighbors() []Edge {
	return []Edge{
		Edge{
			dest:   GridNode{node.x + 1, node.y},
			action: "right",
		},
		Edge{
			dest:   GridNode{node.x - 1, node.y},
			action: "left",
		},
		Edge{
			dest:   GridNode{node.x, node.y + 1},
			action: "up",
		},
		Edge{
			dest:   GridNode{node.x, node.y - 1},
			action: "down",
		}}
}

func (node GridNode) Heuristic(goal Node) float64 {
	gridGoal := goal.(GridNode)
	return math.Sqrt(
		math.Pow(float64(gridGoal.x)-float64(node.x), 2) +
			math.Pow(float64(gridGoal.y)-float64(node.y), 2))
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

func BenchmarkAStar(b *testing.B) {
	start := GridNode{x: 0, y: 0}
	goal := GridNode{x: 4, y: 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AStar(start, goal)
	}
}
