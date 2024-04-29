package model

type Player struct {
	Nickname      string
	CurrentInning *Inning
	//Account       string
	//Password      string
	//WeChat        string
	//QQ            string
	//Coins         int
	//Innings       []*Inning
}

type Inning struct {
	HavingCards []*Card
	Sequence    *string
	IsLandlord  bool
	//Cost        int
	//Reward      int
}

func NewInning() *Inning {
	tmp := ""
	return &Inning{HavingCards: make([]*Card, 0), Sequence: &tmp}
}

/*
func (p *Player) AddInning(inning *Inning) {
	p.Innings = append(p.Innings, inning)
}

func (p *Player) CheckPassword() bool {
	password := p.Password
	if len(password) < 6 || len(password) > 16 {
		return false
	}
	return password != ""
}

func (p *Player) CheckWeChat() bool {
	weChat := p.WeChat
	return weChat != ""
}

func (p *Player) CheckQQ() bool {
	qq := p.QQ
	if len(qq) != 10 {
		return false
	}
	return qq != ""
}
*/
