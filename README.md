## 斗地主 fight-landlords

## Go + MySQL + Redis +...

## 有时会(发牌)运行错误,原因还不知道(等着大佬教教我😭😭)

### rules

54张牌,2张大小王,4个[1,13]
A,2,3,4,5,6,7,8,9,X(罗马数字10),J,Q,K,L(little joker),B(big joker)

牌:
梅花,方块,红心,黑桃
一副牌:deck,suit
梅花(Club)为C，方块(Diamond)为D，红心(Hearts)为H，黑桃(Spade)为S。

3人玩法:
随机使牌乱序,随机约定一张地主牌

每人发牌17张,留3张
3个goroutine轮流起牌,

叫地主,抢地主

地主拿走剩余的3张

每个回合:
要,不要

牌规则:
单张
对子
三张
三带一
三带二
顺子
飞机
炸弹
火箭
四带二
亲亲对

胜负规则:
先出完,胜利,
若是地主,则地主方胜利
若是农民,则农民方胜利

## 对单文件中单个函数测试

在IDE左侧目录结构中右键,选择在Explorer(文件资源管理器)中打开,然后在顶部路径栏,键入powershell,
可以直接进入powershell,类似键入cmd.
然后在powershell中执行命令

```shell
go test -v src.go src_test.go -run <src_test中的Test函数>
go test -v xxx_test.go -run <src_test中的Test函数> 或者这样 
```

## 游戏效果
```sh
欢乐斗地主,游戏开始!!🤗🤗
player1的牌:
4667899XJJQQKAA2B
player2的牌:
3334567789XKKAA22
player3的牌:
34455567889XXJJQQK2L
-------------------------
地主为player3
农民为player1和player2
新回合开始喽----------!
player3出牌:
>>> 5556
player3出的牌"5556"
player3剩余的牌:3447889XXJJQQK2L
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player2放弃出牌
新回合开始喽----------!
player3出牌:
>>> 3
player3出的牌"3"
player3剩余的牌:447889XXJJQQK2L
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> 2
player1出的牌"2"
player1剩余的牌:4667899XJJQQKAAB
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player2放弃出牌
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> L
>>> L
player3出的牌"L"
player3剩余的牌:447889XXJJQQK2
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> B
player1出的牌"B"
player1剩余的牌:4667899XJJQQKAA
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player2放弃出牌
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
>>> 6789X
player1出的牌"6789X"
player1剩余的牌:469JJQQKAA
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player2放弃出牌
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
>>> AA
player1出的牌"AA"
player1剩余的牌:469JJQQK
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> 22
player2出的牌"22"
player2剩余的牌:3334567789XKKAA
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
>>> 3433
player2出的牌"3334"
player2剩余的牌:567789XKKAA
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
>>> X98765
player2出的牌"56789X"
player2剩余的牌:7KKAA
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
>>> 7
player2出的牌"7"
player2剩余的牌:KKAA
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> 2
player3出的牌"2"
player3剩余的牌:447889XXJJQQK
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player2放弃出牌
新回合开始喽----------!
player3出牌:
>>> 88
player3出的牌"88"
player3剩余的牌:4479XXJJQQK
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> AA
不能出没有的牌👿👿,重新出牌
要or不要? 如果要,输入1,不要就输入0
>>> 1
>>> QQ
player1出的牌"QQ"
player1剩余的牌:469JJK
player2出牌:
要or不要? 如果要,输入1,不要就输入0
>>> KK
>>> KK
player2出的牌"KK"
player2剩余的牌:AA
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
>>> AA
player2出的牌"AA"
player2剩余的牌:
新回合开始喽----------!
player3出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player3放弃出牌
player1出牌:
要or不要? 如果要,输入1,不要就输入0
>>> 0
player1放弃出牌
player2出牌:
>>> p
只能输入A23456789XJQKLB!
>>> A
不能出没有的牌👿👿,重新出牌
>>> l
只能输入A23456789XJQKLB!
>>> vhbkjn
只能输入A23456789XJQKLB!
>>> cuvbdha
只能输入A23456789XJQKLB!
>>> 
```
大概就是这样玩的~



