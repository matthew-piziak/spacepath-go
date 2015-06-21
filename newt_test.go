package spacepath

import (
	"fmt"
	"testing"
)

func TestNewt(t *testing.T) {
	start := NewtNode{x: 0, y: 0, vx: 0, vy: 0, angle: 1}
	goal := NewtNode{x: 1024, y: 1024, vx: 0, vy: 0, angle: 1}
	path := AStar(start, goal)
	for _, node := range path {
		fmt.Printf("%v\n", node)
	}
	fmt.Printf("path length %d\n", len(path))
}
