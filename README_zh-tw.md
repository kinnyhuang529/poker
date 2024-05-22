# poker

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## 介紹
`poker` 是一個撲克牌算牌器，使用 Cactus Kev's Poker Hand Evaluator 演算法。 它可以計算玩家手中的撲克牌組合併給出相應的牌力及牌型。

## 先決條件
確保您已經安裝了Go環境。

## 安裝
```bash
go get github.com/kinnyhuang529/poker
```

## 配置
在 Go 程式碼中匯入此套件：
```bash
import "github.com/kinnyhuang529/poker"
```

## 範例
```go
package main

import (
	"fmt"
	"github.com/kinnyhuang529/poker"
)

func main() {
	//建立排組
	deck := poker.NewDeck()

	//洗牌
	deck.Shuffle()

	//五張牌
	five, err := deck.Draw(5)
	//five: [s3 d2 sK c8 s9]

	fiveCardType, fiveCardStrength, fiveBestHand, err := poker.Evaluator(five)
	//fiveCardType: 10, fiveCardStrength: 6952, fiveBestHand: [s3 d2 sK c8 s9]

	//六張牌
	six, err := deck.Draw(6)
	//six: [c7 c5 dK h6 cA hK]

	sixCardType, sixCardStrength, sixBestHand, err := poker.Evaluator(six)
	//sixCardType: 9, sixCardStrength: 3586, sixBestHand: [c7 dK h6 cA hK]

	//七張牌
	seven, err := deck.Draw(7)
	//seven: [c6 cK h8 dQ s8 cT hT]
	sevenCardType, sevenCardStrength, sevenBestHand, err := poker.Evaluator(seven)
	//sevenCardType: 8, sevenCardStrength: 2942, sevenBestHand: [cK h8 s8 cT hT]
}
```
### cardType 說明

- 0: 例外情況
- 1: 皇家同花順
- 2: 同花順
- 3: 四條/鐵支
- 4: 葫蘆
- 5: 同花
- 6: 順子
- 7: 三條
- 8: 兩對
- 9: 一對
- 10: 高牌

### cardStrength 說明
範圍是1到7462，數字越小牌力越強。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)