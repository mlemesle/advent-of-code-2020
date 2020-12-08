package day8

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
)

type vmResult struct {
	acc        int
	terminated bool
}

func execPart(t []string) vmResult {
	acc := 0
	var visitedInstructions map[int]bool = make(map[int]bool)
	for i := 0; i < len(t); {
		if visitedInstructions[i] {
			return vmResult{acc, false}
		}
		visitedInstructions[i] = true
		var code string
		var param int
		fmt.Sscanf(t[i], "%s %d", &code, &param)
		switch code {
		case "acc":
			acc += param
			fallthrough
		case "nop":
			i++
		case "jmp":
			i += param
		}
	}
	return vmResult{acc, true}
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToString("day8/input.txt")
	if err != nil {
		return 0, err
	}
	return execPart(t).acc, nil
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToString("day8/input.txt")
	if err != nil {
		return 0, err
	}
	for index, instruction := range t {
		var code string
		var param int
		fmt.Sscanf(instruction, "%s %d", &code, &param)
		switch code {
		case "jmp":
			inst := fmt.Sprintf("%s %d", "nop", param)
			t[index] = inst
			if res := execPart(t); res.terminated {
				return res.acc, nil
			}
			t[index] = instruction
		case "nop":
			inst := fmt.Sprintf("%s %d", "jmp", param)
			t[index] = inst
			if res := execPart(t); res.terminated {
				return res.acc, nil
			}
			t[index] = instruction
		}
	}
	return 0, errors.New("no fault code found")
}
