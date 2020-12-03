package main

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/day1"
	"github.com/mlemesle/advent-of-code-2020/day2"
	"github.com/mlemesle/advent-of-code-2020/day3"
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
	default:
		res, err = "", errors.New("Unknow exercise")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
