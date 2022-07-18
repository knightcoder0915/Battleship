package main

import (
	"battleship/components/board"
	"battleship/components/player"
	"battleship/components/result"
	"battleship/service"
	"fmt"
)

func main() {
	var row, col int
	// var ship *ships.Ship
	var playerName string
	fmt.Println("Welcome to Battleship!")
begin:
	fmt.Print("Enter row and col of Board: ")
	fmt.Scanln(&row, &col)
	board1 := board.CreateBoard(row, col)
	if board1 == nil {
		goto begin
	}
	service.GenerateShips(board1)
	board1.DisplayBoard()
	board1.DisplayBoardForUser()
	fmt.Println("Enter player name")
	fmt.Scanln(&playerName)
	contestant := player.NewPlayer(playerName, 0, 0)
	resultAnalyser := result.NewResultAnalyzer(board1)
	newGame := service.NewGame(board1, contestant, resultAnalyser)
	newGame.Play()

}
