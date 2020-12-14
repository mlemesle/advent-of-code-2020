package day14

import (
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"strconv"
	"strings"
)

func Part1() (int64, error) {
	t, err := lib.ReadAllLineToString("day14/input.txt")
	if err != nil {
		return 0, err
	}
	var mask []rune
	mem := map[int]int64{}
	for _, line := range t {
		if strings.HasPrefix(line, "mask = ") {
			mask = []rune(line[7:])
		} else if strings.HasPrefix(line, "mem") {
			var index int
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &index, &value)
			tmp := strconv.FormatInt(value, 2)
			valueSlice := []rune(strings.Repeat("0", 36-len(tmp)) + tmp)
			for i, r := range mask {
				if r != 'X' {
					valueSlice[i] = r
				}
			}
			memValue, err := strconv.ParseInt(string(valueSlice), 2, 64)
			if err != nil {
				return 0, err
			}
			mem[index] = memValue
		} else {
			return 0, fmt.Errorf("Bad line %s", line)
		}
	}
	var total int64
	for _, value := range mem {
		total += value
	}
	return total, nil
}

func generateIndexes(baseIndex, mask []rune) []string {
	currentIndex := ""
	init := false
	var indexes []string
	for i, r := range mask {
		switch r {
		case 'X':
			if !init {
				indexes = append(indexes, currentIndex+"0")
				indexes = append(indexes, currentIndex+"1")
				init = true
			} else {
				var tmpIndexes []string
				for _, index := range indexes {
					tmpIndexes = append(tmpIndexes, index+"0")
					tmpIndexes = append(tmpIndexes, index+"1")
				}
				indexes = tmpIndexes
			}
		case '1':
			currentIndex += string(r)
			for j, index := range indexes {
				indexes[j] = index + string(r)
			}
		default:
			currentIndex += string(baseIndex[i])
			for j, index := range indexes {
				indexes[j] = index + string(baseIndex[i])
			}
		}
	}
	return indexes
}

func Part2() (int64, error) {
	t, err := lib.ReadAllLineToString("day14/message.txt")
	if err != nil {
		return 0, err
	}
	var mask []rune
	mem := map[int64]int64{}
	for _, line := range t {
		if strings.HasPrefix(line, "mask = ") {
			mask = []rune(line[7:])
		} else if strings.HasPrefix(line, "mem") {
			var index int64
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &index, &value)
			tmp := strconv.FormatInt(index, 2)
			indexSlice := []rune(strings.Repeat("0", 36-len(tmp)) + tmp)
			indexesToWriteTo := generateIndexes(indexSlice, mask)
			for _, i := range indexesToWriteTo {
				memAddr, err := strconv.ParseInt(i, 2, 64)
				if err != nil {
					return 0, err
				}
				mem[memAddr] = value
			}
		} else {
			return 0, fmt.Errorf("Bad line %s", line)
		}
	}
	var total int64
	for _, value := range mem {
		total += value
	}
	return total, nil
}
