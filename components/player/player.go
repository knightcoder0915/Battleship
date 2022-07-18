package player

type Player struct {
	name string
	hit  int
	miss int
	// mark cell.Mark
}

func NewPlayer(name string, hit int, miss int) *Player {
	var NewPlayer = &Player{
		name: name,
		hit:  hit,
		miss: miss,
	}
	return NewPlayer
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) IncrementHit() {
	p.hit++
}

func (p *Player) IncrementMiss() {
	p.miss++
}

func (p *Player) GetHitCount() int {
	return p.hit
}
func (p *Player) GetMissCount() int {
	return p.miss
}
