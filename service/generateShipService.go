//change loop and if into one
package service

import (
	"battleship/components/board"
	"battleship/components/cell"
	"math/rand"
	"time"
)

func GenerateRandom(row, col int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	randomNoX := rand.Intn(row)
	randomNoY := rand.Intn(col)
	return randomNoX, randomNoY

}

func GenerateShips(b *board.Board) {
	row := b.GetRow()
	col := b.GetCol()
	rand.Seed(time.Now().UnixNano())
	var shipCell cell.Mark
	i := 5
	for i > 0 {
		x, y := GenerateRandom(b.GetRow(), b.GetCol())
		choosingFactor := rand.Intn(2)
		// fmt.Println(choosingFactor)
		if i == 1 {
			shipCell = cell.Ship1
		} else if i == 2 {
			shipCell = cell.Ship2

		} else if i == 3 {
			shipCell = cell.Ship3
		} else if i == 4 {
			shipCell = cell.Ship4
		} else {
			shipCell = cell.Ship5
		}

		if choosingFactor == 1 {
			if shipPlacement(x, y, i, row, "vertical", shipCell, b) { //Vertical
				i--
				continue
			}
			if shipPlacement(x, y, i, col, "horizontal", shipCell, b) { //Horizontal
				i--
			}
		} else {
			if shipPlacement(x, y, i, col, "horizontal", shipCell, b) {
				i--
				continue
			}
			if shipPlacement(x, y, i, row, "vertical", shipCell, b) {
				i--
			}
		}

	}
}

func shipPlacement(randomNoX, randomNoY, size, rc int, placementType string, shipMark cell.Mark, b *board.Board) bool {
	var start, stop int
	var count int = 0
	checkEmpty := b.Get(randomNoX, randomNoY) //checking empty
	if checkEmpty == false {
		return false

	}
	if placementType == "vertical" {
		start = randomNoX
	} else {
		start = randomNoY
	}
	var i int
	for i = start; i < rc; i++ {
		if placementType == "vertical" {
			checkEmpty = b.Get(i, randomNoY) //checking empty
		} else {
			checkEmpty = b.Get(randomNoX, i)
		}
		if checkEmpty == false { //not empty
			goto checkUp
		}
		count += 1
	}
checkUp:
	start = i
	stop = start - 1
	if placementType == "vertical" {
		start = randomNoX
	} else {
		start = randomNoY
	}
	for i := 0; i < rc; i++ {
		start -= 1
		if start >= 0 {
			if placementType == "vertical" {
				checkEmpty = b.Get(start, randomNoY) //checking empty
			} else {
				checkEmpty = b.Get(randomNoX, start)
			}
			if checkEmpty == false {
				if count < size {
					return false
				} else {
					break
				}
			}
			count += 1
		} else {
			break
		}

	}

	start = start + 1
	if stop-start+1 >= size {
		startCount := 0
		rand.Seed(time.Now().UnixNano())
		min := 1
		choosingFactor := rand.Intn(2-min+1) + min
		if choosingFactor == 1 {

			for i := 0; i < size; i++ {
				if start <= stop {
					if placementType == "vertical" {
						b.Set(start, randomNoY, shipMark) //ships will be placed vertically up to down
					} else {
						b.Set(randomNoX, start, shipMark) //ships will be placed horizontally from left
					}
					start += 1
					startCount += 1
					if startCount == size {
						break
					}
				}
			}
		} else {
			for i := 0; i < size; i++ { // ad if in loop ka condition
				if stop >= start {
					if placementType == "vertical" {
						b.Set(stop, randomNoY, shipMark) //ships will be placed vertically from down to up
					} else {
						b.Set(randomNoX, stop, shipMark) //ships will be placed horizontally from right
					}

					stop -= 1
					startCount += 1
					if startCount == size {
						break
					}

				}
			}
		}
		// b.DisplayBoard()

	} else {
		return false
	}
	// b.DisplayBoard()
	return true

}
