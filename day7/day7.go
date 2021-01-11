package main

import (
	"advent2020/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("bags.txt")

	var graph Graph
	graph.vertices = make(map[string]*Vertex)

	for _, line := range data {
		split := strings.Split(line, "contain")
		outermostBag := strings.Split(split[0], " bags")[0]
		innerBagLines := strings.Split(split[1], ",")
		var innerBags []*Edge
		other := false
		for _, bagLine := range innerBagLines {
			innerBagCount, _ := strconv.Atoi(strings.Trim(bagLine, " ")[:1])
			innerBag := strings.TrimLeft(strings.Split(bagLine[3:], " bag")[0], " ")
			if innerBag == "other" {
				other = true
				continue
			}
			var vertex Vertex
			if _, ok := graph.vertices[innerBag]; !ok {
				vertex = Vertex{label: innerBag, edges: []*Edge{}}
			} else {
				vertex = *graph.vertices[innerBag]
			}
			innerBags = append(innerBags, &Edge{&vertex, innerBagCount})
		}
		if other {
			continue
		}
		if _, ok := graph.vertices[outermostBag]; !ok {
			graph.vertices[outermostBag] = &Vertex{label: outermostBag, edges: innerBags}
		} else {
			graph.vertices[outermostBag].edges = append(graph.vertices[outermostBag].edges, innerBags...)
		}
	}

	var shinyGoldBagCount int
	for _, vertex := range graph.vertices {
		found := graph.recursiveBagExistsSearch(vertex)
		if found {
			shinyGoldBagCount++
		}
	}
	fmt.Printf("bag colors that can contain at least one shiny gold bag: %v\n", shinyGoldBagCount)

	foundBags := graph.recursiveBagCountSearch(graph.vertices["shiny gold"])
	fmt.Printf("shiny gold bag contains %v other bags\n", foundBags)

}

func (graph *Graph) recursiveBagExistsSearch(v *Vertex) (found bool) {
	for _, edge := range v.edges {
		if edge.to.label == "shiny gold" {
			return true
		} else if bag, ok := graph.vertices[edge.to.label]; ok {
			found = graph.recursiveBagExistsSearch(bag)
			if found {
				break
			}
		}
	}
	return found
}

func (graph *Graph) recursiveBagCountSearch(v *Vertex) (foundBags int) {
	for _, edge := range v.edges {
		foundBags += edge.weight
		if bag, ok := graph.vertices[edge.to.label]; ok {
			count := edge.weight * graph.recursiveBagCountSearch(bag)
			foundBags += count
		}
	}
	return
}

// Edge points to a Vertex and captures the associated weight
type Edge struct {
	to     *Vertex
	weight int
}

// Vertex represents a Vertex in a Graph data structure, containing a slice of Edges
type Vertex struct {
	label string
	edges []*Edge
}

func (v *Vertex) addEdge(edge Edge) {
	v.edges = append(v.edges, &edge)
}

// Graph represents a weighted Graph data structure, containing a map of Vertices
type Graph struct {
	vertices map[string]*Vertex
}
