package ship

type ShipName string

type Ship struct {
	shipName ShipName
}

const (
	NoShip ShipName = "__"
	Ship1  ShipName = "S1"
	Ship2  ShipName = "S2"
	Ship3  ShipName = "S3"
	Ship4  ShipName = "S4"
	Ship5  ShipName = "S5"
)

var Ships = map[ShipName]int{Ship5: 5, Ship4: 4, Ship3: 3, Ship2: 2, Ship1: 1}

func NewShip() *Ship {
	var newShip = &Ship{
		shipName: NoShip,
	}
	return newShip
}

func (s *Ship) Ship() ShipName {
	return s.shipName
}

func (s *Ship) SetShip(shipName ShipName) {
	s.shipName = shipName

}

func (s *Ship) CheckEmpty() bool {
	if s.shipName == NoShip {
		return true
	}
	return false
}

func (s *Ship) ResetShip() {
	if s.shipName != NoShip {
		s.shipName = NoShip
	}
}
