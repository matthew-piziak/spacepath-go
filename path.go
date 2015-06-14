// Generalized pathing functions
package spacepath

import (
	"container/heap"
	"math"
)

type Node struct {
	x     int8
	y     int8
	vx    int8
	vy    int8
	angle uint8
}

func min(open map[Node]float64) Node {
	waterline := math.Inf(1)
	best := Node{x: 0, y: 0, vx: 0, vy: 0, angle: 0}
	for k, v := range open {
		if v < waterline {
			waterline = v
			best = k
		}
	}
	return best
}

func reconstructPath(cameFrom map[Node]Node, node Node) []Node {
	if _, ok := cameFrom[node]; ok {
		path := append(reconstructPath(cameFrom, cameFrom[node]), node)
		return path
	}
	return []Node{node}
}

// A* pathing algorithm
func AStar(
	start Node,
	goal Node,
	adjacent func(Node) []Node,
	heuristic func(Node, Node) float64,
	success func(Node, Node) bool) []Node {
	seen := make(map[Node]bool)
	openHeap := make(PriorityQueue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[Node]Node)
	gScore := make(map[Node]float64)
	fScore := make(map[Node]float64)
	gScore[start] = 0
	fScore[start] = gScore[start] + heuristic(start, goal)
	heap.Push(&openHeap, &Item{value: start, priority: fScore[start]})
	seen[start] = true
	for {
		node := heap.Pop(&openHeap).(*Item).value
		if success(node, goal) {
			return reconstructPath(cameFrom, node)
		}
		for _, adj := range adjacent(node) {
			if seen[adj] {
				continue
			}
			seen[adj] = true
			cameFrom[adj] = node
			// adjacency is based on a constant time step
			gScore[adj] = gScore[node] + 1
			hScore := heuristic(adj, goal)
			fScore[adj] = gScore[adj] + hScore
			heap.Push(&openHeap, &Item{value: adj, priority: fScore[adj]})
		}
	}
}
