// Possible Improvement: Can we speed this up? 13 seconds is still slow...
//                       Maybe a ring buffer instead of a linked one?

package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/queue"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	hands, err := readStartingHands(input)
	if err != nil {
		return "", err
	}

	resultingHands, err := playCombat(hands)
	if err != nil {
		return "", err
	}

	score := uint64(0)
	for _, hand := range resultingHands {
		score += scoreHand(hand)
	}

	return fmt.Sprint(score), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	hands, err := readStartingHands(input)
	if err != nil {
		return "", err
	}

	seenGames := map[string]int{}
	resultingHands, err := playRecursiveCombat(hands, seenGames)
	if err != nil {
		return "", err
	}

	score := uint64(0)
	for _, hand := range resultingHands {
		score += scoreHand(hand)
	}

	return fmt.Sprint(score), nil
}

func readStartingHands(input string) ([][]uint64, error) {
	handSections := strings.Split(input, "\n\n")

	hands := make([][]uint64, len(handSections))

	for i, handSection := range handSections {
		valueStrs := strings.Split(handSection, "\n")[1:]
		hands[i] = make([]uint64, len(valueStrs))
		for j, valueStr := range valueStrs {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, errors.Errorf("section %d, value %d", i, j)
			}
			hands[i][j] = uint64(value)
		}
	}

	return hands, nil
}

func playCombat(startingHands [][]uint64) ([][]uint64, error) {
	player1Hand := queue.NewLinkQueue()
	player2Hand := queue.NewLinkQueue()
	for i := 0; i < len(startingHands[0]); i++ {
		player1Hand.AppendRight(startingHands[0][i])
		player2Hand.AppendRight(startingHands[1][i])
	}

	for player1Hand.Len() > 0 && player2Hand.Len() > 0 {
		player1CardObj, err := player1Hand.PopLeft()
		if err != nil {
			return nil, err
		}
		player1Card, ok := player1CardObj.(uint64)
		if !ok {
			return nil, errors.Errorf("%v was %T, not uint64!", player1CardObj, player1CardObj)
		}
		player2CardObj, err := player2Hand.PopLeft()
		if err != nil {
			return nil, err
		}
		player2Card, ok := player2CardObj.(uint64)
		if !ok {
			return nil, errors.Errorf("%v was %T, not uint64!", player2CardObj, player2CardObj)
		}

		if player1Card > player2Card {
			player1Hand.AppendRight(player1Card)
			player1Hand.AppendRight(player2Card)
		} else {
			player2Hand.AppendRight(player2Card)
			player2Hand.AppendRight(player1Card)
		}
	}

	resultingHands := make([][]uint64, 2)
	resultingHands[0] = make([]uint64, 0, player1Hand.Len())
	resultingHands[1] = make([]uint64, 0, player2Hand.Len())

	for player1Hand.Len() > 0 {
		cardObj, err := player1Hand.PopLeft()
		if err != nil {
			return nil, err
		}
		card, ok := cardObj.(uint64)
		if !ok {
			return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
		}
		resultingHands[0] = append(resultingHands[0], card)
	}

	for player2Hand.Len() > 0 {
		cardObj, err := player2Hand.PopLeft()
		if err != nil {
			return nil, err
		}
		card, ok := cardObj.(uint64)
		if !ok {
			return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
		}
		resultingHands[1] = append(resultingHands[1], card)
	}

	return resultingHands, nil
}

func playRecursiveCombat(startingHands [][]uint64, seenGames map[string]int) ([][]uint64, error) {
	hands := make([]queue.Queue, 2)
	for i := 0; i < 2; i++ {
		hands[i] = queue.NewLinkQueue()

		for _, card := range startingHands[i] {
			hands[i].AppendRight(card)
		}
	}

	gameStr := stringifyHands(hands)
	if winner, ok := seenGames[gameStr]; ok {
		if winner == 0 {
			for hands[1].Len() > 0 {
				cardObj, err := hands[1].PopLeft()
				if err != nil {
					return nil, err
				}
				hands[0].AppendRight(cardObj)
			}
		} else {
			for hands[0].Len() > 0 {
				cardObj, err := hands[0].PopLeft()
				if err != nil {
					return nil, err
				}
				hands[1].AppendRight(cardObj)
			}
		}

		resultingHands := make([][]uint64, 2)
		for i := 0; i < 2; i++ {
			handLen := hands[i].Len()
			hand := make([]uint64, handLen)
			for j := 0; j < handLen; j++ {
				cardObj, err := hands[i].PopLeft()
				if err != nil {
					return nil, err
				}
				card, ok := cardObj.(uint64)
				if !ok {
					return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
				}
				hand[j] = card
			}
			resultingHands[i] = hand
		}

		return resultingHands, nil
	}

	seenHands := map[string]struct{}{}

	for hands[0].Len() > 0 && hands[1].Len() > 0 {
		handsStr := stringifyHands(hands)
		if _, ok := seenHands[handsStr]; ok {
			for hands[1].Len() > 0 {
				cardObj, err := hands[1].PopLeft()
				if err != nil {
					return nil, err
				}
				card, ok := cardObj.(uint64)
				if !ok {
					return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
				}
				hands[0].AppendRight(card)
			}
			break
		}
		seenHands[handsStr] = struct{}{}

		cards := make([]uint64, 2)
		for i, hand := range hands {
			cardObj, err := hand.PopLeft()
			if err != nil {
				return nil, err
			}
			card, ok := cardObj.(uint64)
			if !ok {
				return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
			}
			cards[i] = card
		}

		winner := 0
		if cards[0] <= uint64(hands[0].Len()) && cards[1] <= uint64(hands[1].Len()) {
			newStartingHands := make([][]uint64, 2)
			for i := 0; i < 2; i++ {
				hand := make([]uint64, cards[i])
				for j := 0; j < hands[i].Len(); j++ {
					cardObj, err := hands[i].PopLeft()
					if err != nil {
						return nil, err
					}
					card, ok := cardObj.(uint64)
					if !ok {
						return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
					}

					if uint64(j) < cards[i] {
						hand[j] = card
					}

					hands[i].AppendRight(card)
				}
				newStartingHands[i] = hand
			}

			resultingHands, err := playRecursiveCombat(newStartingHands, seenGames)
			if err != nil {
				return nil, err
			}

			if len(resultingHands[0]) > len(resultingHands[1]) {
				winner = 0
			} else {
				winner = 1
			}

		} else {
			if cards[0] > cards[1] {
				winner = 0
			} else {
				winner = 1
			}
		}

		if winner == 0 {
			hands[0].AppendRight(cards[0])
			hands[0].AppendRight(cards[1])
		} else {
			hands[1].AppendRight(cards[1])
			hands[1].AppendRight(cards[0])
		}
	}

	resultingHands := make([][]uint64, 2)
	for i := 0; i < 2; i++ {
		handLen := hands[i].Len()
		hand := make([]uint64, handLen)
		for j := 0; j < handLen; j++ {
			cardObj, err := hands[i].PopLeft()
			if err != nil {
				return nil, err
			}
			card, ok := cardObj.(uint64)
			if !ok {
				return nil, errors.Errorf("%v was %T, not uint64!", cardObj, cardObj)
			}
			hand[j] = card
		}
		resultingHands[i] = hand
	}

	if len(resultingHands[0]) > len(resultingHands[1]) {
		seenGames[gameStr] = 0
	} else {
		seenGames[gameStr] = 1
	}

	return resultingHands, nil
}

func scoreHand(hand []uint64) uint64 {
	score := uint64(0)
	for i := 0; i < len(hand); i++ {
		score += (uint64(i) + uint64(1)) * hand[len(hand)-i-1]
	}
	return score
}

func stringifyHands(hands []queue.Queue) string {
	builder := strings.Builder{}
	for _, hand := range hands {
		builder.WriteString(hand.String())
	}
	return builder.String()
}
