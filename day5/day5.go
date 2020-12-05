package day5

import (
	"errors"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"sort"
)

func computeSeatID(seatPlacement string) int {
	rowUp := 127
	rowDown := 0
	columnUp := 7
	columnDown := 0
	for _, seatIntruction := range seatPlacement {
		switch seatIntruction {
		case 'F':
			rowUp = rowUp - ((rowUp - rowDown) / 2)
		case 'B':
			rowDown = rowDown + ((rowUp - rowDown + 1) / 2)
		case 'L':
			columnUp = columnUp - ((columnUp - columnDown) / 2)
		case 'R':
			columnDown = columnDown + ((columnUp - columnDown + 1) / 2)
		}
	}
	return rowDown*8 + columnDown
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToString("day5/input.txt")
	if err != nil {
		return 0, err
	}
	highestSeatID := -1
	for _, seatPlacement := range t {
		seatID := computeSeatID(seatPlacement)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID, nil
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToString("day5/input.txt")
	if err != nil {
		return 0, err
	}
	var seatIDs []int
	for _, seatPlacement := range t {
		seatIDs = append(seatIDs, computeSeatID(seatPlacement))
	}
	sort.Ints(seatIDs)
	s1 := seatIDs[0]
	for _, s2 := range seatIDs[1:] {
		if s2-s1 > 1 {
			return s1 + 1, nil
		}
		s1 = s2
	}
	return 0, errors.New("couldn't find seatID")
}
