package day10

import (
	"github.com/mlemesle/advent-of-code-2020/lib"
	"sort"
)

func Part1() (int, error) {
	t, err := lib.ReadAllLineToInt("day10/input.txt")
	if err != nil {
		return 0, err
	}
	sort.Ints(t)
	var oneDiff, threeDiff, currentAdapter int
	for _, adapter := range t {
		if currentAdapter+1 == adapter {
			oneDiff++
		} else if currentAdapter+3 == adapter {
			threeDiff++
		}
		currentAdapter = adapter
	}
	threeDiff++
	return oneDiff * threeDiff, nil
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToInt("day10/input.txt")
	if err != nil {
		return 0, err
	}
	t = append(t, 0)
	sort.Ints(t)
	t = append(t, t[len(t)-1]+3)
	depth := make([]int, len(t))
	depth[len(t)-1] = 1
	for i := 0; i < len(t); i++ {
		index := len(t) - i - 1
		for j := index - 1; j >= 0 && j >= index-3; j-- {
			if t[index]-t[j] <= 3 {
				depth[j] += depth[index]
			}
		}
	}
	return depth[0], nil
}
