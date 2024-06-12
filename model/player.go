package model

type Player struct {
	Nickname      string
	CurrentInning *Inning
	//Innings       []*Inning
}

type Inning struct {
	HavingCards []*Card
	Sequence    *string
	IsLandlord  bool
}

func NewInning() *Inning {
	tmp := ""
	return &Inning{HavingCards: make([]*Card, 0), Sequence: &tmp}
}

/*
func (p *Player) AddInning(inning *Inning) {
	p.Innings = append(p.Innings, inning)
}
*/
