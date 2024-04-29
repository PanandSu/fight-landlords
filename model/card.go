package model

import (
	"fight-landlords/util"
)

type Card struct {
	Style      string
	Name       string // A,2,3,4,5,6,7,8,9,10,J,Q,K,Little-Joker,Big-Joker
	Number     int    //14,15,3,4,5,6,7,8,9,10,11,12,13,16,17
	IsLandlord bool
}

func NewCard(style, name string, number int, isLandlord bool) *Card {
	return &Card{
		Style:      style,
		Name:       name,
		Number:     number,
		IsLandlord: isLandlord,
	}
}

func (c *Card) String() string {
	return c.Style + "-" + c.Name
}

func (c *Card) setNameByNumber() {
	c.Name = util.Unhash(c.Number)
}
