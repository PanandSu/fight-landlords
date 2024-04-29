package progress

import (
	"fight-landlords/model"
	"fight-landlords/rule"
	"fight-landlords/util"
	"fmt"
	"reflect"
)

func win(player *model.Player) {
	fmt.Printf("%s获胜!\n", player.Nickname)
	if player.CurrentInning.IsLandlord {
		fmt.Println("地主获胜!!")
	} else {
		fmt.Println("农民获胜!!")
	}
}

// 管住,控制住,card,last已为相同类型
// 返回值为0,1,2. 0表示永不可能管住,1表示没管住,重新出牌有可能,2表示管住了
func control(card string, last string) int {
	if last == rule.Rocket {
		fmt.Println("不可能管住火箭!!")
		return 0
	}
	if card == rule.Rocket {
		return 2
	}
	type1 := SwitchToType(card)
	type2 := SwitchToType(last)
	_, ok1 := type1.(rule.Bomb)
	_, ok2 := type2.(rule.Bomb)
	if ok1 && !ok2 {
		return 2
	}
	if ok2 && !ok1 {
		fmt.Println("管不住炸弹,重新选要还是不要!!")
		return 1
	}
	if !IsSameType(card, last) {
		fmt.Println("驴唇不对马嘴,出的牌类型与上家不符!,重新出牌")
		return 1
	}
	bigger := BiggerThan(card, last)
	if !bigger {
		fmt.Printf("您出的牌根本管不住上家的牌,不要乱出,上家的牌为%s\n", last)
		fmt.Println("重新选要还是不要,要就好好出!!")
		return 1
	}
	return 2
}

func IsValidType(s string) bool {
	toType := SwitchToType(s)
	switch toType.(type) {
	case *string, *rule.Single, *rule.Pair, *rule.Three,
		*rule.Bomb, *rule.ThreeWithOne, *rule.FourWithTwo,
		*rule.Airplane, *rule.Straight:
		return true
	default:
		return false
	}
}

func IsSameType(a, b string) bool {
	t1 := SwitchToType(a)
	t2 := SwitchToType(b)
	type1 := reflect.TypeOf(t1)
	type2 := reflect.TypeOf(t2)
	if type1 != type2 {
		return false
	}
	return true
}

func showCards(player *model.Player) {
	fmt.Printf("%s的牌:\n", player.Nickname)
	for _, card := range player.CurrentInning.HavingCards {
		fmt.Printf("%#v	", card)
	}
}

func showCards2(player *model.Player) {
	fmt.Printf("%s的牌:\n", player.Nickname)
	fmt.Println(util.SortedStr(*player.CurrentInning.Sequence))
}

// 能否负担得起要出牌,cards能否由player.CurrentInning.Sequence中的牌组成
func AffordDeal(player *model.Player, cards string) bool {
	seq := player.CurrentInning.Sequence
	seqRunes := []rune(*seq)
	cardRunes := []rune(cards)
	return util.ContainRunes(cardRunes, seqRunes)
}

// 出过的牌要remove
func RemovePlayerCards(player *model.Player, cards string) {
	seq := player.CurrentInning.Sequence
	*seq = AscendStr(*seq)
	cards = AscendStr(cards)
	*seq = string(util.RemoveRunes([]rune(*seq), []rune(cards)))
	havingCards := &(player.CurrentInning.HavingCards)
	var Cards []*model.Card
	for _, v := range *seq {
		Cards = append(Cards, toCard(v))
	}
	*havingCards = Cards
	last = cards
}
