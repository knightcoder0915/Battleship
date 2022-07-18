package result

import (
	"battleship/components/board"
	"battleship/components/cell"
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
	if c.Cell() == cell.Ship1 || c.Cell() == cell.Ship2 || c.Cell() == cell.Ship3 || c.Cell() == cell.Ship4 || c.Cell() == cell.Ship5 {
		return 0 //means Hit
	} else if c.Cell() == cell.NoMark {
		return 1 // means Miss
	} else {
		return -1 // means repeat
	}

}
