package day23

import (
	"fmt"
	"slices"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	graphInfo := parseConnectivityGraph(input)

	size3Cliques := findCliquesOfSize3(graphInfo.Graph)

	cliquesThatHaveATNamedMember :=
		findCliquesWithAtLeastOneNameWithPrefix(size3Cliques, graphInfo.NameLookup, "t")

	return fmt.Sprintf("%d", len(cliquesThatHaveATNamedMember)), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	graphInfo := parseConnectivityGraph(input)

	largestCliqueName := getLargestCliqueByName(graphInfo.Graph, graphInfo.NameLookup)

	return largestCliqueName, nil
}

type connectivityGraphInfo struct {
	IndexLookup map[string]int
	NameLookup  map[int]string
	Graph       [][]int
}

func parseConnectivityGraph(input string) connectivityGraphInfo {
	lines := strings.Split(input, "\n")
	connections := make([][2]string, 0, len(lines))
	for _, line := range lines {
		members := strings.Split(line, "-")
		connections = append(connections, [2]string(members))
	}

	names := []string{}
	nameSet := map[string]struct{}{}
	for _, pair := range connections {
		for _, name := range pair {
			if _, ok := nameSet[name]; !ok {
				names = append(names, name)
				nameSet[name] = struct{}{}
			}
		}
	}

	nameLookup := map[int]string{}
	indexLookup := map[string]int{}
	for i, name := range names {
		nameLookup[i] = name
		indexLookup[name] = i
	}

	graph := make([][]int, 0, len(names))
	for range len(names) {
		graph = append(graph, make([]int, len(names)))
	}

	for _, pair := range connections {
		i, j := indexLookup[pair[0]], indexLookup[pair[1]]
		graph[i][j] = 1
		graph[j][i] = 1
	}

	return connectivityGraphInfo{
		IndexLookup: indexLookup,
		NameLookup:  nameLookup,
		Graph:       graph,
	}
}

func getLargestCliqueByName(graph [][]int, nameLookup map[int]string) string {
	maximumClique := findMaximumCliqueNaive(graph)

	names := make([]string, 0, len(maximumClique))
	for _, id := range maximumClique {
		names = append(names, nameLookup[id])
	}
	slices.Sort(names)

	return strings.Join(names, ",")
}

func findMaximumCliqueNaive(graph [][]int) []int {
	maximumDegree := 0
	for i, row := range graph {
		degree := 0
		for j, weight := range row {
			if i == j {
				continue
			}

			if weight == 1 {
				degree += 1
			}
		}

		if degree > maximumDegree {
			maximumDegree = degree
		}
	}

	maximumClique := []int{}
	for n := maximumDegree; n >= 0; n -= 1 {
		foundCliques := findCliquesOfSizeNIterative(graph, n)
		if len(foundCliques) > 0 {
			maximumClique = foundCliques[0]
			break
		}
	}

	return maximumClique
}

func findCliquesOfSize3(graph [][]int) [][]int {
	cliques := [][]int{}
	for i := range len(graph) {
		for j := i + 1; j < len(graph); j += 1 {
			for k := j + 1; k < len(graph); k += 1 {
				if graph[i][j] == 1 && graph[j][k] == 1 && graph[k][i] == 1 {
					cliques = append(cliques, []int{i, j, k})
				}
			}
		}
	}

	return cliques
}

func findCliquesOfSizeNIterative(graph [][]int, n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	type stackFrame struct {
		I                int
		IncludedVertices []int
	}

	callStack := queue.NewLifo[stackFrame]()
	for i := range len(graph) - n {
		callStack.Push(stackFrame{I: i, IncludedVertices: []int{i}})
	}

	foundCliques := [][]int{}
	for callStack.Len() > 0 {
		frame, ok := callStack.Pop()
		if !ok {
			break
		}

		if len(frame.IncludedVertices) == n {
			foundCliques = append(foundCliques, frame.IncludedVertices)
			continue
		}

		if frame.I >= len(graph) {
			continue
		}

	NEIGHBOR_SEARCH:
		for j := frame.I + 1; j < len(graph)-(n-len(frame.IncludedVertices)); j += 1 {
			for _, k := range frame.IncludedVertices {
				if graph[j][k] != 1 {
					continue NEIGHBOR_SEARCH
				}
			}

			newIncludedVertices := append(make([]int, 0, len(frame.IncludedVertices)+1), frame.IncludedVertices...)
			newIncludedVertices = append(newIncludedVertices, j)

			callStack.Push(stackFrame{I: j, IncludedVertices: newIncludedVertices})
		}

	}

	return foundCliques
}

func findCliquesWithAtLeastOneNameWithPrefix(
	cliques [][]int,
	nameLookup map[int]string,
	prefix string) [][]int {

	matchingCliques := make([][]int, 0, len(cliques))
	for _, clique := range cliques {
		for _, index := range clique {
			if name, ok := nameLookup[index]; ok && strings.HasPrefix(name, prefix) {
				matchingCliques = append(matchingCliques, clique)
				break
			}
		}
	}

	return matchingCliques
}

/**
 * Keeping these Bron-Kerbosch implementations around because I liked them.
 * After all, AoC is about learning something new, and this was a new thing.
 * Also, I made it iterative, which did provide some speedup.
 * It was weird though, this is an "official" algorithm, but it performs about
 * 3 times as slow as my naive find-all-cliques-of-size-N approach. I think
 * It's because that approach starts (somewhat greedily) with the highest
 * degree of any vertex and works down, while Bron-Kerbosh will find all
 * maximal cliques of any size. I do wonder if we can tune Bron-Kerbosch to
 * get more speedup using the maximum degree as a hint (Wikipedia does
 * mention a degeneracy ordering of the graph, after all).
 * I also find it curious that my implementation is "slow" compared to others.
 * Online I'm seeing reports of, like, 7ms, while my times are around
 * 42ms/128ms (naive/Bron-Kerbosh). I wonder if I'm implementing things in a
 * dumb way or doing too much copying or something...
 */
 
/*
type vertexSet map[int]struct{}

func findMaximumCliqueBronKerbosch(graph [][]int) []int {
	foundCliques := findMaximalCliquesBronKerboschWithPivotIterative(graph)

	maximumCliqueIndex := 0
	for i, clique := range foundCliques {
		if len(clique) > len(foundCliques[maximumCliqueIndex]) {
			maximumCliqueIndex = i
		}
	}

	return foundCliques[maximumCliqueIndex]
}

func findMaximalCliquesBronKerboschWithPivotIterative(graph [][]int) [][]int {
	foundCliques := [][]int{}

	type stackFrame struct {
		R vertexSet
		P vertexSet
		X vertexSet
	}

	callStack := queue.NewLifo[stackFrame]()
	r := vertexSet{}
	p := vertexSet{}
	for i := range len(graph) {
		p[i] = struct{}{}
	}
	x := vertexSet{}
	callStack.Push(stackFrame{R: r, P: p, X: x})

	for callStack.Len() > 0 {
		frame, ok := callStack.Pop()
		if !ok {
			break
		}

		if len(frame.P) == 0 && len(frame.X) == 0 {
			clique := make([]int, 0, len(frame.R))
			for x := range frame.R {
				clique = append(clique, x)
			}
			foundCliques = append(foundCliques, clique)
			continue
		}

		pUnionX := findUnion(frame.P, frame.X)
		pivot := choosePivot(pUnionX)
		neighborsOfPivot := vertexSet{}
		for x, weight := range graph[pivot] {
			if x != pivot && weight == 1 {
				neighborsOfPivot[x] = struct{}{}
			}
		}
		pDiffNeighborsOfPivot := findDifference(frame.P, neighborsOfPivot)

		for v := range pDiffNeighborsOfPivot {
			newR := maps.Clone(frame.R)
			newR[v] = struct{}{}

			newP := vertexSet{}
			for x := range frame.P {
				for neighbor, weight := range graph[v] {
					if neighbor == x && weight == 1 {
						newP[neighbor] = struct{}{}
					}
				}
			}

			newX := vertexSet{}
			for x := range frame.X {
				for neighbor, weight := range graph[v] {
					if neighbor == x && weight == 1 {
						newP[neighbor] = struct{}{}
					}
				}
			}

			callStack.Push(stackFrame{R: newR, P: newP, X: newX})

			delete(frame.P, v)
			frame.X[v] = struct{}{}
		}
	}

	return foundCliques
}

func findUnion(a vertexSet, b vertexSet) vertexSet {
	union := maps.Clone(a)
	for x := range b {
		union[x] = struct{}{}
	}
	return union
}

func choosePivot(a vertexSet) int {
	for x := range a {
		return x
	}
	return -1
}

func findDifference(a vertexSet, b vertexSet) vertexSet {
	difference := vertexSet{}
	for x := range a {
		if _, ok := b[x]; !ok {
			difference[x] = struct{}{}
		}
	}
	return difference
}
*/
