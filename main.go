package main

import (
	"errors"
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/day1"
	"os"
)

func main() {
	exerciseToRun := os.Args[1]
	var res string
	var err error
	switch exerciseToRun {
	case "11":
		res, err = day1.Part1(2020)
	case "12":
		res, err = day1.Part2(2020)
	default:
		res, err = "", errors.New("Unknow exercise")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
