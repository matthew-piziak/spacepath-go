package spacepath

import (
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {
	start := GridNode{x: 0, y: 0}
	goal := GridNode{x: 64, y: 64}
	path := AStar(start, goal)
	fmt.Printf("path length %d\n", len(path))
}
