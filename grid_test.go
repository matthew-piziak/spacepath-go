package spacepathgo

import (
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {
	start := GridNode{X: 0, Y: 0}
	goal := GridNode{X: 64, Y: 64}
	path := AStar(start, goal)
	fmt.Printf("path length %d\n", len(path))
}
