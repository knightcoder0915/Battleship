package result

import (
	"battleship/components/board"
	"battleship/components/cell"
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
	c := r.board.GetBoardMark(userRow, userCol)
	if c.Cell() == cell.Ship1 {
		r.board.Reset(userRow, userCol)
		r.board.Set(userRow, userCol, cell.Hit)
		fmt.Println("Ship 1 has been destroyed!")
		return 0 //means Hit
	} else if c.Cell() == cell.Ship2 {
		r.board.Reset(userRow, userCol)
		r.board.Set(userRow, userCol, cell.Hit)
		destroyed := r.CheckShipExist(cell.Ship2)
		if destroyed == true {
			fmt.Println("Ship 2 has been destroyed")
		}
		return 0

	} else if c.Cell() == cell.Ship3 {
		r.board.Reset(userRow, userCol)
		r.board.Set(userRow, userCol, cell.Hit)
		destroyed := r.CheckShipExist(cell.Ship3)
		if destroyed == true {
			fmt.Println("Ship 3 has been destroyed")
		}
		return 0

	} else if c.Cell() == cell.Ship4 {
		r.board.Reset(userRow, userCol)
		r.board.Set(userRow, userCol, cell.Hit)
		destroyed := r.CheckShipExist(cell.Ship4)
		if destroyed == true {
			fmt.Println("Ship 4 has been destroyed")
		}
		return 0

	} else if c.Cell() == cell.Ship5 {
		r.board.Reset(userRow, userCol)
		r.board.Set(userRow, userCol, cell.Hit)
		destroyed := r.CheckShipExist(cell.Ship5)
		if destroyed == true {
			fmt.Println("Ship 5 has been destroyed")
		}
		return 0

	} else if c.Cell() == cell.NoMark {
		return 1 // means Miss
	} else {
		return -1 // means repeat
	}

}

func (r *ResultAnalyser) CheckShipExist(mark cell.Mark) bool {
	var count int = 0
	for i := 0; i < r.board.GetCol(); i++ {
		for j := 0; j < r.board.GetRow(); j++ {
			c := r.board.GetBoardMark(i, j)
			if mark == c.Cell() {
				count++
			}

		}
	}
	if count == 0 {
		return true
	}
	return false
	// fmt.Println(*&allCells)
}
