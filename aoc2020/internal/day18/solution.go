// Improvement: Add tests. This was easy enough to implement I didn't do them. :(
// Possible Improvement: Can we do better than that linked list thing in part 2?
// Possible Improvement: Use int64 instead of int so we aren't relying on 64 bit
//                       platforms.

package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	expressions := splitExpressions(input)

	resultSum := 0
	for i, expression := range expressions {
		result, err := evaluateSamePrecedenceExpression(expression)
		if err != nil {
			return "", errors.Wrapf(err, "processing expression %d", i)
		}
		resultSum += result
	}

	return strconv.Itoa(resultSum), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	expressions := splitExpressions(input)

	resultSum := 0
	for i, expression := range expressions {
		result, err := evaluateSwappedPrecedenceExpressoin(expression)
		if err != nil {
			return "", errors.Wrapf(err, "processing expression %d", i)
		}
		resultSum += result
	}

	return strconv.Itoa(resultSum), nil
}

func splitExpressions(input string) []string {
	return strings.Split(input, "\n")
}

const (
	operatorAdd           = '+'
	operatorMultiply      = '*'
	operatorNotDetermined = ' '
)

type valueAndOperator struct {
	value    int
	operator rune
}

func evaluateSamePrecedenceExpression(expression string) (int, error) {
	var expressionSoFar *valueAndOperator = nil
	outerExpressions := []*valueAndOperator{}

	for i, r := range expression {
		switch r {
		case '+':
			expressionSoFar.operator = operatorAdd
		case '*':
			expressionSoFar.operator = operatorMultiply
		case '(':
			outerExpressions = append(outerExpressions, expressionSoFar)
			expressionSoFar = nil
		case ')':
			outerExpression := outerExpressions[len(outerExpressions)-1]
			outerExpressions = outerExpressions[:len(outerExpressions)-1]

			if outerExpression == nil {
				outerExpression = &valueAndOperator{
					value:    expressionSoFar.value,
					operator: operatorNotDetermined,
				}
			} else {
				switch outerExpression.operator {
				case operatorAdd:
					outerExpression.value += expressionSoFar.value
				case operatorMultiply:
					outerExpression.value *= expressionSoFar.value
				default:
					return 0, errors.Errorf("unable to use %c as an operator (at position %d)", expressionSoFar.operator, i)
				}
				outerExpression.operator = operatorNotDetermined
			}

			expressionSoFar = outerExpression

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value, err := strconv.Atoi(fmt.Sprintf("%c", r))
			if err != nil {
				return 0, errors.Wrapf(err, "processing item at index %d", i)
			}

			if expressionSoFar == nil {
				expressionSoFar = &valueAndOperator{
					value:    value,
					operator: operatorNotDetermined,
				}
			} else {
				switch expressionSoFar.operator {
				case operatorAdd:
					expressionSoFar.value += value
				case operatorMultiply:
					expressionSoFar.value *= value
				default:
					return 0, errors.Errorf("unable to use %c as an operator (position %d)", expressionSoFar.operator, i)
				}
				expressionSoFar.operator = operatorNotDetermined
			}
		case ' ':
			continue
		default:
			return 0, errors.Errorf("unrecognized symbol %c at position %d", r, i)
		}
	}

	return expressionSoFar.value, nil
}

type valueNode struct {
	value    int
	operator *operatorNode
}

type operatorNode struct {
	previousValue *valueNode
	operator      rune
	nextValue     *valueNode
}

type unfinishedExpression struct {
	expressionStart *valueNode
	lastOperator    *operatorNode
}

func evaluateSwappedPrecedenceExpressoin(expression string) (int, error) {
	var expressionStart *valueNode = nil
	var lastOperator *operatorNode = nil
	outerExpressions := []unfinishedExpression{}

	for i, r := range expression {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value, err := strconv.Atoi(fmt.Sprintf("%c", r))
			if err != nil {
				return 0, errors.Wrapf(err, "processing item at index %d", i)
			}

			if expressionStart == nil {
				expressionStart = &valueNode{
					value:    value,
					operator: nil,
				}
			} else if lastOperator != nil {
				lastOperator.nextValue = &valueNode{
					value:    value,
					operator: nil,
				}
			} else {
				return 0, errors.Errorf("lastOperator isn't pointing to anything (at %d)", i)
			}
		case '*':
			if lastOperator == nil {
				lastOperator = &operatorNode{
					previousValue: expressionStart,
					operator:      operatorMultiply,
					nextValue:     nil,
				}
				expressionStart.operator = lastOperator
			} else {
				newOperatorNode := &operatorNode{
					previousValue: lastOperator.nextValue,
					operator:      operatorMultiply,
					nextValue:     nil,
				}
				lastOperator.nextValue.operator = newOperatorNode
				lastOperator = newOperatorNode
			}
		case '+':
			if lastOperator == nil {
				lastOperator = &operatorNode{
					previousValue: expressionStart,
					operator:      operatorAdd,
					nextValue:     nil,
				}
				expressionStart.operator = lastOperator
			} else {
				newOperatorNode := &operatorNode{
					previousValue: lastOperator.nextValue,
					operator:      operatorAdd,
					nextValue:     nil,
				}
				lastOperator.nextValue.operator = newOperatorNode
				lastOperator = newOperatorNode
			}
		case ' ':
			continue
		case '(':
			outerExpressions = append(outerExpressions, unfinishedExpression{
				expressionStart: expressionStart,
				lastOperator:    lastOperator,
			})
			expressionStart = nil
			lastOperator = nil
		case ')':
			value := evaluateFlatSwappedPrecedenceExpression(expressionStart)
			outerExpression := outerExpressions[len(outerExpressions)-1]
			outerExpressions = outerExpressions[:len(outerExpressions)-1]
			expressionStart, lastOperator = outerExpression.expressionStart, outerExpression.lastOperator

			if expressionStart == nil {
				expressionStart = &valueNode{
					value:    value,
					operator: nil,
				}
			} else if lastOperator != nil {
				lastOperator.nextValue = &valueNode{
					value:    value,
					operator: nil,
				}
			} else {
				return 0, errors.Errorf("lastOperator isn't pointing to anything (at %d)", i)
			}

		default:
			return 0, errors.Errorf("unrecognized symbol %c at position %d", r, i)
		}
	}

	return evaluateFlatSwappedPrecedenceExpression(expressionStart), nil
}

func evaluateFlatSwappedPrecedenceExpression(expressionStart *valueNode) int {
	if expressionStart == nil {
		return 0
	}

	// Addition pass.
	operator := expressionStart.operator
	for operator != nil {
		if operator.operator == operatorAdd {
			previousValue := operator.previousValue
			previousValue.value += operator.nextValue.value
			previousValue.operator = operator.nextValue.operator
			operator = operator.nextValue.operator
			if operator != nil {
				operator.previousValue = previousValue
			}
		} else {
			operator = operator.nextValue.operator
		}
	}

	// Multiplication pass.
	operator = expressionStart.operator
	for operator != nil {
		expressionStart.value *= operator.nextValue.value
		operator = operator.nextValue.operator
	}

	return expressionStart.value
}

func (v *valueNode) String() string {
	currentValue := v
	builder := strings.Builder{}

	for currentValue != nil {
		builder.WriteString(strconv.Itoa(currentValue.value))
		if currentValue.operator != nil {
			builder.WriteRune(currentValue.operator.operator)
			currentValue = currentValue.operator.nextValue
		} else {
			currentValue = nil
		}
	}

	return builder.String()
}
