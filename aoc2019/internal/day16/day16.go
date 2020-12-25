package day16

import (
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	signal := inputToIntArray(input)
	cleanedSignal := cleanSignal(signal, 100)
	message := getDigitsFromSignal(cleanedSignal, 0, 8)
	return message, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	signal := inputToIntArray(input)
	signal = replicateSignal(signal, 10000)
	offset, err := strconv.Atoi(getDigitsFromSignal(signal, 0, 7))
	if err != nil {
		panic(err)
	} else if offset < (len(signal) / 2) {
		panic("The trick won't work...")
	}
	signal = signal[offset:]
	cleanedSignal := cleanSignalWithTrick(signal, 100)
	message := getDigitsFromSignal(cleanedSignal, 0, 8)
	return message, nil
}

func inputToIntArray(input string) []int {
	digits := strings.Split(input, "")
	out := make([]int, len(digits))
	for i, d := range digits {
		x, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		out[i] = x
	}
	return out
}

func cleanSignal(signal []int, iterations int) []int {
	cleanedSignal := make([]int, len(signal))
	for i := 0; i < iterations; i++ {
		cleanedSignal = computeFFT(signal)
		signal = cleanedSignal
	}
	return cleanedSignal
}

func computeFFT(in []int) []int {
	out := make([]int, len(in))
	repeatingPatterns := make(map[int]repeatingPattern)
	for i := range out {
		repeatingPatterns[i] = newRepeatingPattern(baseRepeatingPattern, i)
	}

	for i := range out {
		rp := repeatingPatterns[i]
		s := int64(0)
		for j, v := range in {
			multiplier := rp.get(j)
			s += int64(v * multiplier)
		}
		out[i] = int(math.Abs(float64(s % 10)))
	}
	return out
}

func getDigitsFromSignal(signal []int, offset, nDigits int) string {
	var b strings.Builder
	for i := 0; i < nDigits; i++ {
		b.WriteString(strconv.Itoa(signal[i]))
	}
	return b.String()
}

func replicateSignal(signal []int, n int) []int {
	replicatedSignal := make([]int, n*len(signal))
	for j := 0; j < len(replicatedSignal); j += len(signal) {
		for i := 0; i < len(signal); i++ {
			replicatedSignal[j+i] = signal[i]
		}
	}
	return replicatedSignal
}

func cleanSignalWithTrick(signal []int, iterations int) []int {
	cleanedSignal := make([]int, len(signal))
	for i := 0; i < iterations; i++ {
		cleanedSignal = computeFFTWithTrick(signal)
		signal = cleanedSignal
	}
	return cleanedSignal
}

func computeFFTWithTrick(signal []int) []int {
	out := make([]int, len(signal))

	out[len(out)-1] = signal[len(out)-1]
	for i := len(signal) - 2; i >= 0; i-- {
		out[i] = signal[i] + out[i+1]
	}

	for i := 0; i < len(out); i++ {
		out[i] = int(math.Abs(float64(out[i] % 10)))
	}

	return out
}

type repeatingPattern struct {
	base  []int
	len   int
	order int
}

var baseRepeatingPattern = []int{0, 1, 0, -1}

func newRepeatingPattern(basePattern []int, order int) repeatingPattern {
	return repeatingPattern{
		base:  basePattern,
		len:   (order + 1) * len(basePattern),
		order: order,
	}
}

func (r repeatingPattern) get(i int) int {
	i += 1
	i %= r.len
	return r.base[i/(r.order+1)]
}

func (r repeatingPattern) toIntArray() []int {
	rp := make([]int, 0, len(r.base)*(r.order+1))

	for i := 0; i < len(baseRepeatingPattern); i++ {
		for j := 0; j <= r.order; j++ {
			rp = append(rp, baseRepeatingPattern[i])
		}
	}

	rp = append(rp[1:], rp[0])

	return rp
}
