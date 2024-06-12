package model

const playerNum = 3

type Round struct {
	//Id      uint
	Number  int //player 人数
	Players []*Player
	Deck    *Deck
}

// 一副可以直接起的斗地主牌
func NewRound(players []*Player) *Round {
	round := &Round{}
	round.Number = playerNum
	round.Players = players
	poker := NewDeck()
	poker.Shuffle()
	poker.CreateLandlordDeck() //创建地主牌
	round.Deck = poker
	return round
}
