package main

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/day1"
	"github.com/mlemesle/advent-of-code-2020/day10"
	"github.com/mlemesle/advent-of-code-2020/day11"
	"github.com/mlemesle/advent-of-code-2020/day12"
	"github.com/mlemesle/advent-of-code-2020/day13"
	"github.com/mlemesle/advent-of-code-2020/day14"
	"github.com/mlemesle/advent-of-code-2020/day15"
	"github.com/mlemesle/advent-of-code-2020/day2"
	"github.com/mlemesle/advent-of-code-2020/day3"
	"github.com/mlemesle/advent-of-code-2020/day4"
	"github.com/mlemesle/advent-of-code-2020/day5"
	"github.com/mlemesle/advent-of-code-2020/day6"
	"github.com/mlemesle/advent-of-code-2020/day7"
	"github.com/mlemesle/advent-of-code-2020/day8"
	"github.com/mlemesle/advent-of-code-2020/day9"
	"os"
)

func main() {
	exerciseToRun := os.Args[1]
	var res interface{}
	var err error
	switch exerciseToRun {
	case "11":
		res, err = day1.Part1(2020)
	case "12":
		res, err = day1.Part2(2020)
	case "21":
		res, err = day2.Part1()
	case "22":
		res, err = day2.Part2()
	case "31":
		res, err = day3.Part1()
	case "32":
		res, err = day3.Part2()
	case "41":
		res, err = day4.Part1()
	case "42":
		res, err = day4.Part2()
	case "51":
		res, err = day5.Part1()
	case "52":
		res, err = day5.Part2()
	case "61":
		res, err = day6.Part1()
	case "62":
		res, err = day6.Part2()
	case "71":
		res, err = day7.Part1("shiny gold")
	case "72":
		res, err = day7.Part2("shiny gold")
	case "81":
		res, err = day8.Part1()
	case "82":
		res, err = day8.Part2()
	case "91":
		res, err = day9.Part1(25)
	case "92":
		res, err = day9.Part2(25)
	case "101":
		res, err = day10.Part1()
	case "102":
		res, err = day10.Part2()
	case "111":
		res, err = day11.Part1()
	case "112":
		res, err = day11.Part2()
	case "121":
		res, err = day12.Part1()
	case "122":
		res, err = day12.Part2()
	case "131":
		res, err = day13.Part1()
	case "132":
		res, err = day13.Part2()
	case "141":
		res, err = day14.Part1()
	case "142":
		res, err = day14.Part2()
	case "151":
		res, err = day15.Part1()
	case "152":
		res, err = day15.Part2()
	default:
		res, err = "", errors.New("Unknow exercise")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
