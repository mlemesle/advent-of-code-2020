package day11

import (
	"fmt"
	"github.com/mlemesle/advent-of-code-2020/lib"
)

const (
	floor        = '.'
	emptySeat    = 'L'
	occupiedSeat = '#'
)

type seat struct {
	current rune
	x, y    int
}

func newSeat(current rune, x, y int) *seat {
	return &seat{
		current, x, y,
	}
}

func (seat *seat) equals(o *seat) bool {
	return seat.current == o.current
}

type stateComputer func(*seat, []*seat) *seat

func stateComputerGenerator(maxPeopleCanTolerate int) stateComputer {
	return func(s *seat, n []*seat) *seat {
		r := s.current
		if s.current != floor {
			nbOccupied := 0
			for _, neighbor := range n {
				if neighbor.current == occupiedSeat {
					nbOccupied++
				}
			}
			if s.current == emptySeat && nbOccupied == 0 {
				r = occupiedSeat
			}
			if s.current == occupiedSeat && nbOccupied >= maxPeopleCanTolerate {
				r = emptySeat
			}
		}
		return newSeat(r, s.x, s.y)
	}
}

type seatsMatrix [][]*seat

func (sm *seatsMatrix) equals(o *seatsMatrix) bool {
	for x := 0; x < len(*sm); x++ {
		for y := 0; y < len((*sm)[x]); y++ {
			if !(*sm)[x][y].equals((*o)[x][y]) {
				return false
			}
		}
	}
	return true
}

type getNeighborFunction func(*seatsMatrix, *seat) []*seat

func getNeighborsStep1(sm *seatsMatrix, s *seat) []*seat {
	var neighbors []*seat
	canGetTop := s.x > 0
	canGetDown := s.x < len(*sm)-1
	canGetLeft := s.y > 0
	canGetRight := s.y < len((*sm)[0])-1
	if canGetTop {
		neighbors = append(neighbors, (*sm)[s.x-1][s.y])
		if canGetLeft {
			neighbors = append(neighbors, (*sm)[s.x-1][s.y-1])
		}
		if canGetRight {
			neighbors = append(neighbors, (*sm)[s.x-1][s.y+1])
		}
	}
	if canGetDown {
		neighbors = append(neighbors, (*sm)[s.x+1][s.y])
		if canGetLeft {
			neighbors = append(neighbors, (*sm)[s.x+1][s.y-1])
		}
		if canGetRight {
			neighbors = append(neighbors, (*sm)[s.x+1][s.y+1])
		}
	}
	if canGetLeft {
		neighbors = append(neighbors, (*sm)[s.x][s.y-1])
	}
	if canGetRight {
		neighbors = append(neighbors, (*sm)[s.x][s.y+1])
	}
	return neighbors
}

func getNeighborsStep2(sm *seatsMatrix, s *seat) []*seat {
	var neighbors []*seat
	canGetTop := s.x > 0
	canGetDown := s.x < len(*sm)-1
	canGetLeft := s.y > 0
	canGetRight := s.y < len((*sm)[0])-1
	if canGetTop {
		for i := s.x - 1; i >= 0; i-- {
			if (*sm)[i][s.y].current != floor {
				neighbors = append(neighbors, (*sm)[i][s.y])
				break
			}
		}
		if canGetLeft {
			for i, j := s.x-1, s.y-1; i >= 0 && j >= 0; {
				if (*sm)[i][j].current != floor {
					neighbors = append(neighbors, (*sm)[i][j])
					break
				}
				i--
				j--
			}
		}
		if canGetRight {
			for i, j := s.x-1, s.y+1; i >= 0 && j < len((*sm)[i]); {
				if (*sm)[i][j].current != floor {
					neighbors = append(neighbors, (*sm)[i][j])
					break
				}
				i--
				j++
			}
		}
	}
	if canGetDown {
		for i := s.x + 1; i < len(*sm); i++ {
			if (*sm)[i][s.y].current != floor {
				neighbors = append(neighbors, (*sm)[i][s.y])
				break
			}
		}
		if canGetLeft {
			for i, j := s.x+1, s.y-1; i < len(*sm) && j >= 0; {
				if (*sm)[i][j].current != floor {
					neighbors = append(neighbors, (*sm)[i][j])
					break
				}
				i++
				j--
			}
		}
		if canGetRight {
			for i, j := s.x+1, s.y+1; i < len(*sm) && j < len((*sm)[i]); {
				if (*sm)[i][j].current != floor {
					neighbors = append(neighbors, (*sm)[i][j])
					break
				}
				i++
				j++
			}
		}
	}
	if canGetLeft {
		for j := s.y - 1; j >= 0; j-- {
			if (*sm)[s.x][j].current != floor {
				neighbors = append(neighbors, (*sm)[s.x][j])
				break
			}
		}
	}
	if canGetRight {
		for j := s.y + 1; j < len((*sm)[s.x]); j++ {
			if (*sm)[s.x][j].current != floor {
				neighbors = append(neighbors, (*sm)[s.x][j])
				break
			}
		}
	}
	return neighbors
}

func (sm *seatsMatrix) doStep(computer stateComputer, neighborFunc getNeighborFunction) seatsMatrix {
	var newSm seatsMatrix
	for _, row := range *sm {
		var seatRow []*seat
		for _, seat := range row {
			neighbors := neighborFunc(sm, seat)
			s := computer(seat, neighbors)
			seatRow = append(seatRow, s)
		}
		newSm = append(newSm, seatRow)
	}
	return newSm
}

func (sm *seatsMatrix) getNbSeatsOccupied() int {
	nbOccupied := 0
	for _, row := range *sm {
		for _, s := range row {
			if s.current == occupiedSeat {
				nbOccupied++
			}
		}
	}
	return nbOccupied
}

func (sm *seatsMatrix) print() {
	for _, row := range *sm {
		for _, s := range row {
			fmt.Print(string(s.current))
		}
		fmt.Println()
	}
	fmt.Println()
}

func toSeatMatrix(t [][]rune) seatsMatrix {
	var seats [][]*seat
	for x, row := range t {
		var seatRow []*seat
		for y, s := range row {
			seatRow = append(seatRow, &seat{
				x:       x,
				y:       y,
				current: s,
			})
		}
		seats = append(seats, seatRow)
	}
	return seats
}

func Part1() (int, error) {
	t, err := lib.ReadAllLineToRuneSlice("day11/message.txt")
	if err != nil {
		return 0, err
	}
	seats := toSeatMatrix(t)
	stateCmp := stateComputerGenerator(4)
	o := seats.doStep(stateCmp, getNeighborsStep1)
	for !seats.equals(&o) {
		seats = o
		o = seats.doStep(stateCmp, getNeighborsStep1)
	}
	return seats.getNbSeatsOccupied(), nil
}

func Part2() (int, error) {
	t, err := lib.ReadAllLineToRuneSlice("day11/message.txt")
	if err != nil {
		return 0, err
	}
	seats := toSeatMatrix(t)
	stateCmp := stateComputerGenerator(5)
	o := seats.doStep(stateCmp, getNeighborsStep2)
	for !seats.equals(&o) {
		seats = o
		o = seats.doStep(stateCmp, getNeighborsStep2)
	}
	return seats.getNbSeatsOccupied(), nil
}
