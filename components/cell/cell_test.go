package cell

import "testing"

var newCell = NewCell()

func TestNewCell(t *testing.T) {
	actual := newCell.mark
	expected := NoMark
	if actual != expected {
		t.Error("Actual is", actual, "but expected is", expected)
	}

}
