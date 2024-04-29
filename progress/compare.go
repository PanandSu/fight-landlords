package progress

import (
	"fight-landlords/rule"
)

// s1>s2
func BiggerThan(s1, s2 string) bool {
	self := SwitchToType(s1)
	opponent := SwitchToType(s2)
	switch self.(type) {
	case *rule.Single:
		self := self.(*rule.Single)
		opponent := opponent.(*rule.Single)
		return compareSingle(self, opponent)
	case *rule.Pair:
		self := self.(*rule.Pair)
		opponent := opponent.(*rule.Pair)
		return comparePair(self, opponent)
	case *rule.Three:
		self := self.(*rule.Three)
		opponent := opponent.(*rule.Three)
		return compareThree(self, opponent)
	case *rule.Bomb:
		self := self.(*rule.Bomb)
		opponent := opponent.(*rule.Bomb)
		return compareBomb(self, opponent)
	case *rule.ThreeWithOne:
		self := self.(*rule.ThreeWithOne)
		opponent := opponent.(*rule.ThreeWithOne)
		return compareThreeWithOne(self, opponent)
	case *rule.ThreeWithPair:
		self := self.(*rule.ThreeWithPair)
		opponent := opponent.(*rule.ThreeWithPair)
		return compareThreeWithPair(self, opponent)
	case *rule.FourWithTwo:
		self := self.(*rule.FourWithTwo)
		opponent := opponent.(*rule.FourWithTwo)
		return compareFourWithTwo(self, opponent)
	case *rule.ConsecutivePair:
		self := self.(*rule.ConsecutivePair)
		opponent := opponent.(*rule.ConsecutivePair)
		return compareConsecutivePair(self, opponent)
	case *rule.Straight:
		self := self.(*rule.Straight)
		opponent := opponent.(*rule.Straight)
		return compareStraight(self, opponent)
	case *rule.Airplane:
		self := self.(*rule.Airplane)
		opponent := opponent.(*rule.Airplane)
		return compareAirplane(self, opponent)
	default:
		panic("Invalid type!!")
	}
	return false
}

func compareSingle(a, b *rule.Single) bool {
	return a.Card.Number > b.Card.Number
}

func comparePair(a, b *rule.Pair) bool {
	return a.Cards[0].Number > b.Cards[0].Number
}

func compareThree(a, b *rule.Three) bool {
	return a.Triplet[0].Number > b.Triplet[0].Number
}

func compareBomb(a, b *rule.Bomb) bool {
	return a.Cards[0].Number > b.Cards[0].Number
}

func compareThreeWithOne(a, b *rule.ThreeWithOne) bool {
	return a.Main.Triplet[0].Number > b.Main.Triplet[0].Number
}

func compareThreeWithPair(a, b *rule.ThreeWithPair) bool {
	return a.Main.Triplet[0].Number > b.Main.Triplet[0].Number
}

func compareFourWithTwo(a, b *rule.FourWithTwo) bool {
	return a.Main.Cards[0].Number > b.Main.Cards[0].Number
}
func compareConsecutivePair(a, b *rule.ConsecutivePair) bool {
	if a.Len != b.Len {
		panic("Consecutive pair length mismatch!")
		return false
	}
	return a.Pairs[0].Cards[0].Number > b.Pairs[0].Cards[0].Number
}
func compareAirplane(a, b *rule.Airplane) bool {
	return a.Mains[0].Triplet[0].Number > b.Mains[0].Triplet[0].Number
}

func compareStraight(a, b *rule.Straight) bool {
	if a.Len != b.Len {
		panic("Straight lengths do not match!")
		return false
	}
	return a.Cards[0].Number > b.Cards[0].Number
}
