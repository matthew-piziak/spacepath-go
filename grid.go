package spacepathgo

import (
	"math"
)

type GridNode struct {
	X int16
	Y int16
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
			Dest:   GridNode{X: node.X + 1, Y: node.Y},
			Action: "right",
		},
		Edge{
			Dest:   GridNode{X: node.X - 1, Y: node.Y},
			Action: "left",
		},
		Edge{
			Dest:   GridNode{X: node.X, Y: node.Y + 1},
			Action: "up",
		},
		Edge{
			Dest:   GridNode{X: node.X, Y: node.Y - 1},
			Action: "down",
		}}
}

// euclidean norm
func (node GridNode) Heuristic(goal Node) float64 {
	gridGoal := goal.(GridNode)
	return math.Hypot(float64(gridGoal.X-node.X), float64(gridGoal.Y-node.Y))
}

func (node GridNode) Success(goal Node) bool {
	return node == goal
}
