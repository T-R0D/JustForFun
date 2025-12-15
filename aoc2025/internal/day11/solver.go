package day11

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/counter"
	"github.com/T-R0D/aoc2025/v2/internal/queue"
	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	mappings := parseDataFlowGraph(input)

	nPaths, err := countPathsThroughSystem(mappings, youServer, outServer)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(nPaths), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	mappings := parseDataFlowGraph(input)

	nPathsAdjacencyMatrix, labelToIndex := countPathsByMatrixMultiplication(mappings)

	svr, ok := labelToIndex[svrServer]
	if !ok {
		return "", fmt.Errorf("'%s' was not in the graph", svrServer)
	}
	dac, ok := labelToIndex[dacServer]
	if !ok {
		return "", fmt.Errorf("'%s' was not in the graph", dacServer)
	}
	fft, ok := labelToIndex[fftServer]
	if !ok {
		return "", fmt.Errorf("'%s' was not in the graph", fftServer)
	}
	out, ok := labelToIndex[outServer]
	if !ok {
		return "", fmt.Errorf("'%s' was not in the graph", outServer)
	}

	nPathsSvrToDac := nPathsAdjacencyMatrix[svr][dac]
	nPathsSvrToFft := nPathsAdjacencyMatrix[svr][fft] 
	nPathsDacToFft := nPathsAdjacencyMatrix[dac][fft] 
	nPathsFftToDac := nPathsAdjacencyMatrix[fft][dac] 
	nPathsDacToOut := nPathsAdjacencyMatrix[dac][out] 
	nPathsFftToOut := nPathsAdjacencyMatrix[fft][out]

	nPaths := 0
	if nPathsDacToFft > 0 {
		nPaths = nPathsSvrToDac * nPathsDacToFft * nPathsFftToOut
	} else {
		nPaths = nPathsSvrToFft * nPathsFftToDac * nPathsDacToOut
	}

	return strconv.Itoa(nPaths), nil
}

type dataFlowMapping struct {
	src  string
	dsts []string
}

func parseDataFlowGraph(input string) []dataFlowMapping {
	lines := strings.Split(input, "\n")
	mappings := make([]dataFlowMapping, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		dsts := strings.Split(parts[1], " ")
		mappings = append(mappings, dataFlowMapping{
			src:  parts[0],
			dsts: dsts,
		})
	}
	return mappings
}

// This method is faster for small graphs/short paths. Assumes no cycles.
func countPathsThroughSystem(labeledMappings []dataFlowMapping, src string, dst string) (int, error) {
	mappings := listOfListsToMapOfLists(labeledMappings)
	pathThroughCounts := counter.New[string]()

	stack := queue.NewLifo[string]()
	stack.Push(src)

	for stack.Len() > 0 {
		current, ok := stack.Pop()
		if !ok {
			return 0, fmt.Errorf("somehow, the stack had nothing in it")
		}

		pathThroughCounts.Increment(current, 1)

		if nexts, ok := mappings[current]; ok {
			for _, next := range nexts {
				stack.Push(next)
			}
		}
	}

	nPaths, exists := pathThroughCounts.Get(dst)
	if !exists {
		nPaths = 0
	}

	return nPaths, nil
}

// This method is faster for large graphs. Assumes no cycles.
func countPathsByMatrixMultiplication(listOfLists []dataFlowMapping) ([][]int, map[string]int) {
	labelToIndex := map[string]int{}
	labelSet := set.New[string]()
	for i, mapping := range listOfLists {
		labelSet.Add(mapping.src)
		labelToIndex[mapping.src] = i

		for _, dst:=range mapping.dsts {
			labelSet.Add(dst)
		}
	}

	m := len(labelToIndex)
	for label := range labelSet.All() {
		if _, exists := labelToIndex[label]; !exists {
			labelToIndex[label] = m
			m += 1
		}
	}

	n := labelSet.Len()
	adjacencyMatrix := make([][]int, 0, n)
	for range n {
		adjacencyMatrix = append(adjacencyMatrix, slices.Repeat([]int{0}, n))
	}

	for i, mapping := range listOfLists {
		for _, dst := range mapping.dsts {
			if j, ok := labelToIndex[dst]; ok {
				adjacencyMatrix[i][j] = 1
			}
		}
	}

	for k := range n {
		for j := range n {
			for i := range n {
				adjacencyMatrix[i][j] += adjacencyMatrix[i][k] * adjacencyMatrix[k][j]
			}
		}
	}

	return adjacencyMatrix, labelToIndex
}

func listOfListsToMapOfLists(listOfLists []dataFlowMapping) map[string][]string {
	mapOfLists := map[string][]string{}
	for _, mapping := range listOfLists {
		dsts := make([]string, len(mapping.dsts))
		copy(dsts, mapping.dsts)
		mapOfLists[mapping.src] = dsts
	}
	return mapOfLists
}

const (
	dacServer = "dac"
	fftServer = "fft"
	outServer = "out"
	svrServer = "svr"
	youServer = "you"
)
