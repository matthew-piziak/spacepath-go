package spacepath

import (
	"math"
	"testing"
)

func manhattan(node Node) []Node {
	adjacencyList := []Node{
		Node{node.x + 1, node.y, node.vx, node.vy, node.angle},
		Node{node.x - 1, node.y, node.vx, node.vy, node.angle},
		Node{node.x, node.y + 1, node.vx, node.vy, node.angle},
		Node{node.x, node.y - 1, node.vx, node.vy, node.angle}}
	return adjacencyList
}

func euclideanDistance(Node Node, goal Node) float64 {
	return math.Sqrt(
		math.Pow(float64(goal.x)-float64(Node.x), 2) +
			math.Pow(float64(goal.y)-float64(Node.y), 2))
}

func nodeEqual(Node Node, goal Node) bool {
	return Node == goal
}

func TestAStar(t *testing.T) {
	start := Node{x: 0, y: 0, vx: 0, vy: 0, angle: 0}
	goal := Node{x: 8, y: 8, vx: 0, vy: 0, angle: 0}
	adjacent := manhattan
	heuristic := euclideanDistance
	success := nodeEqual
	path := AStar(start, goal, adjacent, heuristic, success)
	if len(path) != 17 {
		t.Errorf("Expected length 17, was length %d.", len(path))
	}
}

func BenchmarkAStar(b *testing.B) {
	start := Node{x: 0, y: 0, vx: 0, vy: 0, angle: 0}
	goal := Node{x: 8, y: 8, vx: 0, vy: 0, angle: 0}
	adjacent := manhattan
	heuristic := euclideanDistance
	success := nodeEqual
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AStar(start, goal, adjacent, heuristic, success)
	}
}
