package day2

import (
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"strings"
)

type passwordPolicy struct {
	x int
	y int
	c string
}

func generatePasswordMap(list []string) map[passwordPolicy][]string {
	var result map[passwordPolicy][]string = make(map[passwordPolicy][]string)
	for _, s := range list {
		policyAndPassword := strings.Split(s, ":")
		var min, max int
		var c string
		fmt.Sscanf(policyAndPassword[0], "%d-%d %s", &min, &max, &c)
		passwordPolicy := passwordPolicy{min, max, c}
		result[passwordPolicy] = append(result[passwordPolicy], strings.TrimSpace(policyAndPassword[1]))
	}
	return result
}

func execPart(f func(passwordPolicy, string) int) (int, error) {
	t, err := lib.ReadAllLineToString("day2/input.txt")
	if err != nil {
		return 0, err
	}
	policiesAndPasswords := generatePasswordMap(t)
	valids := 0
	for policy, passwords := range policiesAndPasswords {
		for _, password := range passwords {
			valids += f(policy, password)
		}
	}
	return valids, nil
}

func Part1() (int, error) {
	f := func(policy passwordPolicy, password string) int {
		count := strings.Count(password, policy.c)
		if count >= policy.x && count <= policy.y {
			return 1
		}
		return 0
	}
	return execPart(f)
}

func Part2() (int, error) {
	f := func(policy passwordPolicy, password string) int {
		runes := []rune(password)
		var letter rune
		fmt.Sscanf(policy.c, "%c", &letter)
		if runes[policy.x-1] == letter && runes[policy.y-1] != letter || runes[policy.x-1] != letter && runes[policy.y-1] == letter {
			return 1
		}
		return 0
	}
	return execPart(f)
}
