package day22

import "strings"

import "fmt"

import "math"

import "math/big"

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	instructions := parseInstructions(input)
	shuffledDeck := shuffleCards(instructions, 10007)
	var posOf2019 int
	for i, v := range shuffledDeck {
		if v == 2019 {
			posOf2019 = i
			break
		}
	}
	return posOf2019, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	instructions := parseInstructions(input)
	result := shuffleCardsWithModularArithmetic(instructions, 119315717514047, 101741582076661)
	return result, nil
}

const (
	DEAL_NEW_STACK   = "DEAL_NEW_STACK"
	CUT              = "CUT"
	DEAL_W_INCREMENT = "DEAL_W_INCREMENT"
)

type instruction struct {
	t   string
	val int
}

func parseInstructions(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		if strings.Contains(line, "deal into new stack") {
			instructions[i] = instruction{
				t:   DEAL_NEW_STACK,
				val: 0,
			}
		} else if strings.Contains(line, "cut") {
			var val int
			_, err := fmt.Sscanf(line, "cut %d", &val)
			if err != nil {
				panic("Value could not be extracted from cut instruction!")
			}
			instructions[i] = instruction{
				t:   CUT,
				val: val,
			}
		} else if strings.Contains(line, "deal with increment") {
			var val int
			_, err := fmt.Sscanf(line, "deal with increment %d", &val)
			if err != nil {
				panic("Value could not be extracted from deal with increment instruction!")
			}
			instructions[i] = instruction{
				t:   DEAL_W_INCREMENT,
				val: val,
			}
		} else {
			panic("Line could not be parsed!")
		}
	}
	return instructions
}

func shuffleCards(instructions []instruction, deckSize int) []int {
	deck := make([]int, deckSize)
	for i := range deck {
		deck[i] = i
	}

	for _, instruction := range instructions {
		switch instruction.t {
		case DEAL_NEW_STACK:
			deck = dealIntoNewStack(deck)
		case CUT:
			deck = cut(deck, instruction.val)
		case DEAL_W_INCREMENT:
			deck = dealWithIncrement(deck, instruction.val)
		}
	}

	return deck
}

func dealIntoNewStack(deck []int) []int {
	for i, j := 0, len(deck)-1; i < j; i, j = i+1, j-1 {
		temp := deck[i]
		deck[i] = deck[j]
		deck[j] = temp
	}
	return deck
}

func cut(deck []int, n int) []int {
	if n >= 0 {
		deck = append(deck[n:], deck[:n]...)
	} else {
		n = int(math.Abs(float64(n)))
		deck = append(deck[len(deck)-n:], deck[:len(deck)-n]...)
	}
	return deck
}

func dealWithIncrement(deck []int, n int) []int {
	newDeck := make([]int, len(deck))
	for i, v := range deck {
		newDeck[(i*n)%len(deck)] = v
	}
	return newDeck
}

func shuffleCardsWithModularArithmetic(instructions []instruction, deckSize, iterations int) *big.Int {
	// Shamefully copied to get past this problem. I don't understand the modular arithmetic.
	n, iters := big.NewInt(int64(deckSize)), big.NewInt(int64(iterations))
	offset, increment := big.NewInt(0), big.NewInt(1)
	for _, instruction := range instructions {
		switch instruction.t {
		case DEAL_NEW_STACK:
			increment.Mul(increment, big.NewInt(-1))
			offset.Add(offset, increment)
		case CUT:
			offset.Add(offset, big.NewInt(0).Mul(big.NewInt(int64(instruction.val)), increment))
		case DEAL_W_INCREMENT:
			increment.Mul(increment, big.NewInt(0).Exp(big.NewInt(int64(instruction.val)), big.NewInt(0).Sub(n, big.NewInt(2)), n))
		}
	}

	finalIncrement := big.NewInt(0).Exp(increment, iters, n)
	finalOffset := big.NewInt(0).Exp(increment, iters, n)
	finalOffset.Sub(big.NewInt(1), finalOffset)
	inverseMod := big.NewInt(0).Exp(big.NewInt(0).Sub(big.NewInt(1), increment), big.NewInt(0).Sub(n, big.NewInt(2)), n)
	finalOffset.Mul(finalOffset, inverseMod)
	finalOffset.Mul(finalOffset, offset)

	answer := big.NewInt(0).Mul(big.NewInt(2020), finalIncrement)
	answer.Add(answer, finalOffset)
	answer.Mod(answer, n)

	return answer
}
