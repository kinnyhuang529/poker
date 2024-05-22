# poker

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## 介绍
`poker` 是一个扑克牌算牌器，使用 Cactus Kev's Poker Hand Evaluator 演算法。它可以计算玩家手中的扑克牌组合并给出相应的牌力及牌型。

## 先决条件
确保您已经安装了Go环境。

## 安装
```bash
go get github.com/kinnyhuang529/poker
```

## 配置
在 Go 代码中导入此包：
```bash
import "github.com/kinnyhuang529/poker"
```

## 范例
```go
package main

import (
	"fmt"
	"github.com/kinnyhuang529/poker"
)

func main() {
	//建立牌组
	deck := poker.NewDeck()

	//洗牌
	deck.Shuffle()

	//五张牌
	five, err := deck.Draw(5)
	//five: [s3 d2 sK c8 s9]

	fiveCardType, fiveCardStrength, fiveBestHand, err := poker.Evaluator(five)
	//fiveCardType: 10, fiveCardStrength: 6952, fiveBestHand: [s3 d2 sK c8 s9]

	//六张牌
	six, err := deck.Draw(6)
	//six: [c7 c5 dK h6 cA hK]

	sixCardType, sixCardStrength, sixBestHand, err := poker.Evaluator(six)
	//sixCardType: 9, sixCardStrength: 3586, sixBestHand: [c7 dK h6 cA hK]

	//七张牌
	seven, err := deck.Draw(7)
	//seven: [c6 cK h8 dQ s8 cT hT]
	sevenCardType, sevenCardStrength, sevenBestHand, err := poker.Evaluator(seven)
	//sevenCardType: 8, sevenCardStrength: 2942, sevenBestHand: [cK h8 s8 cT hT]
}
```
### cardType 说明

- 0: 例外情况
- 1: 皇家同花顺
- 2: 同花顺
- 3: 四条/金刚
- 4: 葫芦
- 5: 同花
- 6: 顺子
- 7: 三条
- 8: 两对
- 9: 一对
- 10: 高牌

### cardStrength 说明
范围是1到7462，数字越小牌力越强。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)