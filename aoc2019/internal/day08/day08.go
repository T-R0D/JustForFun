package day08

import (
	"fmt"
	"strconv"
)

const (
	kIMAGE_WIDTH_PX  = 25
	kIMAGE_HEIGHT_PX = 6
	kLAYER_PX        = kIMAGE_WIDTH_PX * kIMAGE_HEIGHT_PX

	COLOR_BLACK       = 0
	COLOR_WHITE       = 1
	COLOR_TRANSPARENT = 2
)

type Solver struct{}

type layer []int

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	layers, err := inputStreamToLayers(input, kIMAGE_WIDTH_PX, kIMAGE_HEIGHT_PX)
	if err != nil {
		return nil, err
	}

	layerWithFewestZeros := findLayerWithFewestZeros(layers, kLAYER_PX)

	ones := countDigits(layers[layerWithFewestZeros], 1)
	twos := countDigits(layers[layerWithFewestZeros], 2)

	// 2352 is too high.
	return ones * twos, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	layers, err := inputStreamToLayers(input, kIMAGE_WIDTH_PX, kIMAGE_HEIGHT_PX)
	if err != nil {
		return nil, err
	}

	image := decodeImage(layers, kLAYER_PX)

	return image.ToString(kIMAGE_WIDTH_PX), nil
}

func inputStreamToLayers(input string, w, h int) ([]layer, error) {
	sLen := len(input)
	layerLen := h * w
	layers := make([]layer, sLen/layerLen)
	for i := range layers {
		layers[i] = make(layer, layerLen)
	}

	for i, b := range input {
		pxVal, err := strconv.Atoi(string(b))
		if err != nil {
			return nil, fmt.Errorf("Error decoding pixel value from stream - value: %v byte: %v", b, i)
		}
		layer := i / layerLen
		layers[layer][i%layerLen] = pxVal
	}
	return layers, nil
}

func countDigits(l layer, d int) int {
	digits := 0
	for _, v := range l {
		if v == d {
			digits++
		}
	}
	return digits
}

func findLayerWithFewestZeros(layers []layer, layerLen int) int {
	layerWithFewestZeros := 0
	zeros := layerLen + 1
	for i, layer := range layers {
		zs := countDigits(layer, 0)
		if zs < zeros {
			zeros = zs
			layerWithFewestZeros = i
		}
	}
	return layerWithFewestZeros
}

func decodeImage(layers []layer, layerLen int) layer {
	r := make(layer, layerLen)
	for i := 0; i < layerLen; i++ {
		for _, layer := range layers {
			r[i] = layer[i]
			if layer[i] != COLOR_TRANSPARENT {
				break
			}
		}
	}
	return r
}

func (l layer) ToString(width int) string {
	r := ""
	for i, v := range l {
		if i%width == 0 {
			r += "\n"
		}

		switch v {
		case COLOR_TRANSPARENT:
			r += " "
		case COLOR_BLACK:
			r += "X"
		case COLOR_WHITE:
			r += "."
		}
	}
	r += "\n"
	return r
}
