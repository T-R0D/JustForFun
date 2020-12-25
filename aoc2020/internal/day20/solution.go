// TODO: Add tests for sea monster tagging.
// Possible Improvement: See if we can make a more generic imageTile struct
//                       that could replace mergedImageContainer
// Possible Improvement: Instead of making copies of tiles for the backtracking
//                       algorithm, create an order of transformations and
//                       perform them on the same object, reverting each time.

package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	tiles, err := readTiles(input)
	if err != nil {
		return "", err
	}

	fullImage, err := constructFullImage(tiles)
	if err != nil {
		return "", err
	}

	checkSum := fullImage.TopLeftTileID() * fullImage.TopRightTileID() *
		fullImage.BottomLeftTileID() * fullImage.BottomRightTileID()

	return fmt.Sprintf("%d", checkSum), nil
}

// Part2 solves part 2 of the day's problem.
// !4868
func (s *Solver) Part2(input string) (string, error) {
	tiles, err := readTiles(input)
	if err != nil {
		return "", err
	}

	fullImage, err := constructFullImage(tiles)
	if err != nil {
		return "", err
	}

	mergedImage := fullImage.MergeTiles()

	mergedImage.ManipulateUntilSeaMonstersTagged()

	waterRoughness := mergedImage.DetermineWaterRoughness()

	return fmt.Sprintf("%d", waterRoughness), nil
}

const (
	tileDim = 10
)

func readTiles(input string) ([]*imageTile, error) {
	idRegex, err := regexp.Compile(`^Tile (\d+):$`)
	if err != nil {
		return nil, err
	}

	tilesWithID := strings.Split(input, "\n\n")

	tiles := make([]*imageTile, len(tilesWithID))
	for k, tileWithID := range tilesWithID {
		idLineAndImage := strings.SplitN(tileWithID, "\n", 2)

		id := uint64(0)
		if matches := idRegex.FindStringSubmatch(idLineAndImage[0]); len(matches) > 0 {
			idValue, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, errors.Wrapf(err, "tile %d didn't have a valid ID", k)
			}
			id = uint64(idValue)
		} else {
			return nil, errors.Errorf("unable to match id line of tile %d", k)
		}

		tile := [tileDim][tileDim]rune{}
		i, j := 0, 0
		for _, r := range idLineAndImage[1] {

			switch r {
			case '.', '#':
				tile[i][j] = r
				j++
			case '\n':
				i, j = i+1, 0
			default:
				return nil, errors.Errorf("unrecognized rune %c when processing tile %d", r, k)
			}
		}

		// TODO: verify image was filled

		tiles[k] = &imageTile{
			id:     id,
			pixels: tile,
		}
	}

	return tiles, nil
}

type fullImageContainer struct {
	dim  int
	data [][]*imageTile
}

func constructFullImage(tiles []*imageTile) (*fullImageContainer, error) {
	dimAsFloat := math.Sqrt(float64(len(tiles)))
	dim := int(dimAsFloat)

	data := make([][]*imageTile, dim)
	for i := 0; i < dim; i++ {
		data[i] = make([]*imageTile, dim)
	}

	image := &fullImageContainer{
		dim:  dim,
		data: data,
	}

	unusedTiles := make([]*imageTile, len(tiles))
	for i, unusedTile := range tiles {
		unusedTiles[i] = unusedTile
	}

	tilesFit := image.fitTile(tiles, unusedTiles, 0, 0)
	if !tilesFit {
		return nil, errors.New("unable to fit tiles into image")
	}

	return image, nil
}

func (f *fullImageContainer) fitTile(tiles []*imageTile, unusedTiles []*imageTile, i int, j int) bool {
	if len(unusedTiles) == 0 {
		return true
	}

	nextI, nextJ := i, j+1
	if nextJ == f.dim {
		nextI, nextJ = i+1, 0
	}

	for k, tile := range unusedTiles {
		newUnusedTiles := make([]*imageTile, 0, len(tiles)-1)
		for l, unusedTile := range unusedTiles {
			if l != k {
				newUnusedTiles = append(newUnusedTiles, unusedTile)
			}
		}

		cloneTile := tile.Clone()
		flippedTile := cloneTile.FlipAcrossXAxis()
		flippedAndRotated90 := flippedTile.Rotate90()
		flippedAndRotate180 := flippedAndRotated90.Rotate90()
		flippedAndRotate270 := flippedAndRotate180.Rotate90()
		rotated90Tile := cloneTile.Rotate90()
		rotated180Tile := rotated90Tile.Rotate90()
		rotated270Tile := rotated180Tile.Rotate90()

		tileVariations := []*imageTile{
			cloneTile,
			flippedTile, flippedAndRotated90, flippedAndRotate180, flippedAndRotate270,
			rotated90Tile, rotated180Tile, rotated270Tile,
		}

		for _, variation := range tileVariations {
			if i > 0 {
				if !variation.TopEdgeMatchesUp(f.data[i-1][j]) {
					continue
				}
			}
			if j > 0 {
				if !variation.LeftEdgeMatchesUp(f.data[i][j-1]) {
					continue
				}
			}

			// Tile fits, place it.
			f.data[i][j] = variation

			imageIsSuccessful := f.fitTile(tiles, newUnusedTiles, nextI, nextJ)
			if imageIsSuccessful {
				return true
			}

			f.data[i][j] = nil
		}
	}

	return false
}

func (f *fullImageContainer) TopLeftTileID() uint64 {
	return f.data[0][0].id
}

func (f *fullImageContainer) TopRightTileID() uint64 {
	return f.data[0][f.dim-1].id
}

func (f *fullImageContainer) BottomLeftTileID() uint64 {
	return f.data[f.dim-1][0].id
}

func (f *fullImageContainer) BottomRightTileID() uint64 {
	return f.data[f.dim-1][f.dim-1].id
}

func (f *fullImageContainer) String() string {
	pixelDim := f.dim * tileDim

	builder := strings.Builder{}
	for i := 0; i < pixelDim; i++ {
		for j := 0; j < pixelDim; j++ {
			tileI, tileJ := i/tileDim, j/tileDim
			pixelI, pixelJ := i%tileDim, j%tileDim

			r := f.data[tileI][tileJ].pixels[pixelI][pixelJ]

			builder.WriteRune(r)

			if j%tileDim == tileDim-1 {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune('\n')

		if i%tileDim == tileDim-1 {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

func (f *fullImageContainer) MergeTiles() *mergedImageContainer {
	fullDim := f.dim * tileDim
	mergedDim := f.dim * (tileDim - 2)

	pixels := make([][]rune, mergedDim)
	for i := 0; i < mergedDim; i++ {
		pixels[i] = make([]rune, mergedDim)
	}

	mergedI, mergedJ := 0, 0

	for i := 0; i < fullDim; i++ {
		if i%tileDim == 0 || i%tileDim == tileDim-1 {
			continue
		}

		for j := 0; j < fullDim; j++ {
			if j%tileDim == 0 || j%tileDim == tileDim-1 {
				continue
			}

			tileI, tileJ := i/tileDim, j/tileDim
			pixelI, pixelJ := i%tileDim, j%tileDim

			pixelValue := f.data[tileI][tileJ].pixels[pixelI][pixelJ]
			pixels[mergedI][mergedJ] = pixelValue

			mergedJ++
		}

		mergedI, mergedJ = mergedI+1, 0
	}

	return &mergedImageContainer{
		dim:    mergedDim,
		pixels: pixels,
	}
}

type imageTile struct {
	id     uint64
	pixels [tileDim][tileDim]rune
}

func (t *imageTile) String() string {
	builder := strings.Builder{}

	for i := 0; i < tileDim; i++ {
		for j := 0; j < tileDim; j++ {
			r := t.pixels[i][j]
			if r == '.' {
				r = ' '
			}
			builder.WriteRune(r)
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

func (t *imageTile) TopEdgeMatchesUp(above *imageTile) bool {
	for j := 0; j < tileDim; j++ {
		if t.pixels[0][j] != above.pixels[tileDim-1][j] {
			return false
		}
	}
	return true
}

func (t *imageTile) LeftEdgeMatchesUp(adjacent *imageTile) bool {
	for i := 0; i < tileDim; i++ {
		if adjacent.pixels[i][tileDim-1] != t.pixels[i][0] {
			return false
		}
	}
	return true
}

func (t *imageTile) Clone() *imageTile {
	pixels := [tileDim][tileDim]rune{}
	for i := 0; i < tileDim; i++ {
		for j := 0; j < tileDim; j++ {
			pixels[i][j] = t.pixels[i][j]
		}
	}
	return &imageTile{
		id:     t.id,
		pixels: pixels,
	}
}

func (t *imageTile) Rotate90() *imageTile {
	pixels := [tileDim][tileDim]rune{}

	for i := 0; i < tileDim/2; i++ {
		for j := i; j < tileDim-i-1; j++ {
			pixels[i][j] = t.pixels[tileDim-1-j][i]
			pixels[tileDim-1-j][i] = t.pixels[tileDim-1-i][tileDim-1-j]
			pixels[tileDim-1-i][tileDim-1-j] = t.pixels[j][tileDim-1-i]
			pixels[j][tileDim-1-i] = t.pixels[i][j]
		}
	}

	return &imageTile{
		id:     t.id,
		pixels: pixels,
	}
}

func (t *imageTile) FlipAcrossXAxis() *imageTile {
	pixels := [tileDim][tileDim]rune{}

	for i := 0; i < tileDim; i++ {
		for j := 0; j < tileDim; j++ {
			pixels[i][j] = t.pixels[tileDim-i-1][j]
		}
	}

	return &imageTile{
		id:     t.id,
		pixels: pixels,
	}
}

const seaMonsterImage = "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "

type mergedImageContainer struct {
	dim    int
	pixels [][]rune
}

func (m *mergedImageContainer) Rotate90() {
	for i := 0; i < m.dim/2; i++ {
		for j := i; j < m.dim-i-1; j++ {
			temp := m.pixels[i][j]
			m.pixels[i][j] = m.pixels[m.dim-j-1][i]
			m.pixels[m.dim-j-1][i] = m.pixels[m.dim-i-1][m.dim-j-1]
			m.pixels[m.dim-i-1][m.dim-j-1] = m.pixels[j][m.dim-i-1]
			m.pixels[j][m.dim-i-1] = temp
		}
	}
}

func (m *mergedImageContainer) FlipAcrossXAxis() {
	for i := 0; i < m.dim/2; i++ {
		for j := 0; j < m.dim; j++ {
			temp := m.pixels[i][j]
			m.pixels[i][j] = m.pixels[m.dim-i-1][j]
			m.pixels[m.dim-i-1][j] = temp
		}
	}
}

func (m *mergedImageContainer) Clone() *mergedImageContainer {
	pixels := make([][]rune, m.dim)
	for i := 0; i < m.dim; i++ {
		pixels[i] = make([]rune, m.dim)
	}

	for i := 0; i < m.dim; i++ {
		for j := 0; j < m.dim; j++ {
			pixels[i][j] = m.pixels[i][j]
		}
	}

	return &mergedImageContainer{
		dim:    m.dim,
		pixels: m.pixels,
	}
}

func (m *mergedImageContainer) ManipulateUntilSeaMonstersTagged() int {
	// Original.
	nMonstersTagged := m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Rotated 90 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Rotated 180 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Rotated 270 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Flipped, original rotation.
	m.Rotate90()
	m.FlipAcrossXAxis()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Flipped, rotated 90 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Flipped, rotated 180 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// Flipped, rotated 270 degrees.
	m.Rotate90()
	nMonstersTagged = m.TagSeaMonsters(seaMonsterImage)
	if nMonstersTagged > 0 {
		return nMonstersTagged
	}

	// We failed to find sea monsters.
	m.Rotate90()
	m.FlipAcrossXAxis()
	return 0
}

func (m *mergedImageContainer) TagSeaMonsters(seaMonsterImage string) int {
	monsterHeight, monsterLength, visiblePortionOffsets := examineSeaMonsterImage(seaMonsterImage)

	nMonstersTagged := 0
	for i := 0; i < m.dim-monsterHeight+1; i++ {
		for j := 0; j < m.dim-monsterLength+1; j++ {
			allVisiblePortionsPresent := true
			for _, offsets := range visiblePortionOffsets {
				if m.pixels[i+offsets.I][j+offsets.J] != '#' {
					allVisiblePortionsPresent = false
					break
				}
			}

			if allVisiblePortionsPresent {
				for _, offsets := range visiblePortionOffsets {
					m.pixels[i+offsets.I][j+offsets.J] = 'O'
				}
				nMonstersTagged++
			}
		}
	}

	return nMonstersTagged
}

func (m *mergedImageContainer) DetermineWaterRoughness() int {
	nRoughSpots := 0
	for i := 0; i < m.dim; i++ {
		for j := 0; j < m.dim; j++ {
			if m.pixels[i][j] == '#' {
				nRoughSpots++
			}
		}
	}
	return nRoughSpots
}

func (m *mergedImageContainer) String() string {
	builder := strings.Builder{}
	for i := 0; i < m.dim; i++ {
		for j := 0; j < m.dim; j++ {
			builder.WriteRune(m.pixels[i][j])
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func examineSeaMonsterImage(seaMonsterImage string) (int, int, []grid.Point) {
	monsterHeight, monsterLength := 0, 0
	visiblePortionOffsets := []grid.Point{}

	i, j := 0, 0
	for _, r := range seaMonsterImage {
		switch r {
		case '#':
			visiblePortionOffsets = append(visiblePortionOffsets, grid.Point{I: i, J: j})
			j++
		case '\n':
			monsterLength = j
			i, j = i+1, 0
		default:
			j++
		}
	}
	monsterHeight = i + 1

	return monsterHeight, monsterLength, visiblePortionOffsets
}
