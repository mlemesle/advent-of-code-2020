package day6

import (
	"github.com/mlemesle/advent-of-code-2020/lib"
)

func emptyLine(s []rune) bool {
	return len(s) == 0
}

func getNumberQuestionsAnswered(m map[rune]int, nbPerson int) int {
	total := 0
	for _, v := range m {
		if v == nbPerson {
			total++
		}
	}
	return total
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToRuneSlice("day6/input.txt")
	if err != nil {
		return 0, err
	}
	var set map[rune]int = make(map[rune]int)
	total := 0
	for _, answer := range t {
		if emptyLine(answer) {
			total += getNumberQuestionsAnswered(set, 1)
			set = make(map[rune]int)
		}
		for _, question := range answer {
			set[question] = 1
		}
	}
	total += getNumberQuestionsAnswered(set, 1)
	return total, nil
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToRuneSlice("day6/input.txt")
	if err != nil {
		return 0, err
	}
	var set map[rune]int = make(map[rune]int)
	total := 0
	nbPerson := 0
	for _, answer := range t {
		if emptyLine(answer) {
			total += getNumberQuestionsAnswered(set, nbPerson)
			nbPerson = 0
			set = make(map[rune]int)
		} else {
			for _, question := range answer {
				set[question]++
			}
			nbPerson++
		}
	}
	total += getNumberQuestionsAnswered(set, nbPerson)
	return total, nil
}
