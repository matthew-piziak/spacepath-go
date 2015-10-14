package spacepathgo

import (
	"fmt"
	"testing"
)

func TestNewt(t *testing.T) {
	start := NewtNode{X: 0, Y: 0, ΔX: 0, ΔY: 0, Θ: 1}
	goal := NewtNode{X: 512, Y: 512, ΔX: 0, ΔY: 0, Θ: 1}
	path := AStar(start, goal)
	fmt.Printf("path length %d\n", len(path))
}
