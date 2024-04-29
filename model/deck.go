package model

import "math/rand"

type Deck struct {
	Cards []*Card
}

type Poker Deck

// 一副崭新的牌
func NewDeck() *Deck {
	var poker Deck
	var cards []*Card
	var styles = []string{"Club", "Diamond", "Heart", "Spade"}
	for i := 3; i <= 15; i++ {
		for _, style := range styles {
			card := &Card{Number: i, Style: style}
			card.setNameByNumber()
			cards = append(cards, card)
		}
	}
	littleJoker := &Card{Name: "Little-Joker", Number: 16}
	bigJoker := &Card{Name: "Big-Joker", Number: 17}
	cards = append(cards, littleJoker, bigJoker)
	poker = Deck{
		Cards: cards,
	}
	deck := &poker
	return deck
}

func (d *Deck) Shuffle() {
	cards := d.Cards
	n := d.Len()
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func (d *Deck) CreateLandlordDeck() {
	i := rand.Intn(d.Len() - 3) //在前51张牌之中选定地主牌
	d.Cards[i].IsLandlord = true
}

func (d *Deck) Pop() *Card {
	if d.IsEmpty() {
		panic("Empty deck")
	}
	top := d.Cards[d.Len()-1]
	d.Cards = d.Cards[:d.Len()-1]
	return top
}

func (d *Deck) String() string {
	res := ""
	for _, card := range d.Cards {
		res += card.String()
	}
	return res
}

func (d *Deck) Len() int {
	return len(d.Cards)
}

func (d *Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}
