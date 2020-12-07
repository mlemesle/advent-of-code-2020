package day7

import (
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"strings"
)

type bag struct {
	color     string
	childrens map[string]int
}

func newBag(color string, childrens map[string]int) bag {
	return bag{color, childrens}
}

func bagCanContain(bag bag, allBags map[string]bag, bagToFind string) bool {
	for color := range bag.childrens {
		if color == bagToFind {
			return true
		}
	}
	for color := range bag.childrens {
		if bagCanContain(allBags[color], allBags, bagToFind) {
			return true
		}
	}
	return false
}

func parseAllBags(t []string) map[string]bag {
	var bags map[string]bag = make(map[string]bag)
	for _, line := range t {
		if strings.Contains(line, "bags contain no other bags.") {
			s := strings.Split(line, " bags contain")
			bags[s[0]] = newBag(s[0], nil)
		} else {
			line = line[:len(line)]
			s := strings.Replace(line, " contain", ",", 1)
			sTab := strings.Split(s, ", ")
			var tone, color string
			fmt.Sscanf(sTab[0], "%s %s bags", &tone, &color)
			key := tone + " " + color
			var c map[string]int = make(map[string]int)
			for _, child := range sTab[1:] {
				var nbBags int
				fmt.Sscanf(child, "%d %s %s bags", &nbBags, &tone, &color)
				c[tone+" "+color] = nbBags
			}
			bags[key] = newBag(key, c)
		}
	}
	return bags
}

func Part1(bagToFind string) (int, error) {
	t, err := lib.ReadAllLineToString("day7/input.txt")
	if err != nil {
		return 0, err
	}

	bags := parseAllBags(t)
	count := 0
	for _, b := range bags {
		if bagCanContain(b, bags, bagToFind) {
			count++
		}
	}
	return count, nil
}

func countBagsNeeded(allBags map[string]bag, bagName string) int {
	count := 0
	for color, n := range allBags[bagName].childrens {
		count += n * countBagsNeeded(allBags, color)
	}
	return count + 1
}

func Part2(bagToCount string) (int, error) {
	t, err := lib.ReadAllLineToString("day7/input.txt")
	if err != nil {
		return 0, err
	}

	bags := parseAllBags(t)
	return countBagsNeeded(bags, bagToCount) - 1, nil
}
