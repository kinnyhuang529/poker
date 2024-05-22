package poker

type Card int32

var (
	suitKey    = []uint8{'s', 'h', 'd', 'c'}
	suitValue  = []int32{1, 2, 4, 8}
	rankKey    = []uint8{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	rankValue  = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	primeValue = []int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	suits      = map[uint8]int32{}
	ranks      = map[uint8]int32{}
	primes     = map[uint8]int32{}
)

var (
	suitStr = "xshxdxxxc"
	rankStr = "23456789TJQKA"
)

func init() {
	for i, s := range suitKey {
		suits[s] = suitValue[i]
	}
	for i, r := range rankKey {
		ranks[r] = rankValue[i]
		primes[r] = primeValue[i]
	}
}
func NewCard(card string) Card {
	//xxxAKQJT 98765432 CHDSrrrr xxpppppp
	prime := primes[card[1]]
	rank := ranks[card[1]] << 8
	suit := suits[card[0]] << 12
	bit := int32(1) << ranks[card[1]] << 16
	return Card(bit | suit | rank | prime)
}

func (c Card) Prime() int32 {
	return int32(c) & 0x3F
}

func (c Card) Rank() int32 {
	return (int32(c) >> 8) & 0xF
}

func (c Card) Suit() int32 {
	return (int32(c) >> 12) & 0xF
}

func (c Card) BitRank() int32 {
	return (int32(c) >> 16) & 0x1FFF
}

func (c Card) String() string {
	return string(suitStr[c.Suit()]) + string(rankStr[c.Rank()])
}
