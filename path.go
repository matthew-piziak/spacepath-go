// Generalized pathing functions
package spacepath

import (
	"container/heap"
)

type Node struct {
	x     int8
	y     int8
	vx    int8
	vy    int8
	angle uint8
}

type Action string

const (
	UP    Action = "up"
	DOWN  Action = "down"
	LEFT  Action = "left"
	RIGHT Action = "right"
)

type Edge struct {
	dest   Node
	action Action
}

func reconstructPath(cameFrom map[Node]Edge, node Node) []Action {
	if edge, ok := cameFrom[node]; ok {
		return append(reconstructPath(cameFrom, edge.dest), edge.action)
	}
	return make([]Action, 0)
}

// A* pathing algorithm
func AStar(
	start Node,
	goal Node,
	adjacent func(Node) []Edge,
	heuristic func(Node, Node) float64,
	success func(Node, Node) bool) []Action {
	seen := make(map[Node]bool)
	openHeap := make(PriorityQueue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[Node]Edge)
	gScore := make(map[Node]float64)
	fScore := make(map[Node]float64)
	gScore[start] = 0
	fScore[start] = gScore[start] + heuristic(start, goal)
	heap.Push(&openHeap, &Item{node: start, priority: fScore[start]})
	seen[start] = true
	for {
		node := heap.Pop(&openHeap).(*Item).node
		if success(node, goal) {
			return reconstructPath(cameFrom, node)
		}
		for _, edge := range adjacent(node) {
			adj := edge.dest
			action := edge.action
			if seen[adj] {
				continue
			}
			seen[adj] = true
			// reverse the edge for reconstruction
			cameFrom[adj] = Edge{dest: node, action: action}
			// adjacency is based on a constant time step
			gScore[adj] = gScore[node] + 1
			hScore := heuristic(adj, goal)
			fScore[adj] = gScore[adj] + hScore
			heap.Push(&openHeap, &Item{node: adj, priority: fScore[adj]})
		}
	}
}
