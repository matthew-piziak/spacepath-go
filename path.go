package spacepath

import (
	"container/heap"
)

type Node interface {
	Neighbors() []Edge
	Heuristic(goal Node) float64
	Success(goal Node) bool
}

type Action string

type Edge struct {
	dest   Node
	action Action
	score  float64
}

func AStar(start Node, goal Node) []Edge {
	seen := make(map[Node]bool)
	openHeap := make(PriorityQueue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[Node]Edge)
	gScore := make(map[Node]float64)
	fScore := make(map[Node]float64)
	gScore[start] = 0
	fScore[start] = gScore[start] + start.Heuristic(goal)
	heap.Push(&openHeap, &Item{node: start, priority: fScore[start]})
	seen[start] = true
	for {
		node := heap.Pop(&openHeap).(*Item).node
		if node.Success(goal) {
			return reconstructPath(cameFrom, node)
		}
		for _, edge := range node.Neighbors() {
			adj := edge.dest
			action := edge.action
			if seen[adj] {
				continue
			}
			seen[adj] = true
			// adjacency cost is based on a constant step
			gScore[adj] = gScore[node] + 1
			hScore := adj.Heuristic(goal)
			fScore[adj] = gScore[adj] + hScore
			heap.Push(&openHeap, &Item{node: adj, priority: fScore[adj]})
			// reverse the edge for reconstruction
			cameFrom[adj] = Edge{
				dest:   node,
				action: action,
				score:  fScore[adj],
			}
		}
	}
}

func reconstructPath(cameFrom map[Node]Edge, node Node) []Edge {
	if edge, ok := cameFrom[node]; ok {
		return append(reconstructPath(cameFrom, edge.dest), edge)
	}
	return make([]Edge, 0)
}
