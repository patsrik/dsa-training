package main

import (
	"dsa-training/graph"
	"fmt"
)

/*
Given a graph, and 2 nodes as input, i.e., source and destination,
determine if a path exists from the source to destination
The input is a undirected graph, i.e, you can traverse in both directions between
the given 2 nodes for an edge

Leetcode  :https://leetcode.com/problems/find-if-path-exists-in-graph/submissions/1825284004/
*/

func main() {
	givenGraph := graph.GetUndirectedGraphForHasPath()
	source := "2"
	dest := "3"
	visited := map[string]bool{}
	hasPath := hasPathFunc(source, dest, givenGraph, visited)
	fmt.Printf("has path from %s to %s %v", source, dest, hasPath)
}

func hasPathFunc(s, d string, graph map[string][]string, visited map[string]bool) bool {
	if _, ok := visited[s]; ok {
		return false
	}

	if s == d {
		return true
	}

	visited[s] = true

	for _, neighbour := range graph[s] {
		if hasPathFunc(neighbour, d, graph, visited) {
			return true
		}
	}

	return false
}
