package cell

type Mark string

const (
	NoMark Mark = "__"
	Ship1  Mark = "S1"
	Ship2  Mark = "S2"
	Ship3  Mark = "S3"
	Ship4  Mark = "S4"
	Ship5  Mark = "S5"
	Hit    Mark = "*"
	Miss   Mark = "M"
)

type Cell struct {
	mark Mark
}

func NewCell() *Cell {
	var newCell = &Cell{
		mark: NoMark,
	}
	return newCell
}

func (c *Cell) Cell() Mark {
	return c.mark
}
func (c *Cell) SetMark(userMark Mark) {
	if c.mark == NoMark {
		c.mark = userMark
	}

}

func (c *Cell) CheckEmpty() bool {
	if c.mark == NoMark {
		return true
	}
	return false
}

func (c *Cell) ResetMark() {
	if c.mark == Ship1 || c.mark == Ship2 || c.mark == Ship3 || c.mark == Ship4 || c.mark == Ship5 {
		c.mark = NoMark
	}
}
