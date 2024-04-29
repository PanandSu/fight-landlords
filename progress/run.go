package progress

import (
	"fight-landlords/model"
	"fight-landlords/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	card1  string
	card2  string
	card3  string
	last   string
	master string
)

func Run() {
	initial() //正确
	player1 := CurrentRound.Players[0]
	player2 := CurrentRound.Players[1]
	player3 := CurrentRound.Players[2]
	pc1 := &player1.CurrentInning.HavingCards
	pc2 := &player2.CurrentInning.HavingCards
	pc3 := &player3.CurrentInning.HavingCards
	fmt.Println("欢乐斗地主,游戏开始!!🤗🤗")

	showCards2(player1)
	showCards2(player2)
	showCards2(player3)

	fmt.Println("-------------------------")
	landlord, farmers := apartLandlordFarmers(CurrentRound.Players)
	assert.Equal(&testing.T{}, len(farmers), 2)
	fmt.Printf("地主为%s\n", landlord.Nickname)
	fmt.Printf("农民为%s和%s\n", farmers[0].Nickname, farmers[1].Nickname)
	master = landlord.Nickname
	for {
		fmt.Println("新回合开始喽----------!")

		perform(landlord, &card1)
		if len(*pc1) <= 10 {
			win(player1)
			break
		}
		perform(farmers[0], &card2)
		if len(*pc2) <= 10 {
			win(player2)
			break
		}
		perform(farmers[1], &card3)
		if len(*pc3) <= 10 {
			win(player3)
			break
		}
	}
	fmt.Println("游戏结束,Have a good time!!🤗🤗")
}

// card必须为*string形式,因为每次调用perform函数都修改了card
func perform(player *model.Player, card *string) {
	fmt.Printf("%s出牌:(%#v)\n", player.Nickname, util.SortedStr(*player.CurrentInning.Sequence))
	if master == player.Nickname {
		positive(player, card)
	} else {
		negative(player, card)
	}
}

func positive(player *model.Player, card *string) {
	posValidInput(player, card)
	RemovePlayerCards(player, *card)
	fmt.Printf("%s出的牌%#v,剩余的牌:%#v\n",
		player.Nickname, *card, *player.CurrentInning.Sequence)
	master = player.Nickname
}

func checkInput(input string) bool {
	s := "A23456789XJQKLB"
	for _, v := range input {
		index := -1
		for i, char := range s {
			if v == char {
				index = i
				break
			}
		}
		if index == -1 {
			return false
		}
	}
	return true
}

func posValidInput(player *model.Player, card *string) {
	fmt.Print(">>> ")
	_, _ = fmt.Scanln(card)
	if !checkInput(*card) {
		fmt.Println("只能输入A23456789XJQKLB!")
		posValidInput(player, card)
	}
	*card, last = AscendStr(*card), AscendStr(last)
	if !AffordDeal(player, *card) {
		fmt.Println("不能出没有的牌👿👿,重新出牌")
		posValidInput(player, card)
	}
	if !IsValidType(*card) {
		fmt.Println("出的牌不符合规则!😠😠,重新出牌")
		posValidInput(player, card)
	}
	return
}

func negValidInput(player *model.Player, card *string) bool {
	fmt.Println("要or不要? 如果要,输入出的牌,不要就输入0")
	fmt.Print(">>> ")
	_, _ = fmt.Scanln(card)
	if *card == "0" {
		fmt.Printf("%s放弃出牌\n", player.Nickname)
		return false
	}
	if !checkInput(*card) {
		fmt.Println("只能输入A23456789XJQKLB!")
		negValidInput(player, card)
	}
	*card, last = AscendStr(*card), AscendStr(last)
	if !AffordDeal(player, *card) {
		fmt.Println("不能出没有的牌👿👿,重新出牌")
		negValidInput(player, card)
	}
	if !(IsValidType(*card) && IsValidType(last)) {
		fmt.Println("出的牌不符合规则!😠😠,重新出牌")
		negValidInput(player, card)
	}
	return true
}

func negative(player *model.Player, card *string) {
	if !negValidInput(player, card) {
		return
	}
	n := control(*card, last)
	if n == 0 {
		return
	} else if n == 1 {
		negative(player, card)
	} else if n == 2 {
		RemovePlayerCards(player, *card)
		fmt.Printf("%s出的牌%#v,剩余的牌:%#v\n",
			player.Nickname, *card, *player.CurrentInning.Sequence)
		master = player.Nickname
	}
}
