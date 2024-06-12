package main

import (
	"fight-landlords/model"
	"fight-landlords/progress"
	"fight-landlords/rule"
	"fight-landlords/util"
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	check := func(c1, c2, c3 string) bool {
		s := c1 + c2 + c3
		sorted := progress.AscendStr(s)
		fmt.Println(sorted)
		return true
	}
	check("J798K975LKK36XQAQ", "6Q538XJ6475423X6A", "48XJ2954JQ4A22A8K")
	println(util.SortedStr("73J9579685368A4527634AA5922A4424886QJKKQQXJXQJXXKL"))

	println(util.SortedStr("A23456789XJQKLB"))
}

func TestAscendStr(t *testing.T) {
	s := "73J9579685368A4527634AA5922A4424886QJKKQQXJXQJXXKLK"
	fmt.Println(len(s))
	fmt.Println(util.SortedStr(s), len(util.SortedStr(s)))
}

func TestDeal(t *testing.T) {
	pc1 := "Q3X5A38Q94X949K75"
	pc2 := "AJKK64B2K49827J3A"
	pc3 := "X7862LAQ66Q5JX835J72"
	joker := util.SortedStr(pc1 + pc2 + pc3)
	fmt.Println(joker)
	if joker != "3333444455556666777788889999XXXXJJJJQQQQKKKKAAAA2222LB" {
		panic("error")
	}
}

// 测试通过
func TestSorted(t *testing.T) {
	s := "9JJX8Q46X274377273288Q9Q5Q69952J4ALX5K33KBK6J4KAAAX685"
	println(util.SortedStr(s))
	/*
		PS P:\Projects\GoLand\fight-landlords> go test -v super_test.go -run TestSorted
		=== RUN   TestSorted
		AAAA22223333444455556666777788889999XXXXJJJJQQQQKKKKLB
		--- PASS: TestSorted (0.00s)
		PASS
		ok      command-line-arguments  2.183s
		PS P:\Projects\GoLand\fight-landlords>
	*/
}

func TestBiggerThan(t *testing.T) {
	println(progress.BiggerThan("A", "K"))
	println(progress.BiggerThan("2", "A"))
	println(progress.BiggerThan("B", "L"))
	println(progress.BiggerThan("L", "2"))

	println(progress.BiggerThan("77", "55"))
	println(progress.BiggerThan("AA", "JJ"))
	println(progress.BiggerThan("22", "55"))

	println(progress.BiggerThan("AAA", "KKK"))
	println(progress.BiggerThan("222", "888"))

	println(progress.BiggerThan("3888", "777X"))
	println(progress.BiggerThan("3222", "7AAA"))
	println(progress.BiggerThan("2222", "AAAA"))
	println(progress.BiggerThan("AAAA", "4444"))

	println(progress.BiggerThan("6789X", "34567"))
	println(progress.BiggerThan("XJQKA", "34567"))

	println(progress.BiggerThan("9XJQKA", "345678"))
	println(progress.BiggerThan("X2222L", "79AAAA"))
	println(progress.BiggerThan("9XJQKA", "345678"))

	println(progress.BiggerThan("55566678", "33344456"))
	println(progress.BiggerThan("34AAA222", "XJKKKAAA"))

	println(progress.BiggerThan("456789XJQ", "3456789XJ"))
}

func TestSwitchToType(t *testing.T) {
	println(progress.SwitchToType("A"))
}

func TestHandle1(t *testing.T) {
	arr := []string{"A", "3", "X", "J", "Q", "K", "L", "B"}
	for _, v := range arr {
		_, ok := progress.Handle1(v).(*rule.Single)
		println(ok)
	}
}
func TestHandle2(t *testing.T) {
	arr := []string{"AA", "33", "XX", "JJ", "QQ", "KK"}
	for _, v := range arr {
		_, ok := progress.Handle2(v).(*rule.Pair)
		println(ok)
	}
	_, ok := progress.Handle2("LB").(*string)
	println(ok)
}
func TestHandle3(t *testing.T) {
	arr := []string{"AAA", "333", "XXX", "JJJ", "QQQ", "KKK"}
	for _, v := range arr {
		_, ok := progress.Handle3(v).(*rule.Three)
		println(ok)
	}
}
func TestHandle4(t *testing.T) {
	arr := []string{"AAAA", "3333", "XXXX", "JJJJ", "QQQQ", "KKKK"}
	for _, v := range arr {
		_, ok := progress.Handle4(v).(*rule.Bomb)
		println(ok)
	}
	arr2 := []string{"3332", "222A", "XXX7", "QQQL"}
	for _, v := range arr2 {
		_, ok := progress.Handle4(v).(*rule.ThreeWithOne)
		println(ok)
	}
}

func TestHandle5(t *testing.T) {
	arr := []string{"34567", "9XJQK", "XJQKA"}
	for _, v := range arr {
		_, ok := progress.Handle5(v).(*rule.Straight)
		println(ok)
	}
}

func TestHandle6(t *testing.T) {
	arr := []string{"2222XX", "999933", "KKKKQQ"}
	for _, v := range arr {
		_, ok := progress.Handle6(v).(*rule.FourWithTwo)
		println(ok)
	}
	arr2 := []string{"345678", "9XJQKA"}
	for _, v := range arr2 {
		_, ok := progress.Handle6(v).(*rule.Straight)
		println(ok)
	}
	arr3 := []string{"334455", "QQKKAA"}
	for _, v := range arr3 {
		_, ok := progress.Handle6(v).(*rule.ConsecutivePair)
		println(ok)
	}
}

func TestHandle8(t *testing.T) {
	arr := []string{"33344456", "XJJJQQQK", "46AAA222"}
	for _, v := range arr {
		_, ok := progress.Handle8(v).(*rule.Airplane)
		println(ok)
	}
	arr2 := []string{"3456789X", "789XJQKA"}
	for _, v := range arr2 {
		_, ok := progress.Handle8(v).(*rule.Straight)
		println(ok)
	}
	arr3 := []string{"33445566", "JJQQKKAA"}
	for _, v := range arr3 {
		_, ok := progress.Handle8(v).(*rule.ConsecutivePair)
		println(ok)
	}
	//_, ok := progress.Handle8("222333XJ").(*rule.Airplane)
	//println(ok)
}

func TestHandle0(t *testing.T) {
	arr := []string{"3456789", "56789XJQKA"}
	for _, v := range arr {
		_, ok := progress.Handle0(v).(*rule.Straight)
		println(ok)
	}
}

func TestRemovePlayerCards(t *testing.T) {
	seq := "45566677889JJQKKAAAB"
	inning := &model.Inning{Sequence: &seq}
	p := &model.Player{CurrentInning: inning}
	progress.RemovePlayerCards(p, "88")
	println(*p.CurrentInning.Sequence)
}

func TestRemoveRunes(t *testing.T) {
	rs := util.RemoveRunes([]rune("45566677889JJQKKAAAB"), []rune("88"))
	println(string(rs))
}

func TestSortedStr(t *testing.T) {
	println(util.SortedStr("AA234435BK9JJ876XJQQ"))
}
