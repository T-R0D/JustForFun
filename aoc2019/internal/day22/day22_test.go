package day22

import (
	"reflect"
	"testing"
)

var testInput = `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`

func TestParseInstructions(t *testing.T) {
	input := testInput
	result := parseInstructions(input)
	expected := []instruction{
		instruction{t: DEAL_NEW_STACK, val: 0},
		instruction{t: CUT, val: -2},
		instruction{t: DEAL_W_INCREMENT, val: 7},
		instruction{t: CUT, val: 8},
		instruction{t: CUT, val: -4},
		instruction{t: DEAL_W_INCREMENT, val: 7},
		instruction{t: CUT, val: 3},
		instruction{t: DEAL_W_INCREMENT, val: 9},
		instruction{t: DEAL_W_INCREMENT, val: 3},
		instruction{t: CUT, val: -1},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestDealIntoNewStack(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := dealIntoNewStack(input)
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestCut(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := cut(input, 3)
	expected := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestCutNegative(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := cut(input, -4)
	expected := []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestDealWithIncrement(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := dealWithIncrement(input, 3)
	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestShuffleCards(t *testing.T) {
	input := testInput
	instructions := parseInstructions(input)
	result := shuffleCards(instructions, 10)
	expected := []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestShuffleCardsSmall1(t *testing.T) {
	input := "deal with increment 7\ndeal into new stack\ndeal into new stack"
	instructions := parseInstructions(input)
	result := shuffleCards(instructions, 10)
	expected := []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestShuffleCardsSmall2(t *testing.T) {
	input := "cut 6\ndeal with increment 7\ndeal into new stack"
	instructions := parseInstructions(input)
	result := shuffleCards(instructions, 10)
	expected := []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestShuffleCardsSmall3(t *testing.T) {
	input := "deal with increment 7\ndeal with increment 9\ncut -2"
	instructions := parseInstructions(input)
	result := shuffleCards(instructions, 10)
	expected := []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}