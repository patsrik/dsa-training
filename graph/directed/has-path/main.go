package main

import (
	"dsa-training/graph"
	"fmt"
)

/*
Given a graph, and 2 nodes as input, i.e., source and destination,
determine if a path exists from the source to destination
The input is a directed graph, i.e, there is only one direction to travel between
any two given nodes
*/

func main() {
	givenGraph := graph.GetGraphForHasPath()
	source := "f"
	dest := "h"
	hasPath := hasPathFunc(source, dest, givenGraph)
	fmt.Printf("has path from %s to %s %v", source, dest, hasPath)
}

func hasPathFunc(s, d string, graph map[string][]string) bool {
	if s == d {
		return true
	}

	for _, neighbour := range graph[s] {
		if hasPathFunc(neighbour, d, graph) {
			return true
		}
	}

	return false
}
