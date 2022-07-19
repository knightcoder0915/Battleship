package cell

type Mark string

const (
	NoMark Mark = "__"
	Hit    Mark = "()"
	Miss   Mark = "XX"
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
	// fmt.Println(c.mark)

}

func (c *Cell) ResetMark() {
	if c.mark != NoMark {
		c.mark = NoMark
	}
}
