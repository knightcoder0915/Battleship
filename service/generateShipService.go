//change loop and if into one
package service

import (
	"battleship/components/board"
	"battleship/components/ship"
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
	var shipCell ship.ShipName

	for shipName, i := range ship.Ships {
		j := 1
		for j > 0 {
			x, y := GenerateRandom(b.GetRow(), b.GetCol())
			choosingFactor := rand.Intn(2)
			shipCell = shipName //ship Name
			if choosingFactor == 1 {
				if shipPlacement(x, y, i, row, "vertical", shipCell, b) { //Vertical
					j--
					continue
				}
				if shipPlacement(x, y, i, col, "horizontal", shipCell, b) { //Horizontal
					j--
				}
			} else {
				if shipPlacement(x, y, i, col, "horizontal", shipCell, b) {
					j--
					continue
				}
				if shipPlacement(x, y, i, row, "vertical", shipCell, b) {
					j--
				}
			}

		}
	}

}

func shipPlacement(randomNoX, randomNoY, size, rc int, placementType string, shipMark ship.ShipName, b *board.Board) bool {
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
