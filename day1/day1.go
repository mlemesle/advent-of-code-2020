package day1

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"sort"
	"strconv"
)

func Part1(target int) (string, error) {
	t, err := lib.ReadAllLineToInt("day1/input.txt")
	if err != nil {
		return "", err
	}
	sort.Ints(t)
	for i := len(t) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			temp := t[i] + t[j]
			if temp == target {
				return strconv.Itoa(t[i] * t[j]), nil
			} else if temp > target {
				continue
			}
		}
	}
	return "", errors.New(fmt.Sprint("no couple match target : ", target))
}

func Part2(target int) (string, error) {
	t, err := lib.ReadAllLineToInt("day1/input.txt")
	if err != nil {
		return "", err
	}
	for i, x := range t[:len(t)-2] {
		for j, y := range t[i : len(t)-1] {
			for _, z := range t[j:len(t)] {
				if x+y+z == target {
					return strconv.Itoa(x * y * z), nil
				}
			}
		}
	}
	return "", errors.New(fmt.Sprint("no couple match target : ", target))
}
