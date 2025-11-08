package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

// graph represented as map of string of array of strings
// key is vertex / node of graph
// value is list of neighbours from each node
func getGraph() map[string][]string {
	newGraph := map[string][]string{}
	newGraph["a"] = []string{"b", "c"}
	newGraph["b"] = []string{"d"}
	newGraph["d"] = []string{"f"}
	newGraph["c"] = []string{"e"}
	newGraph["f"] = []string{}
	newGraph["e"] = []string{}

	return newGraph
}

func main() {
	fmt.Print("in main \n")
	graph := getGraph()
	traverseInDFS(graph, "a")
}

func traverseInDFS(graph map[string][]string, startNode string) {
	fmt.Print("in traverse \n")
	stack := lane.NewStack()
	stack.Push(startNode)

	for stack.Size() > 0 {
		current := stack.Pop().(string)
		fmt.Println(current)
		for _, neighbour := range graph[current] {
			stack.Push(neighbour)
		}
	}
}
