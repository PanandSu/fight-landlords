package model

import "sort"

type Cards []*Card

func (cs Cards) Len() int {
	return len(cs)
}

func (cs Cards) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Cards) Less(i, j int) bool {
	card1 := cs[i]
	card2 := cs[j]
	return card1.Number < card2.Number
}

func (cs Cards) Ascend() {
	sort.Sort(cs)
}

func (cs Cards) String() string {
	s := ""
	for _, card := range cs {
		s += card.String()
	}
	return s
}
