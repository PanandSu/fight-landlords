# 斗地主 fight-landlords

# Go + MySQL + Redis +...

# 有时会(发牌)运行错误,原因还不知道(等着大佬带我😭😭)

### rules

54张牌,2张大小王,4个[1,13]

牌:
梅花,方块,红心,黑桃
一副牌:deck,suit
梅花(Club)为C，方块(Diamond)为D，红心(Hearts)为H，黑桃(Spade)为S。

3人玩法:
随机使牌乱序,约定一张地主牌

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

# 对单文件中单个函数测试

在IDE左侧目录结构中右键,选择在Explorer(文件资源管理器)中打开,然后在顶部路径栏,键入powershell,
可以直接进入powershell,类似键入cmd.
然后在powershell中执行命令

```shell
go test -v src.go src_test.go -run <src_test中的Test函数>
go test -v xxx_test.go -run <src_test中的Test函数> 或者这样 
```


