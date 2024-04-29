package progress

import (
	"fight-landlords/model"
	"fight-landlords/util"
)

var CurrentRound *model.Round

func initial() {
	round := begin()
	deal(round)
	dealLandlordCards(round)
	CurrentRound = round
}

func begin() *model.Round {
	players := createPlayers()
	round := model.NewRound(players)
	return round
}

func createPlayers() []*model.Player {
	player1 := &model.Player{Nickname: "player1", CurrentInning: model.NewInning()}
	player2 := &model.Player{Nickname: "player2", CurrentInning: model.NewInning()}
	player3 := &model.Player{Nickname: "player3", CurrentInning: model.NewInning()}
	players := []*model.Player{player1, player2, player3}
	return players
}

func deal(round *model.Round) {
	players := round.Players
	var top *model.Card
	// 抽牌17轮
	for range 17 {
		for j := range 3 {
			top = round.Deck.Pop()
			dealToPlayer(top, players[j])
		}
	}
}

func dealToPlayer(card *model.Card, player *model.Player) {
	havingCards := &player.CurrentInning.HavingCards
	*havingCards = append(*havingCards, card)
	sequence := player.CurrentInning.Sequence
	*sequence += toSequence(card)
}

func toSequence(card *model.Card) string {
	return util.Unhash(card.Number)
}

// 第一个返回值是地主,第二个返回值是农民数组
func apartLandlordFarmers(players []*model.Player) (*model.Player, []*model.Player) {
	var landlord *model.Player
	for i, player := range players {
		for _, card := range player.CurrentInning.HavingCards {
			if card.IsLandlord {
				landlord = player
				landlord.CurrentInning.IsLandlord = true // 多余
				farmers := append(players[:i], players[i+1:]...)
				return landlord, farmers
			}
		}
	}
	panic("No Landlord")
}

func dealLandlordCards(round *model.Round) {
	landlord, _ := apartLandlordFarmers(round.Players)
	landlord.CurrentInning.IsLandlord = true
	lc := &landlord.CurrentInning.HavingCards
	// 地主额外拿走最后的3张
	seq := landlord.CurrentInning.Sequence
	for range 3 {
		card := round.Deck.Pop()
		*seq += util.Unhash(card.Number)
		*lc = append(*lc, card)
	}
}
