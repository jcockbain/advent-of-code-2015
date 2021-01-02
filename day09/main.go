package main

import (
	input "aoc2015/inpututils"

	"fmt"
	"math"
	"regexp"
	"strconv"
)

var (
	re      = regexp.MustCompile(`([A-Za-z]+) to ([A-za-z]+) = (\d+)`)
	minPath = math.MaxInt32
	maxPath = 0
)

type Candidate struct {
	node Node
	edge Edge
}

type Node struct {
	loc   string
	edges []*Edge
}

type Edge struct {
	dest   *Node
	weight int
}

type Graph []*Node

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	nodes := buildGraph(lines)
	nodesLength := len(nodes)
	for _, node := range nodes {
		a := make([]Node, nodesLength)
		a[0] = *node
		backtrack(a, 0, nodes, 0)
	}
	return minPath
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	nodes := buildGraph(lines)
	nodesLength := len(nodes)
	for _, node := range nodes {
		a := make([]Node, nodesLength)
		a[0] = *node
		backtrack(a, 0, nodes, 0)
	}
	return maxPath
}

func backtrack(a []Node, k int, graph Graph, total int) {
	if isSolution(a, k, graph) {
		processSolution(a, graph, total)
	} else {
		k = k + 1
		c := constructCandidates(a, k, graph)
		for _, candidate := range c {
			a[k] = candidate.node
			backtrack(a, k, graph, total+candidate.edge.weight)
			a[k] = Node{}
		}
	}
}

func isSolution(a []Node, k int, graph Graph) bool {
	if k == len(graph)-1 {
		return true
	}
	return false
}

func processSolution(a []Node, graph Graph, total int) {
	if total < minPath {
		minPath = total
	}
	if total > maxPath {
		maxPath = total
	}
}

func constructCandidates(a []Node, k int, graph Graph) (c []Candidate) {
	currentNode := a[k-1]
	for _, edge := range currentNode.edges {
		destNode := edge.dest
		if !contains(a, *destNode) {
			c = append(c, Candidate{
				*destNode,
				*edge,
			})
		}
	}
	return c
}

func contains(slice []Node, node Node) bool {
	for _, comp := range slice {
		if comp.loc == node.loc {
			return true
		}
	}
	return false
}

func buildGraph(lines []string) Graph {
	nodes := map[string]*Node{}
	for _, line := range lines {
		groups := re.FindStringSubmatch(line)
		start, dest, weight := groups[1], groups[2], toInt(groups[3])
		if _, ok := nodes[start]; !ok {
			nodes[start] = &Node{
				start,
				[]*Edge{},
			}
		}
		if _, ok := nodes[dest]; !ok {
			nodes[dest] = &Node{
				dest,
				[]*Edge{},
			}
		}
		newEdgeToStart := Edge{
			nodes[start],
			weight,
		}
		newEdgeToDest := Edge{
			nodes[dest],
			weight,
		}
		startNode := nodes[start]
		startNode.edges = append((*startNode).edges, &newEdgeToDest)
		destNode := nodes[dest]
		destNode.edges = append((*destNode).edges, &newEdgeToStart)
	}

	nodeList := []*Node{}
	for _, node := range nodes {
		nodeList = append(nodeList, node)
	}
	return nodeList
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
