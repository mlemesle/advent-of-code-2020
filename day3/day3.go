package day3

import (
	"github.com/mlemesle/advent-of-code-2020/lib"
)

type slope struct {
	right, down int
}

func execPart(t [][]rune, s slope) int {
	count := 0
	x := s.right
	for i := s.down; i < len(t); i += s.down {
		if t[i][x%len(t[i])] == '#' {
			count++
		}
		x += s.right
	}
	return count
}

func Part1() (interface{}, error) {
	t, err := lib.ReadAllLineToRuneSlice("day3/input.txt")
	if err != nil {
		return 0, err
	}
	s := slope{
		right: 3,
		down:  1,
	}
	return execPart(t, s), nil
}

func Part2() (interface{}, error) {
	t, err := lib.ReadAllLineToRuneSlice("day3/input.txt")
	if err != nil {
		return 0, err
	}
	s1 := slope{
		right: 1,
		down:  1,
	}
	s2 := slope{
		right: 3,
		down:  1,
	}
	s3 := slope{
		right: 5,
		down:  1,
	}
	s4 := slope{
		right: 7,
		down:  1,
	}
	s5 := slope{
		right: 1,
		down:  2,
	}
	return execPart(t, s1) * execPart(t, s2) * execPart(t, s3) * execPart(t, s4) * execPart(t, s5), nil
}
