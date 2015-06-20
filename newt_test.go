package spacepath

import (
	"fmt"
	"math"
	"testing"
)

type NewtNode struct {
	x     int16
	y     int16
	vx    int16
	vy    int16
	angle int8
}

const (
	CRUISE_STRAIGHT Action = "cruise straight"
	CRUISE_LEFT     Action = "cruise left"
	CRUISE_RIGHT    Action = "cruise right"
	BURN_STRAIGHT   Action = "burn straight"
	BURN_LEFT       Action = "burn left"
	BURN_RIGHT      Action = "burn right"
)

const ACCELERATION int8 = 2

func (node NewtNode) Neighbors() []Edge {
	x := node.x + node.vx
	y := node.y + node.vy
	vx := node.vx
	vy := node.vy
	angle := node.angle
	sin := map[int8]int16{0: 0, 1: 1, 2: 1, 3: 1, 4: 0, 5: -1, 6: -1, 7: -1}
	cos := map[int8]int16{0: 1, 1: 1, 2: 0, 3: -1, 4: -1, 5: -1, 6: 0, 7: 1}
	Δvx := cos[angle]
	Δvy := sin[angle]
	left_angle := (angle - 1) % 8
	right_angle := (angle + 1) % 8
	return []Edge{
		Edge{
			dest:   NewtNode{x: x, y: y, vx: vx, vy: vy, angle: angle},
			action: "cruise straight",
		},
		Edge{
			dest: NewtNode{
				x:     x,
				y:     y,
				vx:    vx,
				vy:    vy,
				angle: left_angle,
			},
			action: "cruise left",
		},
		Edge{
			dest: NewtNode{
				x:     x,
				y:     y,
				vx:    vx,
				vy:    vy,
				angle: right_angle,
			},
			action: "cruise right",
		},
		Edge{
			dest: NewtNode{
				x:     x,
				y:     y,
				vx:    vx + Δvx,
				vy:    vy + Δvy,
				angle: angle,
			},
			action: "burn straight",
		},
		Edge{
			dest: NewtNode{
				x:     x,
				y:     y,
				vx:    vx + Δvx,
				vy:    vy + Δvy,
				angle: left_angle,
			},
			action: "burn left",
		},
		Edge{
			dest: NewtNode{
				x:     x,
				y:     y,
				vx:    vx + Δvx,
				vy:    vy + Δvy,
				angle: right_angle,
			},
			action: "burn right",
		},
	}
}

func (node NewtNode) Heuristic(goal Node) float64 {
	newtGoal := goal.(NewtNode)
	hMax := math.MaxFloat64
	if outsideArena(node, 50, 50) {
		return hMax
	}
	if leavingArena(node, 50, 50) {
		return hMax
	}
	hx := (-float64(node.vx) + math.Sqrt((2 * math.Pow(float64(node.vx), 2))) + (float64(4*ACCELERATION) * math.Abs(float64(newtGoal.x-node.x)))) / 2
	hy := (-float64(node.vy) + math.Sqrt((2 * math.Pow(float64(node.vy), 2))) + (float64(4*ACCELERATION) * math.Abs(float64(newtGoal.y-node.y)))) / 2
	return 1.02 * (hx + hy)
}

func (node NewtNode) Success(goal Node) bool {
	newtGoal := goal.(NewtNode)
	speed := math.Abs(float64(node.vx-newtGoal.vx)) +
		math.Abs(float64(node.vx-newtGoal.vy))
	distance := math.Sqrt(math.Pow(float64(newtGoal.x)-float64(node.x), 2) +
		math.Pow(float64(newtGoal.y)-float64(node.y), 2))
	return speed == 0 && distance < 6
}

func outsideArena(node NewtNode, boundX int16, boundY int16) bool {
	if node.x < 0 || node.y < 0 {
		return false
	}
	if node.x > boundX || node.y > boundY {
		return false
	}
	return true
}

func leavingArena(node NewtNode, boundX int16, boundY int16) bool {
	brakingTimeX := math.Abs(float64(node.vx)) // acceleration
	brakingTimeY := math.Abs(float64(node.vy)) // acceleration
	vComponentX := math.Abs(float64(node.vx)) * brakingTimeX
	vComponentY := math.Abs(float64(node.vy)) * brakingTimeY
	aComponentX := (float64(ACCELERATION) * brakingTimeX) / 2
	aComponentY := (float64(ACCELERATION) * brakingTimeY) / 2
	brakingDistX := vComponentX + aComponentX
	brakingDistY := vComponentY + aComponentY
	if node.vx > 0 {
		if brakingDistX > float64(boundX-node.x) {
			return true
		}
	}
	if node.vy > 0 {
		if brakingDistY > float64(boundY-node.y) {
			return true
		}
	}
	if node.vx < -1 {
		if brakingDistX > float64(node.x) {
			return true
		}
	}
	if node.vy < -1 {
		if brakingDistY > float64(node.y) {
			return true
		}
	}
	return false
}

func TestNewt(t *testing.T) {
	start := NewtNode{x: 0, y: 0, vx: 0, vy: 0, angle: 1}
	goal := NewtNode{x: 9, y: 9, vx: 0, vy: 0, angle: 0}
	path := AStar(start, goal)
	for i, path := range path {
		fmt.Printf("%d: %s\n", i, path)
	}
}
