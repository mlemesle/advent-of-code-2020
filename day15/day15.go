package day15

import ()

type gameNumber struct {
	nbTimesSeen, previousTurn, lastTurn int
}

func execPart(nth int) (int, error) {
	t := [6]int{1, 12, 0, 20, 8, 16}
	m := map[int]*gameNumber{}
	turn := 1
	var lastSpoken int
	for turn <= nth {
		if turn <= len(t) {
			lastSpoken = t[turn-1]
			m[lastSpoken] = &gameNumber{
				nbTimesSeen:  0,
				previousTurn: turn,
				lastTurn:     0,
			}
		} else {
			spokenTurn, exist := m[lastSpoken]
			if exist && spokenTurn.nbTimesSeen == 0 {
				lastSpoken = 0
			} else {
				lastSpoken = spokenTurn.previousTurn - spokenTurn.lastTurn
			}

			spokenTurn, exist = m[lastSpoken]
			if exist {
				spokenTurn.nbTimesSeen++
				spokenTurn.previousTurn, spokenTurn.lastTurn = turn, spokenTurn.previousTurn
			} else {
				m[lastSpoken] = &gameNumber{
					nbTimesSeen:  0,
					previousTurn: turn,
					lastTurn:     0,
				}
			}
		}
		turn++
	}

	return lastSpoken, nil
}

func Part1() (int, error) {
	return execPart(2020)
}

func Part2() (int, error) {
	return execPart(30000000)
}
