package day16

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleData1 = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestScanTicketForErrorValue(t *testing.T) {
	testCases := []struct {
		notes              string
		ticketValues       []int
		expectedErrorValue int
	}{
		{
			notes:              sampleData1,
			ticketValues:       []int{7,3,47},
			expectedErrorValue: noErrorValue,
		},
		{
			notes:              sampleData1,
			ticketValues:       []int{40,4,50},
			expectedErrorValue: 4,
		},
		{
			notes:              sampleData1,
			ticketValues:       []int{55,2,20},
			expectedErrorValue: 55,
		},
		{
			notes:              sampleData1,
			ticketValues:       []int{38,6,12},
			expectedErrorValue: 12,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v is has an error value of %d", tc.ticketValues, tc.expectedErrorValue)
		t.Run(name, func(t *testing.T){
			// Arrange.
			collectedInformation, err := parseAllInformation(tc.notes)
			assert.NoError(t, err)
			ticket := trainTicket{values: tc.ticketValues}

			// Act.
			errorValue := scanTicketForErrorValue(collectedInformation.rules, ticket)
			
			// Assert.
			assert.Equal(t, tc.expectedErrorValue, errorValue)
		})
	}
}

const sampleData2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestRuleDecoderAnalyzeTicket(t *testing.T) {
	testCases := []struct {
		notes              string
		ticketValues       []int
		expectedErrorValue int
	}{}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T){
			fmt.Println(tc.notes)
		})
	}
}
