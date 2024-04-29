package progress

import (
	"fight-landlords/model"
	"fight-landlords/rule"
	"strconv"
)

func IsPair(seq string) bool {
	return seq[0] == seq[1]
}

func IsThree(seq string) bool {
	return seq[0] == seq[1] && seq[1] == seq[2]
}

func IsBomb(seq string) bool {
	return seq[0] == seq[1] && seq[1] == seq[2] && seq[2] == seq[3]
	//return seq[0]^seq[1]^seq[2]^seq[3] == 0
}

func IsThreeWithOne(seq string) bool {
	return IsThreeWithOneType1(seq) || IsThreeWithOneType2(seq)
}

func IsThreeWithOneType1(seq string) bool {
	return IsThree(seq[0:3])
}
func IsThreeWithOneType2(seq string) bool {
	return IsThree(seq[1:4])
}

// 33222 55566
func IsThreeWithPair(seq string) bool {
	return IsThreeWithPairType1(seq) || IsThreeWithPairType2(seq)
}
func IsThreeWithPairType1(seq string) bool {
	return IsThree(seq[0:3]) && IsPair(seq[3:5])
}
func IsThreeWithPairType2(seq string) bool {
	return IsPair(seq[0:2]) && IsThree(seq[2:5])
}

func IsFourWithTwo(seq string) bool {
	return IsFourWithTwoType1(seq) || IsFourWithTwoType2(seq) || IsFourWithTwoType3(seq)
}

func IsFourWithTwoType1(seq string) bool {
	return IsBomb(seq[0:4])
}
func IsFourWithTwoType2(seq string) bool {
	return IsBomb(seq[1:5])
}
func IsFourWithTwoType3(seq string) bool {
	return IsBomb(seq[2:6])
}
func IsRocket(seq string) bool {
	return seq == rule.Rocket
}
func IsConsecutivePair(seq string) bool {
	if len(seq)%2 == 1 {
		panic("Consecutive Pair must be even")
		return false
	}
	for i := 0; i < len(seq)-2; i = i + 2 {
		if !IsPair(seq[i : i+2]) {
			return false
		}
		num1 := toCard(rune(seq[i])).Number
		num2 := toCard(rune(seq[i+2])).Number
		if !(num1 == num2-1) {
			return false
		}
	}
	if !IsPair(seq[len(seq)-2:]) {
		return false
	}
	return true
}
func IsStraight(seq string) bool {
	n := len(seq)
	if !(n >= 5 && n <= 12) {
		return false
	}
	first := seq[0]
	// 大于3
	if !(toCard(rune(first)).Number >= 3) {
		return false
	}
	last := seq[n-1]
	// 小于A
	if !(toCard(rune(last)).Number <= 14) {
		return false
	}
	flag := true
	for i := 1; i < n; i++ {
		j := i - 1
		iNum := toCard(rune(seq[i])).Number
		jNum := toCard(rune(seq[j])).Number
		if iNum != jNum+1 {
			flag = false
		}
	}
	return flag
}

func IsAirplane(seq string) bool {
	// airplane共有3中情况,满足一个就是airplane
	return IsAirplaneType1(seq) || IsAirplaneType2(seq) || IsAirplaneType3(seq)
}

func IsAirplaneType1(seq string) bool {
	//"33344456"
	return when(seq, 0)
}

func IsAirplaneType2(seq string) bool {
	//"34445556"
	return when(seq, 1)
}

func IsAirplaneType3(seq string) bool {
	//"34555666"
	return when(seq, 2)
}

// opt=0,1,2
func when(seq string, opt int) bool {
	var arr [4]int
	var a, b, c, d int
	f := func(i int) {
		a, b, c, d = i, i+3, (i+6)%8, (i+7)%8
	}
	switch opt {
	case 0, 1, 2:
		f(opt)
		arr = [4]int{a, b, c, d}
	default:
		panic("opt must be 0 or 1 or 2")
	}
	return cond(seq, arr)
}

func cond(seq string, arr [4]int) bool {
	a, b, c, d := arr[0], arr[1], arr[2], arr[3]
	i, j, k, l := seq[a], seq[b], seq[c], seq[d]
	iNum := toCard(rune(i)).Number
	jNum := toCard(rune(j)).Number
	kNum := toCard(rune(k)).Number
	lNum := toCard(rune(l)).Number
	if !IsThree(seq[a : a+3]) {
		return false
	}
	if !IsThree(seq[b : b+3]) {
		return false
	}
	if !(iNum == jNum-1) {
		return false
	}
	if !(kNum != jNum) {
		return false
	}
	if !(lNum != kNum) {
		return false
	}
	return true
}

func toCard(r rune) *model.Card {
	switch r {
	case '3', '4', '5', '6', '7', '8', '9':
		n, _ := strconv.Atoi(string(r))
		return &model.Card{Name: string(r), Number: n}
	case 'X':
		return &model.Card{Name: string(r), Number: 10}
	case 'J':
		return &model.Card{Name: string(r), Number: 11}
	case 'Q':
		return &model.Card{Name: string(r), Number: 12}
	case 'K':
		return &model.Card{Name: string(r), Number: 13}
	case 'A':
		return &model.Card{Name: string(r), Number: 14}
	case '2':
		return &model.Card{Name: string(r), Number: 15}
	case 'L':
		return &model.Card{Name: "Little-Joker", Number: 16}
	case 'B':
		return &model.Card{Name: "Big-Joker", Number: 17}
	}
	panic("Invalid rule sequence")
	return nil
}
