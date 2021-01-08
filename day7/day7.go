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
		found := graph.recursiveBagSearch(vertex)
		//fmt.Printf("bag: %v, count: %v\n", vertex.label, found)
		if found {
			shinyGoldBagCount++
		}
	}
	fmt.Printf("bag colors that can contain at least one shiny gold bag: %v\n", shinyGoldBagCount)
}

func (graph *Graph) recursiveBagSearch(v *Vertex) (found bool) {
	for _, edge := range v.edges {
		if edge.to.label == "shiny gold" {
			//fmt.Printf("\tvertex edge %v\n", v.label)
			return true
		} else if bag, ok := graph.vertices[edge.to.label]; ok {
			found = graph.recursiveBagSearch(bag)
			if found {
				break
			}
		}
	}
	return found
}

type Edge struct {
	to     *Vertex
	weight int
}

type Vertex struct {
	label string
	edges []*Edge
}

func (v *Vertex) addEdge(edge Edge) {
	v.edges = append(v.edges, &edge)
}

type Graph struct {
	vertices map[string]*Vertex
}
