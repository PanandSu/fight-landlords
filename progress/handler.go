package progress

import (
	"fight-landlords/model"
	"fight-landlords/rule"
)

// Handle_n 表示处理n张牌的情况,handel0表示处理其他张数的牌

func Handle1(seq string) any {
	return &rule.Single{Card: toCard([]rune(seq)[0])}
}

func Handle2(seq string) any {
	if IsPair(seq) {
		s := seq[0]
		card := toCard(rune(s))
		return &rule.Pair{Cards: [2]*model.Card{card, card}}
	} else if IsRocket(seq) {
		return &rule.Rocket
	} else {
		panic("Invalid pair sequence")
	}
}

func Handle3(seq string) any {
	if IsThree(seq) {
		s := seq[0]
		card := toCard(rune(s))
		return &rule.Three{Triplet: [3]*model.Card{card, card, card}}
	} else {
		panic("Invalid three sequence")
	}
}

func Handle4(seq string) any {
	if IsBomb(seq) {
		s := seq[0]
		card := toCard(rune(s))
		return &rule.Bomb{Cards: [4]*model.Card{card, card, card, card}}
	} else if IsThreeWithOne(seq) {
		var a, b int
		if IsThreeWithOneType1(seq) {
			a, b = 0, 3
		} else if IsThreeWithOneType2(seq) {
			a, b = 1, 0
		}
		main := Handle3(seq[a : a+3]).(*rule.Three)
		slave := Handle1(seq[b : b+1]).(*rule.Single)
		return &rule.ThreeWithOne{Main: main, Slave: slave}
	} else {
		panic("Invalid bomb or 3with1 sequence")
	}
}

func Handle5(seq string) any {
	if IsThreeWithPair(seq) {
		var a, b int
		if IsThreeWithPairType1(seq) {
			a, b = 0, 3
		} else if IsThreeWithPairType2(seq) {
			a, b = 2, 0
		}
		main := Handle3(seq[a : a+3]).(*rule.Three)
		slave := Handle2(seq[b : b+2]).(*rule.Pair)
		return &rule.ThreeWithPair{Main: main, Slave: slave}
	} else if IsStraight(seq) {
		var cards []*model.Card
		for _, v := range seq {
			card := toCard(v)
			cards = append(cards, card)
		}
		return &rule.Straight{Cards: cards, Len: 5}
	} else {
		panic("Invalid straight sequence")
	}
}

// 222234  0,4,5
// 455556  1,5,0
// 567777  2,0,1
func Handle6(seq string) any {
	if IsFourWithTwo(seq) {
		var a, b, c int
		f := func(n int) {
			a, b, c = n, (n+4)%6, (n+5)%6
		}
		if IsFourWithTwoType1(seq) {
			f(0)
		} else if IsFourWithTwoType2(seq) {
			f(1)
		} else if IsFourWithTwoType3(seq) {
			f(2)
		}
		main := Handle4(seq[a : a+4]).(*rule.Bomb)
		slave1 := Handle1(seq[b : b+1]).(*rule.Single)
		slave2 := Handle1(seq[c : c+1]).(*rule.Single)
		return &rule.FourWithTwo{Main: main, Slave1: slave1, Slave2: slave2}
	} else if IsStraight(seq) {
		var cards []*model.Card
		for _, v := range seq {
			card := toCard(v)
			cards = append(cards, card)
		}
		return &rule.Straight{Cards: cards, Len: 6}
	} else if IsConsecutivePair(seq) {
		pairs := make([]*rule.Pair, 0)
		for i := 0; i < len(seq); i = i + 2 {
			pair := [2]*model.Card{}
			pair[0] = toCard(rune(seq[i]))
			pair[1] = toCard(rune(seq[i+1]))
			pairs = append(pairs, &rule.Pair{Cards: pair})
		}
		return &rule.ConsecutivePair{Pairs: pairs, Len: 3}
	} else {
		panic("Invalid 4with2 or straight sequence")
	}
}

func Handle8(seq string) any {
	if IsAirplane(seq) {
		var leftMain, rightMain *rule.Three
		var leftSlave, rightSlave *rule.Single
		airplane := &rule.Airplane{}
		var a, b, c, d int
		f := func(i int) {
			a, b, c, d = i, i+3, (i+6)%8, (i+7)%8
		}
		if IsAirplaneType1(seq) {
			f(0)
		} else if IsAirplaneType2(seq) {
			f(1)
		} else if IsAirplaneType3(seq) {
			f(2)
		}
		leftMain = Handle3(seq[a : a+3]).(*rule.Three)
		rightMain = Handle3(seq[b : b+3]).(*rule.Three)
		leftSlave = Handle1(seq[c : c+1]).(*rule.Single)
		rightSlave = Handle1(seq[d : d+1]).(*rule.Single)
		airplane.Mains = [2]*rule.Three{leftMain, rightMain}
		airplane.Slaves = [2]*rule.Single{leftSlave, rightSlave}
		return airplane
	} else if IsStraight(seq) {
		var cards []*model.Card
		for _, v := range seq {
			card := toCard(v)
			cards = append(cards, card)
		}
		return &rule.Straight{Cards: cards, Len: 8}
	} else if IsConsecutivePair(seq) {
		pairs := make([]*rule.Pair, 0)
		for i := 0; i < len(seq); i = i + 2 {
			pair := [2]*model.Card{}
			pair[0] = toCard(rune(seq[i]))
			pair[1] = toCard(rune(seq[i+1]))
			pairs = append(pairs, &rule.Pair{Cards: pair})
		}
		return &rule.ConsecutivePair{Pairs: pairs, Len: 4}
	} else {
		panic("Invalid airplane or straight or consecutive_pair sequence")
	}
	return nil
}

func Handle0(seq string) any {
	if IsStraight(seq) {
		var cards []*model.Card
		for _, v := range seq {
			card := toCard(v)
			cards = append(cards, card)
		}
		return &rule.Straight{Cards: cards, Len: len(seq)}
	} else if IsConsecutivePair(seq) {
		if len(seq)%2 == 1 {
			panic("Invalid consecutive pair sequence")
		}
		pairs := make([]*rule.Pair, 0)
		for i := 0; i < len(seq); i = i + 2 {
			pair := [2]*model.Card{}
			pair[0] = toCard(rune(seq[i]))
			pair[1] = toCard(rune(seq[i+1]))
			pairs = append(pairs, &rule.Pair{Cards: pair})
		}
		return &rule.ConsecutivePair{Pairs: pairs, Len: len(seq) / 2}
	} else {
		panic("Invalid straight sequence")
	}
}
