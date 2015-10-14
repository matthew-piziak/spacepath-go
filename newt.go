package spacepathgo

import (
	"math"
)

type NewtNode struct {
	X  int
	Y  int
	ΔX int
	ΔY int
	Θ  int
}

const (
	CRUISE_STRAIGHT Action = "cruise straight"
	CRUISE_LEFT     Action = "cruise left"
	CRUISE_RIGHT    Action = "cruise right"
	BURN_STRAIGHT   Action = "burn straight"
	BURN_LEFT       Action = "burn left"
	BURN_RIGHT      Action = "burn right"
)

func (node NewtNode) Neighbors() []Edge {
	X := node.X + node.ΔX
	Y := node.Y + node.ΔY
	ΔX := node.ΔX
	ΔY := node.ΔY
	angle := node.Θ
	sin := map[int]int{0: 0, 1: 1, 2: 1, 3: 1, 4: 0, 5: -1, 6: -1, 7: -1}
	cos := map[int]int{0: 1, 1: 1, 2: 0, 3: -1, 4: -1, 5: -1, 6: 0, 7: 1}
	ΔΔX := cos[angle]
	ΔΔY := sin[angle]
	left_angle := (angle - 1) % 8
	right_angle := (angle + 1) % 8
	return []Edge{
		Edge{
			Dest:   NewtNode{X: X, Y: Y, ΔX: ΔX, ΔY: ΔY, Θ: angle},
			Action: "cruise straight",
		},
		Edge{
			Dest:   NewtNode{X: X, Y: Y, ΔX: ΔX, ΔY: ΔY, Θ: left_angle},
			Action: "Cruise left",
		},
		Edge{
			Dest:   NewtNode{X: X, Y: Y, ΔX: ΔX, ΔY: ΔY, Θ: right_angle},
			Action: "Cruise right",
		},
		Edge{
			Dest: NewtNode{
				X:  X,
				Y:  Y,
				ΔX: ΔX + ΔΔX,
				ΔY: ΔY + ΔΔY,
				Θ:  angle,
			},
			Action: "burn straight",
		},
		Edge{
			Dest: NewtNode{
				X:  X,
				Y:  Y,
				ΔX: ΔX + ΔΔX,
				ΔY: ΔY + ΔΔY,
				Θ:  left_angle,
			},
			Action: "burn left",
		},
		Edge{
			Dest: NewtNode{
				X:  X,
				Y:  Y,
				ΔX: ΔX + ΔΔX,
				ΔY: ΔY + ΔΔY,
				Θ:  right_angle},
			Action: "burn right",
		},
	}
}

func (node NewtNode) Heuristic(goal Node) float64 {
	newtGoal := goal.(NewtNode)
	hMaX := math.MaxFloat64
	boundX := (newtGoal.X * 11) / 10
	boundY := (newtGoal.Y * 11) / 10
	if outsideArena(node, boundX, boundY) {
		return hMaX
	}
	if leavingArena(node, boundX, boundY) {
		return hMaX
	}
	hX := heuristic(
		float64(node.X),
		float64(node.ΔX),
		float64(newtGoal.X),
		float64(newtGoal.ΔX))
	hY := heuristic(
		float64(node.Y),
		float64(node.ΔY),
		float64(newtGoal.Y),
		float64(newtGoal.ΔY))
	return 1.02 * (hX + hY)
}

func heuristic(np float64, nv float64, gp float64, gv float64) float64 {
	return ((-1 * nv) + (-2 * gv) + math.Sqrt((7*gv*gv)+(2*nv*nv)+(8*math.Abs(gp-np)))) / 2
}

func (node NewtNode) Success(goal Node) bool {
	newtGoal := goal.(NewtNode)
	speed := math.Abs(float64(node.ΔX-newtGoal.ΔX)) +
		math.Abs(float64(node.ΔX-newtGoal.ΔY))
	distance := math.Sqrt(math.Pow(float64(newtGoal.X)-float64(node.X), 2) +
		math.Pow(float64(newtGoal.Y)-float64(node.Y), 2))
	return speed == 0 && distance < 1
}

func outsideArena(node NewtNode, boundX int, boundY int) bool {
	if node.X < 0 || node.Y < 0 {
		return true
	}
	if node.X > boundX || node.Y > boundY {
		return true
	}
	return false
}

func leavingArena(node NewtNode, boundX int, boundY int) bool {
	brakingTimeX := math.Abs(float64(node.ΔX)) // acceleration
	brakingTimeY := math.Abs(float64(node.ΔY)) // acceleration
	vComponentX := math.Abs(float64(node.ΔX)) * brakingTimeX
	vComponentY := math.Abs(float64(node.ΔY)) * brakingTimeY
	aComponentX := (2 * brakingTimeX) / 2
	aComponentY := (2 * brakingTimeY) / 2
	brakingDistX := vComponentX + aComponentX
	brakingDistY := vComponentY + aComponentY
	if node.ΔX > 0 {
		if brakingDistX > float64(boundX-node.X) {
			return true
		}
	}
	if node.ΔY > 0 {
		if brakingDistY > float64(boundY-node.Y) {
			return true
		}
	}
	if node.ΔX < -1 {
		if brakingDistX > float64(node.X) {
			return true
		}
	}
	if node.ΔY < -1 {
		if brakingDistY > float64(node.Y) {
			return true
		}
	}
	return false
}
