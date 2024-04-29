package model

import "math/rand"

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

func generateId(length int) uint {
	bs := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	randomStr := make([]byte, length)
	for i := range length {
		randomStr[i] = bs[rand.Intn(26)]
	}
	mod := 1000_000_000_7
	h := 0
	radix := 31
	for i := range length {
		h = h*radix + int(randomStr[i])
	}
	return uint(h % mod)
}
