package day4

import (
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

type passports []passport

func newPassport() passport {
	return passport{}
}

func execPart(passportValidator func(p *passport) bool) (int, error) {
	t, err := lib.ReadAllLineToString("day4/input.txt")
	if err != nil {
		return 0, err
	}
	var passports passports
	passport := newPassport()
	for _, line := range t {
		if line == "" {
			passports = append(passports, passport)
			passport = newPassport()
			continue
		}
		for _, pair := range strings.Split(line, " ") {
			var key, value string
			pair = strings.ReplaceAll(pair, ":", " ")
			fmt.Sscanf(pair, "%s %s", &key, &value)
			switch key {
			case "byr":
				passport.byr = value
			case "iyr":
				passport.iyr = value
			case "eyr":
				passport.eyr = value
			case "hgt":
				passport.hgt = value
			case "hcl":
				passport.hcl = value
			case "ecl":
				passport.ecl = value
			case "pid":
				passport.pid = value
			case "cid":
				passport.cid = value
			}
		}
	}
	passports = append(passports, passport)

	count := 0
	for _, p := range passports {
		if passportValidator(&p) {
			count++
			continue
		}
		fmt.Printf("%+v\n", p)
	}

	return count, nil
}

func Part1() (int, error) {
	f := func(p *passport) bool {
		return p.byr != "" &&
			p.iyr != "" &&
			p.eyr != "" &&
			p.hgt != "" &&
			p.hcl != "" &&
			p.ecl != "" &&
			p.pid != ""
	}
	return execPart(f)
}

func Part2() (int, error) {
	f := func(p *passport) bool {
		// check byr
		if byr, err := strconv.Atoi(p.byr); err != nil || byr < 1920 || byr > 2002 {
			return false
		}
		// check iyr
		if iyr, err := strconv.Atoi(p.iyr); err != nil || iyr < 2010 || iyr > 2020 {
			return false
		}
		// check eyr
		if eyr, err := strconv.Atoi(p.eyr); err != nil || eyr < 2020 || eyr > 2030 {
			return false
		}
		// check hgt
		var hgt int = 0
		var unit string
		fmt.Sscanf(p.hgt, "%d%s", &hgt, &unit)
		if (unit != "cm" && unit != "in") || unit == "cm" && (hgt < 150 || hgt > 193) || unit == "in" && (hgt < 59 || hgt > 76) {
			return false
		}
		if ok, err := regexp.MatchString("^#[0-9a-f]{6}$", p.hcl); err != nil || !ok {
			return false
		}
		if !(p.ecl == "amb" || p.ecl == "blu" || p.ecl == "brn" || p.ecl == "gry" || p.ecl == "grn" || p.ecl == "hzl" || p.ecl == "oth") {
			return false
		}
		if ok, err := regexp.MatchString("^[0-9]{9}$", p.pid); err != nil || !ok {
			return false
		}
		return true
	}
	return execPart(f)
}
