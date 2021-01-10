package main

import (
	"fmt"
	"math"
	"strings"

	"regexp"
	"strconv"

	input "aoc2015/inpututils"
)

var (
	re1 = regexp.MustCompile(`([A-Za-z]+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	re2 = regexp.MustCompile(`([A-Za-z]+\s+\+\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	re3 = regexp.MustCompile(`(.*): (\d+)`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt", 100))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt", 100))
}

type resource struct {
	name                string
	cost, damage, armor int
}

type player struct {
	hitPoints, damage, armor int
}

func (p player) beats(boss player) bool {
	for p.hitPoints > 0 && boss.hitPoints > 0 {
		boss.hitPoints -= p.getDamageDoneOn(boss)
		p.hitPoints -= boss.getDamageDoneOn(p)
	}
	return boss.hitPoints <= 0
}

func (p player) getDamageDoneOn(p2 player) int {
	if p2.armor >= p.damage {
		return 1
	}
	return p.damage - p2.armor
}

func newPlayer(filename string) player {
	inp := input.ReadLines(filename)
	p := player{}
	for _, line := range inp {
		parts := re3.FindStringSubmatch(line)
		field, val := parts[1], parts[2]
		switch field {
		case "Hit Points":
			p.hitPoints = toInt(val)
		case "Damage":
			p.damage = toInt(val)
		case "Armor":
			p.armor = toInt(val)
		}
	}
	return p
}

func getResources() ([]resource, []resource, []resource) {
	items := input.ReadRaw("shop.txt")
	p := strings.Split(items, "\n\n")

	weaponsInput, armorInput, ringsInput := p[0], p[1], p[2]
	weapons, armors, rings := []resource{}, []resource{}, []resource{}

	for _, weapon := range strings.Split(weaponsInput, "\n")[1:] {
		p := re1.FindStringSubmatch(weapon)
		w := resource{
			p[1],
			toInt(p[2]),
			toInt(p[3]),
			toInt(p[4]),
		}
		weapons = append(weapons, w)
	}

	for _, armor := range strings.Split(armorInput, "\n")[1:] {
		p := re1.FindStringSubmatch(armor)
		a := resource{
			p[1],
			toInt(p[2]),
			toInt(p[3]),
			toInt(p[4]),
		}
		armors = append(armors, a)
	}

	for _, ring := range strings.Split(ringsInput, "\n")[1:] {
		p := re2.FindStringSubmatch(ring)
		r := resource{
			p[1],
			toInt(p[2]),
			toInt(p[3]),
			toInt(p[4]),
		}
		rings = append(rings, r)
	}

	return weapons, armors, rings
}

func Part1(filename string, hitPoints int) int {
	boss := newPlayer(filename)
	return findLeastExpensiveWinner(boss, hitPoints)
}

func Part2(filename string, hitPoints int) int {
	boss := newPlayer(filename)
	return findMostExpensiveLoser(boss, hitPoints)
}

func findLeastExpensiveWinner(boss player, hp int) int {
	weapons, armor, rings := getResources()
	minGold := math.MaxInt32

	var backtrack func(*player, int, int)

	isSolution := func(p player) bool {
		return p.beats(boss)
	}

	processSolution := func(c int) {
		if c < minGold {
			minGold = c
		}
	}

	constructCandidates := func(idx int) []resource {
		if idx == 0 {
			return weapons
		}
		if idx == 1 {
			return armor
		}
		if idx == 2 || idx == 3 {
			return rings
		}
		return []resource{}
	}

	backtrack = func(play *player, idx int, cost int) {
		if idx > 3 {
			return
		}
		if isSolution(*play) {
			processSolution(cost)
		} else {
			candidates := constructCandidates(idx)
			if idx != 0 {
				backtrack(play, idx+1, cost)
			}
			for _, c := range candidates {
				p := (*play)
				p.damage += c.damage
				p.armor += c.armor
				backtrack(&p, idx+1, cost+c.cost)
				p.damage -= c.damage
				p.armor -= c.armor
			}
		}
	}

	backtrack(&player{hp, 0, 0}, 0, 0)
	return minGold
}

func findMostExpensiveLoser(boss player, hp int) int {
	weapons, armor, rings := getResources()
	maxGold := math.MinInt32

	var backtrack func([]*resource, *player, int, int)

	isSolution := func(p player) bool {
		return !p.beats(boss)
	}

	processSolution := func(c int) {
		if c > maxGold {
			maxGold = c
		}
	}

	constructCandidates := func(a []*resource, idx int) []resource {
		if idx == 0 {
			return weapons
		}
		if idx == 1 {
			return armor
		}
		if idx == 2 || idx == 3 {
			res := []resource{}
			for _, r := range rings {
				if !containsResource(a, r) {
					res = append(res, r)
				}
			}
			return res
		}
		return []resource{}
	}

	backtrack = func(a []*resource, play *player, idx int, cost int) {
		if idx > 3 {
			return
		}
		if isSolution(*play) {
			processSolution(cost)
		}
		candidates := constructCandidates(a, idx)
		if idx != 0 {
			backtrack(a, play, idx+1, cost)
		}
		for _, c := range candidates {
			p := (*play)
			a[idx] = &c
			p.damage += c.damage
			p.armor += c.armor
			backtrack(a, &p, idx+1, cost+c.cost)
			a[idx] = nil
			p.damage -= c.damage
			p.armor -= c.armor
		}
	}
	a := make([]*resource, 4)
	backtrack(a, &player{hp, 0, 0}, 0, 0)
	return maxGold
}

func containsResource(s []*resource, r resource) bool {
	for _, rpointer := range s {
		if rpointer != nil {
			if (*rpointer) == r {
				return true
			}
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
