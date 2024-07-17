# Cactus-go

cactus-go is a Go implementation of a poker hand evaluator. It is based on the [Cactus Kev's (Kevin Suffercool) Poker Hand Evaluator](http://suffe.cool/poker/evaluator.html) algorithm. 

## Table of Contents

- [Cactus-go](#cactus-go)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
  - [Functions](#functions)
  - [Examples](#examples)
    - [Output](#output)

## Usage

1. Clone the repository:
    ```sh
    git clone https://github.com/FeloMechail/Cactus-go.git
    ```
2. Navigate to the project directory:
    ```sh
    cd Cactus-go
    ```
3. Run the project:
    ```sh
    go run .
    ```

## Functions

- `fullDeck`: The fullDeck function generates a full deck of 52 cards using Cactus Kev's Poker Hand Evaluator's 4-byte card format. Each card is represented as a 32-bit integer with specific bits assigned to different card properties.

    Card Format:
    ```plaintext
    +--------+--------+--------+--------+
    |xxxAKQJT|98765432|cdhsrrrr|xxpppppp|
    +--------+--------+--------+--------+
    ```
    * `p` = prime number of rank (deuce=2,trey=3,four=5,five=7,...,ace=41)
    * `r` = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
    * `cdhs` = suit of card
    * `AKQJT987654` = bit turned on depending on rank of card
    * `xxx` = unused



- `getCardName`: The getCardName function returns the name of the card based on the card's rank and suit.

- `drawHand`: The drawHand function draws a hand of 5 cards from the deck of cards.

- `handRank`: The handRank function evaluates the rank of the hand based on Kevin Suffercool's Poker Hand Evaluator algorithm.

## Examples

```go
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
```

### Output
```plaintext
A♣ T♣ J♣ Q♣ K♣ 
Straight Flush
```

