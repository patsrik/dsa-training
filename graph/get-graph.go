package graph

// GetGraph graph represented as map of string of array of strings
// key is vertex / node of graph
// value is list of neighbours from each node
func GetGraph() map[string][]string {
	newGraph := map[string][]string{}
	newGraph["a"] = []string{"b", "c"}
	newGraph["b"] = []string{"d"}
	newGraph["d"] = []string{"f"}
	newGraph["c"] = []string{"e"}
	newGraph["f"] = []string{}
	newGraph["e"] = []string{}

	return newGraph
}

func GetGraphForHasPath() map[string][]string {
	newGraph := map[string][]string{}
	newGraph["f"] = []string{"g", "i"}
	newGraph["g"] = []string{"h"}
	newGraph["h"] = []string{}
	newGraph["i"] = []string{"g", "k"}
	newGraph["j"] = []string{"i"}
	newGraph["k"] = []string{}

	return newGraph
}

func GetUndirectedGraphForHasPath() map[string][]string {
	newGraph := map[string][]string{}
	newGraph["0"] = []string{"7", "8", "4", "2"}
	newGraph["6"] = []string{"1", "5"}
	newGraph["2"] = []string{"0"}
	newGraph["5"] = []string{"8", "6", "3"}
	newGraph["4"] = []string{"7", "0"}
	newGraph["1"] = []string{"3"}
	newGraph["3"] = []string{"5", "1"}
	newGraph["7"] = []string{"0", "4"}
	newGraph["8"] = []string{"5", "0"}

	return newGraph
}
