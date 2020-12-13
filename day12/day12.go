package day12

import (
	"github.com/mlemesle/advent-of-code-2020/lib"
	"math"
	"strconv"
)

var cardinals [4]string = [4]string{"N", "E", "S", "W"}

type waypoint struct {
	x, y int
}

func rotate(currentDirection, rotateDirection string, degree int) string {
	var index int
	for i, d := range cardinals {
		if currentDirection == d {
			index = i
			break
		}
	}
	switch rotateDirection {
	case "R":
		index = (index + degree/90) % len(cardinals)
	case "L":
		index = (index - degree/90 + len(cardinals)) % len(cardinals)
	}
	return cardinals[index]
}

func move(x, y, step int, direction string) (int, int) {
	switch direction {
	case "N":
		y += step
	case "S":
		y -= step
	case "E":
		x += step
	case "W":
		x -= step
	}
	return x, y
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToString("day12/input.txt")
	if err != nil {
		return 0, err
	}
	w := waypoint{0, 0}
	direction := "E"
	for _, line := range t {
		indication := line[:1]
		n, _ := strconv.Atoi(line[1:])
		switch indication {
		case "N", "S", "E", "W":
			w.x, w.y = move(w.x, w.y, n, indication)
		case "R", "L":
			direction = rotate(direction, indication, n)
		case "F":
			w.x, w.y = move(w.x, w.y, n, direction)
		}
	}
	return int(math.Abs(float64(w.x)) + math.Abs(float64(w.y))), nil
}

func (w *waypoint) rotateWaypoint(indication string, degree int) {
	switch degree {
	case 0, 360:
		break
	case 270:
		switch indication {
		case "L":
			w.x, w.y = w.y, -w.x
		case "R":
			w.x, w.y = -w.y, w.x
		}
	case 180:
		w.x, w.y = -w.x, -w.y
	case 90:
		switch indication {
		case "L":
			w.x, w.y = -w.y, w.x
		case "R":
			w.x, w.y = w.y, -w.x
		}
	}
}

func moveTowardWaypoint(x, y, n int, w waypoint) (int, int) {
	return x + w.x*n, y + w.y*n
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToString("day12/input.txt")
	if err != nil {
		return 0, err
	}
	w := waypoint{10, 1}
	x, y := 0, 0
	for _, line := range t {
		indication := line[:1]
		n, _ := strconv.Atoi(line[1:])
		switch indication {
		case "N", "S", "E", "W":
			w.x, w.y = move(w.x, w.y, n, indication)
		case "R", "L":
			w.rotateWaypoint(indication, n)
		case "F":
			x, y = moveTowardWaypoint(x, y, n, w)
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y))), nil
}
