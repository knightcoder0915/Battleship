package result

import (
	"battleship/components/board"
	"battleship/components/cell"
	"battleship/components/ship"
	"fmt"
)

type Status string

const (
	Win     Status = "Win"
	Lose    Status = "Lose"
	Running Status = "Running"
)

type ResultAnalyser struct {
	board *board.Board
}

func NewResultAnalyzer(board *board.Board) *ResultAnalyser {
	var newResultAnalyser = &ResultAnalyser{
		board: board,
	}
	return newResultAnalyser
}

func (r *ResultAnalyser) CheckHitMiss(userRow, userCol int) int {
	s := r.board.GetBoardShipMark(userRow, userCol)
	c := r.board.GetBoardCellMark(userRow, userCol)
	shipName := s.Ship()
	if shipName == ship.NoShip && c.Cell() == cell.NoMark {
		return 1
	}

	for shipMark, _ := range ship.Ships {
		if shipName == shipMark && c.Cell() == cell.NoMark {
			ship.Ships[shipMark]--
			if r.CheckDestruction(shipMark, userRow, userCol) == 0 {
				return 0
			}
			return 0
		}
	}

	return -1

}

func (r *ResultAnalyser) CheckDestruction(shipName ship.ShipName, userRow, userCol int) int {
	// fmt.Println("Map of ships ", ship.Ships)
	// fmt.Println("THis is shipName", shipName)
	if ship.Ships[shipName] == 0 && shipName != ship.NoShip {
		fmt.Println(shipName, "has been destroyed!!!!")
		return 0
	}
	return 1

}
