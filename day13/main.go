package main

import (
	input "aoc2015/inpututils"
	"fmt"

	"math"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`(.*) would (gain|lose) (\d+) happiness units by sitting next to (.*).`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

type person struct {
	name      string
	relations []relation
}

type relation struct {
	subject   string
	happiness int
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	graph := buildGraph(lines)
	maxHappiness, _ := calculateMaxOrdering(graph)
	return maxHappiness
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	graph := buildGraph(lines)
	graph = addYou(graph)
	maxHappiness, _ := calculateMaxOrdering(graph)
	return maxHappiness
}

func buildGraph(lines []string) map[string]*person {
	people := map[string]*person{}
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		personName, positive, units, subject := parts[1], parts[2], toInt(parts[3]), parts[4]
		if positive == "lose" {
			units *= -1
		}
		if _, ok := people[personName]; !ok {
			people[personName] = &person{
				personName,
				[]relation{},
			}
		}
		newRelation := relation{
			subject,
			units,
		}
		p := people[personName]
		p.relations = append((*p).relations, newRelation)
	}
	return people
}

func addYou(graph map[string]*person) map[string]*person {
	newGraph := map[string]*person{}
	you := person{
		"You",
		[]relation{},
	}
	for _, p := range graph {
		you.relations = append(you.relations, relation{
			(*p).name,
			0,
		})
		newRelations := append((*p).relations, relation{
			"You",
			0,
		})
		newGraph[(*p).name] = &person{
			(*p).name,
			newRelations,
		}
	}
	newGraph["You"] = &you
	return newGraph
}

func calculateMaxOrdering(graph map[string]*person) (int, []string) {
	maximumHappiness := math.MinInt32
	maxOrdering := []string{""}

	processSolution := func(a []string, total int, graph map[string]*person) {
		total += getRelation(a[0], a[len(a)-1], graph)
		if total > maximumHappiness {
			maximumHappiness = total
			maxOrdering = a
		}
	}

	var backtrack func(a []string, k int, graph map[string]*person, total int)

	backtrack = func(a []string, k int, graph map[string]*person, total int) {
		if isSolution(k, graph) {
			processSolution(a, total, graph)
		} else {
			k = k + 1
			c := constructCandidates(a, k, graph)
			for _, candidate := range c {
				a[k] = candidate.subject
				backtrack(a, k, graph, total+getRelation(a[k-1], a[k], graph))
				a[k] = ""
			}
		}
	}

	for _, node := range graph {
		a := make([]string, len(graph))
		a[0] = (*node).name
		backtrack(a, 0, graph, 0)
	}

	return maximumHappiness, maxOrdering
}

func getRelation(a string, b string, graph map[string]*person) int {
	personA := (*graph[a])
	personB := (*graph[b])
	total := 0

	for _, r := range personA.relations {
		if r.subject == b {
			total += r.happiness
		}
	}
	for _, r := range personB.relations {
		if r.subject == a {
			total += r.happiness
		}
	}
	return total
}

func isSolution(k int, graph map[string]*person) bool {
	return k == len(graph)-1
}

func constructCandidates(a []string, k int, graph map[string]*person) (c []relation) {
	current := a[k-1]
	personNode := (*graph[current])
	for _, r := range personNode.relations {
		if !contains(a, r.subject) {
			c = append(c, r)
		}
	}
	return c
}

func contains(a []string, s string) bool {
	for _, comp := range a {
		if comp == s {
			return true
		}
	}
	return false
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
