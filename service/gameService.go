package service

import (
	"battleship/components/board"
	"battleship/components/cell"
	"battleship/components/player"
	"battleship/components/result"
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Game struct {
	board  *board.Board
	player *player.Player
	result *result.ResultAnalyser
}

func NewGame(b *board.Board, p *player.Player, r *result.ResultAnalyser) *Game {
	var newGame = &Game{
		board:  b,
		player: p,
		result: r,
	}
	return newGame
}

func (g *Game) Play() {
	var userRow int
	var status string
	var trial int
	var count int = 0
	row := g.board.GetRow()
	col := g.board.GetCol()
	// status = result.Running
	trial = (row*col)/2 + 10
	for trial != 0 {
		count += 1
		fmt.Println(g.player.GetName(), "Enter row position to attack ")
		fmt.Scanln(&userRow)
		fmt.Println(g.player.GetName(), "Enter coloumn position to attack ")
		reader := bufio.NewReader(os.Stdin)

		userC, _, _ := reader.ReadRune()
		userC = unicode.ToUpper(userC)
		var userCol int = int(userC)
		// fmt.Println(userCol)
		userCol -= 65
		if userRow > row-1 || userCol > col-1 || userCol < 0 {
			fmt.Println("Invalid position")
			continue
		}
		if g.result.CheckHitMiss(userRow, userCol) == 0 { //check if user attacked a ship or not
			fmt.Println("It was a hit at ", userRow, " ", userCol)
			g.player.IncrementHit()
			// x := g.player.GetHitCount(hit)
			g.board.Reset(userRow, userCol)
			g.board.Set(userRow, userCol, cell.Hit)
			// fmt.Println(g.board.GetBoardMark(userRow, userCol))
			fmt.Println(g.player.GetHitCount())

		} else if g.result.CheckHitMiss(userRow, userCol) == 1 {
			fmt.Println("It was a miss at ", userRow, " ", userCol)
			g.player.IncrementMiss()
			g.board.Set(userRow, userCol, cell.Miss)
			fmt.Println(g.player.GetMissCount())

		} else {
			fmt.Println("Already entered this position")
			continue
		}

		if g.player.GetHitCount() == 15 {
			status = "win"
			fmt.Println("You win")
			break
		}

		trial -= 1
		g.board.DisplayBoardForUser()

	}
	if status != "win" {
		fmt.Println("Sorry you lose!")
	}
	fmt.Println("No of hits:->", g.player.GetHitCount(), " No. of Misses:->", g.player.GetMissCount(), " No of Chances:->", count)
	fmt.Println("Game Over")

}
