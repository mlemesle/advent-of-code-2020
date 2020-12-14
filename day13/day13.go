package day13

import (
	"github.com/mlemesle/advent-of-code-2020/lib"
	"strconv"
	"strings"
)

func toBusSlice(input string) []int {
	var buses []int
	for _, b := range strings.Split(input, ",") {
		if b != "x" {
			bus, _ := strconv.Atoi(b)
			buses = append(buses, bus)
		}
	}
	return buses
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToString("day13/input.txt")
	if err != nil {
		return 0, err
	}
	timestamp, _ := strconv.Atoi(t[0])
	buses := toBusSlice(t[1])
	var timeRemaining, busID int
	for i, bus := range buses {
		if i == 0 {
			timeRemaining, busID = bus-timestamp%bus, bus
		} else if bus != 0 {
			t := bus - timestamp%bus
			if t < timeRemaining {
				timeRemaining, busID = t, bus
			}
		}
	}
	return timeRemaining * busID, nil
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToString("day13/input.txt")
	if err != nil {
		return 0, err
	}
	buses := map[int]int{}
	for i, bus := range strings.Split(t[1], ",") {
		if bus != "x" {
			buses[i], _ = strconv.Atoi(bus)
		}
	}
	timestamp := buses[0]
	diff := buses[0]
	for i, bus := range buses {
		for {
			if (timestamp+i)%bus == 0 {
				break
			}
			timestamp += diff
		}
		diff = leastCommonMultiple(diff, bus)
	}
	return timestamp, nil
}
