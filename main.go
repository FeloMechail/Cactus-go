package main

import (
	"fmt"
	"math/rand"
)

func fullDeck() []uint32 {
	/*
		Cactus Kev's Poker Hand Evaluator
		4-byte Card Format:
		 * A card is a 32 bit integer with the following bits:
		 * xxxAKQJT 98765432 CDHSrrrr xxPPPPPP
		 * r = rank of card (0 duece - 12 Ace)
		 * 13 bits each represents a specific card rank and exactly one is set in each card
		 * CDHS = suit of card
		 * 6 bit to store the prime number of the card rank

	*/

	rankPrimes := []uint32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	suits := []uint32{8, 4, 2, 1} // CDHS
	cards := make([]uint32, 0, 52)

	for rank := 0; rank < 13; rank++ {
		for _, suit := range suits {
			card := (rankPrimes[rank]) | (uint32(rank) << 8) | (suit << 12) | ((1 << rank) << 16)
			cards = append(cards, card)
		}
	}
	return cards
}

func getCardName(card uint32) string {
	var rank = byte((card >> 8) & 0xF) // 0 represents duece, 12 represents Ace
	var suit = byte((card >> 12) & 0xF)

	//map suits num to string
	suits := map[byte]string{
		8: "♣",
		4: "♦",
		2: "♥",
		1: "♠",
	}

	//map rank num to string
	ranks := map[byte]string{
		0: "2", 1: "3", 2: "4", 3: "5", 4: "6", 5: "7", 6: "8", 7: "9", 8: "T", 9: "J", 10: "Q", 11: "K", 12: "A",
	}

	return ranks[rank] + suits[suit]
}

func flush(cards []uint32) bool {
	suit := uint32(0xF000)
	for _, card := range cards {
		suit &= card
	}

	return suit != 0
}

func flushBitPattern(cards []uint32) uint32 {
	total := uint32(0)
	for _, card := range cards {
		total |= card
	}

	shiftedTotal := total >> 16

	return shiftedTotal
}

func primeMultiplicant(cards []uint32) uint32 {
	multiplicant := uint32(1)
	for _, card := range cards {
		multiplicant *= card & 0xFF
	}

	return multiplicant
}

func fastRank(cards []uint32) uint32 {
	u := primeMultiplicant(cards)
	u += 0xe91aaa35
	u ^= u >> 16
	u += u << 8
	u ^= u >> 4
	a := (u + (u << 2)) >> 19
	return a ^ uint32(HashAdjust[(u>>8)&0x1ff])
}

func handRank(cards []uint32) string {
	var rank int = 0

	if flush(cards) {

		rank = Flushes[int(flushBitPattern(cards))]

	} else if FiveUniqueCards[int(flushBitPattern(cards))] != 0 {

		rank = FiveUniqueCards[int(flushBitPattern(cards))]

	} else {
		rank = HashValues[fastRank(cards)]
	}

	if rank > 6185 {
		return "High Card"
	} else if rank > 3325 {
		return "One Pair"
	} else if rank > 2467 {
		return "Two Pair"
	} else if rank > 1609 {
		return "Three of a Kind"
	} else if rank > 1599 {
		return "Straight"
	} else if rank > 322 {
		return "Flush"
	} else if rank > 166 {
		return "Full House"
	} else if rank > 10 {
		return "Four of a Kind"
	} else {
		return "Straight Flush"
	}

}

func drawHand(cards []uint32) []uint32 {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards[:5]
}

func main() {
	// //draws random 5 cards from a full deck
	// deck := fullDeck()
	// hands := drawHand(deck)

	// straight flush
	hands := []uint32{268471337, 16812055, 33589533, 67144223, 134253349}

	for _, card := range hands {
		fmt.Print(getCardName(card) + " ")
	}

	fmt.Println("\n" + handRank(hands))

}
