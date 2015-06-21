package spacepath

import (
	"math"
)

type GridNode struct {
	x int16
	y int16
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
