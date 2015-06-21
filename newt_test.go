package spacepathgo

import (
	"fmt"
	"testing"
)

func TestNewt(t *testing.T) {
	start := NewtNode{x: 0, y: 0, vx: 0, vy: 0, angle: 1}
	goal := NewtNode{x: 512, y: 512, vx: 0, vy: 0, angle: 1}
	path := AStar(start, goal)
	fmt.Printf("path length %d\n", len(path))
}
