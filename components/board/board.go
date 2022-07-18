package board

import (
	"battleship/components/cell"
	"fmt"
)

type Board struct {
	rows      int
	cols      int
	boardMark [][]*cell.Cell // 2D slice of pointer of cell containing all marks on the board
	userMark  [][]*cell.Cell // 2D slice of pointer of cell containing all usermarks on the board
}

func CreateBoard(rows, cols int) *Board { //Creating a board
	if rows < cols && rows >= 5 && cols >= 5 {
		fmt.Println("Board cannot be created")
		return nil
	}
	var cells = make([][]*cell.Cell, rows*cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cells[i] = append(cells[i], cell.NewCell())

		}
	}
	var board = &Board{
		rows:      rows,
		cols:      cols,
		boardMark: cells,
		userMark:  cells,
	}
	return board

}

func (b *Board) DisplayBoard() { //Printing the board
	// count := 0
	fmt.Print("\t")
	for ch := 65; ch < 65+b.cols; ch++ {
		fmt.Printf("       %c", ch)
	}
	fmt.Println()

	for i := 0; i < b.rows; i++ {
		fmt.Print(i, " \t")
		for j := 0; j < b.cols; j++ {
			c := b.boardMark[i][j]
			// fmt.Print(i, j)
			fmt.Print("   |   ", c.Cell())
			// count++

		}
		fmt.Print("   |  ")
		fmt.Print("\t", i, "\n")

	}
	fmt.Println()
	fmt.Print("\t")
	for ch := 65; ch < 65+b.cols; ch++ {
		fmt.Printf("       %c", ch)
	}
	fmt.Println()
}

func (b *Board) DisplayBoardForUser() { //Printing the board
	// count := 0
	fmt.Print("\t")
	for ch := 65; ch < 65+b.cols; ch++ {
		fmt.Printf("       %c", ch)
	}
	fmt.Println()

	for i := 0; i < b.rows; i++ {
		fmt.Print(i, " \t")
		for j := 0; j < b.cols; j++ {
			c := b.userMark[i][j]
			if c.Cell() == cell.Ship1 || c.Cell() == cell.Ship2 || c.Cell() == cell.Ship3 || c.Cell() == cell.Ship4 || c.Cell() == cell.Ship5 && c.Cell() != cell.Hit {
				fmt.Print("   |   ", cell.NoMark)
				// fmt.Print(c.GetMark())
				// count++}
			} else {
				fmt.Print("   |   ", c.Cell())
			}

		}
		fmt.Print("   |  ")
		fmt.Print("\t", i, "\n")

	}
	fmt.Println()
	fmt.Print("\t")
	for ch := 65; ch < 65+b.cols; ch++ {
		fmt.Printf("       %c", ch)
	}
	fmt.Println()
}

func (b *Board) Get(row, col int) bool { //check if cell is empty -->If empty returns true
	return b.boardMark[row][col].CheckEmpty()

}

func (b *Board) Set(row, col int, mark cell.Mark) { //to set the ship
	if b.Get(row, col) {
		b.boardMark[row][col].SetMark(mark)

	}
}

func (b *Board) Reset(row, col int) { //to set the ship

	b.boardMark[row][col].ResetMark()

}

func (b *Board) GetRow() int {
	return b.rows
}

func (b *Board) GetCol() int {
	return b.cols
}

func (b *Board) GetBoardMark(row, col int) *cell.Cell {
	return b.boardMark[row][col]
}

func (b *Board) SetUserMark(row, col int, mark cell.Mark) {
	b.userMark[row][col].SetMark(mark)
}
