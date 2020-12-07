package main

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/day1"
	"github.com/mlemesle/advent-of-code-2020/day2"
	"github.com/mlemesle/advent-of-code-2020/day3"
	"github.com/mlemesle/advent-of-code-2020/day4"
	"github.com/mlemesle/advent-of-code-2020/day5"
	"github.com/mlemesle/advent-of-code-2020/day6"
	"github.com/mlemesle/advent-of-code-2020/day7"
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
	default:
		res, err = "", errors.New("Unknow exercise")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
