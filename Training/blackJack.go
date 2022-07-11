package Training

import (
	"fmt"
	"math/rand"
)

//Module 6: Black Jack
type card struct {
	val   int // 1 = ace, 10 = jack, 11 = queen, 12 = king
	suite string
	num   int
}

type deck struct {
	Deck []card
	pos  int
}

type player struct {
	score int
	cards []card
	done  bool
}

func generateDeck() []card {
	var cards []card
	for ii := 0; ii < 52; ii++ {
		cc := card{}
		cc.num = ii
		cc.val = (ii % 13) + 1
		switch ii / 13 {
		case 0:
			cc.suite = "diamond"
		case 1:
			cc.suite = "club"
		case 2:
			cc.suite = "heart"
		case 3:
			cc.suite = "spade"
		}
		cards = append(cards, cc)
	}

	return cards
}

func (d *deck) shuffleDeck(cards []card) {
	if len(cards) != 52 {
		fmt.Println("Not all cards are present")
		return
	}

	var c []card
	d.Deck = c

	for ii := 0; ii < 52; ii++ {
		xx := rand.Intn(len(cards))
		d.Deck = append(d.Deck, cards[xx])
	}
}

func (d *deck) draw() card {
	cc := d.Deck[d.pos]
	d.pos += 1

	return cc
}

func (p *player) updateScore() {
	score := 0
	numAces := 0
	for _, cc := range p.cards { //ignore all aces
		switch cc.val {
		case 1:
			numAces += 1
		case 11:
			score += 10
		case 12:
			score += 10
		case 13:
			score += 10
		default:
			score += cc.val
		}
	}

	for ii := 0; ii < numAces; ii++ {
		if score+11+numAces-1 <= 21 {
			score += 11
		} else {
			score += 1
		}
	}

	p.score = score
}

func BlackJack() {

	// for _, ii := range cards {
	// 	fmt.Println(ii)
	// }

	// fmt.Println()

	// for _, ii := range Deck.Deck {
	// 	fmt.Println(ii)
	// }

	input := ""
	for true {
		cards := generateDeck()
		Deck := deck{}
		Deck.pos = 0
		Deck.shuffleDeck(cards)

		dealer := player{}
		dealer.score = 0
		dealer.done = false
		var hand1 []card
		dealer.cards = hand1

		player1 := player{}
		player1.score = 0
		player1.done = false
		var hand2 []card
		player1.cards = hand2

		for !dealer.done || !player1.done {
			if dealer.score == 0 && player1.score == 0 {

				cc := Deck.draw()
				player1.cards = append(player1.cards, cc)

				cc = Deck.draw()
				dealer.cards = append(dealer.cards, cc)

				cc = Deck.draw()
				player1.cards = append(player1.cards, cc)

				cc = Deck.draw()
				dealer.cards = append(dealer.cards, cc)

				dealer.updateScore()
				player1.updateScore()
				fmt.Println("Your score = ", player1.score)
			} else {
				fmt.Println("Your cards:")
				for _, zz := range player1.cards {
					switch zz.val {
					case 1:
						fmt.Println("Ace	", zz.suite)
					case 11:
						fmt.Println("Jack	", zz.suite)
					case 12:
						fmt.Println("Queen	", zz.suite)
					case 13:
						fmt.Println("King	", zz.suite)
					default:
						fmt.Println(zz.val, "	", zz.suite)
					}
				}

				fmt.Println("Dealer's cards:")
				for nn, zz := range dealer.cards {
					if nn != 0 {
						switch zz.val {
						case 1:
							fmt.Println("Ace	", zz.suite)
						case 11:
							fmt.Println("Jack	", zz.suite)
						case 12:
							fmt.Println("Queen	", zz.suite)
						case 13:
							fmt.Println("King	", zz.suite)
						default:
							fmt.Println(zz.val, "	", zz.suite)
						}
					}
				}

				if !player1.done {
					input = ""
					fmt.Println("Hit? Enter z to hit. Enter anything else to stand")
					fmt.Scanln(&input)
					if input == "z" {
						xx := Deck.draw()
						player1.cards = append(player1.cards, xx)
						player1.updateScore()
						fmt.Println("Your score = ", player1.score)

						if player1.score >= 21 {
							break
						}

					} else {
						player1.done = true
					}
				}
				if !dealer.done {
					if dealer.score <= 17 {
						yy := Deck.draw()
						dealer.cards = append(dealer.cards, yy)
						dealer.updateScore()

						if dealer.score >= 21 {
							break
						}
					} else {
						dealer.done = true
					}
				}
			}

		}
		player1.updateScore()
		dealer.updateScore()

		fmt.Println("Your score = ", player1.score)
		fmt.Println("Your cards:")
		for _, zz := range player1.cards {
			switch zz.val {
			case 1:
				fmt.Println("Ace	", zz.suite)
			case 11:
				fmt.Println("Jack	", zz.suite)
			case 12:
				fmt.Println("Queen	", zz.suite)
			case 13:
				fmt.Println("King	", zz.suite)
			default:
				fmt.Println(zz.val, "	", zz.suite)
			}
		}
		fmt.Println("Dealer's score = ", dealer.score)
		fmt.Println("Dealer's cards:")
		for _, zz := range dealer.cards {

			switch zz.val {
			case 1:
				fmt.Println("Ace	", zz.suite)
			case 11:
				fmt.Println("Jack	", zz.suite)
			case 12:
				fmt.Println("Queen	", zz.suite)
			case 13:
				fmt.Println("King	", zz.suite)
			default:
				fmt.Println(zz.val, "	", zz.suite)
			}

		}

		if player1.score > 21 {
			fmt.Println("You lose")
		} else if dealer.score > 21 {
			fmt.Println("You win")
		} else if player1.score > dealer.score {
			fmt.Println("You win")
		} else {
			fmt.Println("You lose")
		}

		fmt.Println("Stop playing? Enter x to stop playing.")
		fmt.Scanln(&input)
		if input == "x" {
			break
		}
	}

}
