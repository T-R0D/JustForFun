package day18

import (
	"aoc2019/internal/location"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	w := inputToWorld(input)
	steps := findAllKeys(w)
	return steps, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	w := inputToWorld(input)
	w.updateTo4BotWorld()
	totalSteps := findAllKeys2(w)
	return totalSteps, nil
}

const (
	WALL     = int('#')
	EMPTY    = int('.')
	ENTRANCE = int('@')
)

type doorSet map[int]struct{}

type keySet map[int]struct{}

func (k keySet) Clone() keySet {
	clone := keySet{}
	for key := range k {
		clone[key] = struct{}{}
	}
	return clone
}

func (k keySet) String() string {
	l := make([]int, 0, len(k))
	for key := range k {
		l = append(l, key)
	}
	sort.Ints(l)
	var b strings.Builder
	for i, k := range l {
		b.WriteString(fmt.Sprintf("%c", rune(k)))
		if i < len(l)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}

func doorToKey(door int) int {
	return door + 0x20
}

func isDoor(obj int) bool {
	return int('A') <= obj && obj <= int('Z')
}

func isKey(obj int) bool {
	return int('a') <= obj && obj <= int('z')
}

type world struct {
	entrance  location.Point
	entrances [4]location.Point
	m         map[location.Point]int
	keys      map[int]location.Point
	heldKeys  map[int]struct{}
	doors     map[int]location.Point
}

func inputToWorld(input string) *world {
	m := make(map[location.Point]int)
	keys := make(map[int]location.Point)
	doors := make(map[int]location.Point)
	var entrance location.Point

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, obj := range line {
			o := int(obj)
			loc := location.Point{X: j, Y: i}
			m[loc] = o
			if o == ENTRANCE {
				entrance = loc
			} else if 'a' <= obj && obj <= 'z' {
				keys[o] = loc
			} else if 'A' <= obj && obj <= 'Z' {
				doors[o] = loc
			}
		}
	}

	return &world{
		entrance: entrance,
		m:        m,
		keys:     keys,
		heldKeys: make(map[int]struct{}),
		doors:    doors,
	}
}

func (w *world) updateTo4BotWorld() {
	w.entrances = [4]location.Point{
		location.Point{X: w.entrance.X - 1, Y: w.entrance.Y - 1},
		location.Point{X: w.entrance.X - 1, Y: w.entrance.Y + 1},
		location.Point{X: w.entrance.X + 1, Y: w.entrance.Y - 1},
		location.Point{X: w.entrance.X + 1, Y: w.entrance.Y + 1},
	}

	for _, e := range w.entrances {
		w.m[e] = ENTRANCE
	}

	w.m[location.Point{X: w.entrance.X, Y: w.entrance.Y}] = WALL
	w.m[location.Point{X: w.entrance.X, Y: w.entrance.Y + 1}] = WALL
	w.m[location.Point{X: w.entrance.X + 1, Y: w.entrance.Y}] = WALL
	w.m[location.Point{X: w.entrance.X, Y: w.entrance.Y - 1}] = WALL
	w.m[location.Point{X: w.entrance.X - 1, Y: w.entrance.Y}] = WALL
}

func (w *world) dumpState() string {
	maxX := 0
	x := 0
	_, ok := w.m[location.Point{X: x, Y: 0}]
	for ok {
		maxX = x
		x += 1
		_, ok = w.m[location.Point{X: x, Y: 0}]
	}

	maxY := 0
	y := 0
	_, ok = w.m[location.Point{X: 0, Y: y}]
	for ok {
		maxY = y
		y += 1
		_, ok = w.m[location.Point{X: 0, Y: y}]
	}

	var b strings.Builder
	for y = 0; y <= maxY; y += 1 {
		for x = 0; x <= maxX; x += 1 {
			obj := w.m[location.Point{X: x, Y: y}]
			if obj == WALL {
				b.WriteRune('.')
			} else if obj == EMPTY {
				b.WriteRune(' ')
			} else {
				b.WriteRune(rune(obj))
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (w *world) getKeySet() keySet {
	ks := keySet{}
	for k := range w.keys {
		ks[k] = struct{}{}
	}
	return ks
}

func findAllKeys(w *world) int {
	cache := make(map[string]int)
	keyRequirements := generateLocToKeyRequirements(w)
	return distanceToFindAllKeys(w, w.entrance, w.getKeySet(), cache, keyRequirements)
}

func distanceToFindAllKeys(w *world, loc location.Point, keysRemaining keySet,
	cache map[string]int, keyRequirements map[int]map[int]keyRequirement) int {

	obj := w.m[loc]
	if !(isKey(obj) || obj == ENTRANCE) {
		panic("obj is not key or entrance")
	}

	if len(keysRemaining) == 0 {
		return 0
	}

	cacheKey := createCacheKey(obj, keysRemaining)
	if dist, ok := cache[cacheKey]; ok {
		return dist
	}

	result := math.MaxInt32
	reachableKeys := getReachableKeys(obj, keysRemaining, keyRequirements)
	for key := range reachableKeys {
		d := keyRequirements[obj][key].dist
		keysRem := keysRemaining.Clone()
		delete(keysRem, key)
		dist := distanceToFindAllKeys(w, w.keys[key], keysRem, cache, keyRequirements)
		d += dist

		result = int(math.Min(float64(d), float64(result)))
	}
	cache[cacheKey] = result
	return result
}

func getReachableKeys(obj int, keysRemaining keySet, keyRequirements map[int]map[int]keyRequirement) keySet {
	reachableKeys := keySet{}
	for key := range keysRemaining {
		req := keyRequirements[obj][key]
		needKeyToGetKey := false
		for door := range req.doors {
			doorKey := doorToKey(door)
			if _, ok := keysRemaining[doorKey]; ok {
				needKeyToGetKey = true
				break
			}
		}
		if needKeyToGetKey {
			continue
		}
		reachableKeys[key] = struct{}{}
	}
	return reachableKeys
}

func createCacheKey(obj1 int, keysRemaining keySet) string {
	return fmt.Sprintf("%c [%v]", rune(obj1), keysRemaining.String())
}

type searchState struct {
	heldKeys keySet
	loc      location.Point
	steps    int
}

func (s searchState) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("loc: %v\n", s.loc))
	keys := make([]int, 0, len(s.heldKeys))
	for k := range s.heldKeys {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	b.WriteString(fmt.Sprintf("keys: %v", keys))
	return b.String()
}

type keyRequirement struct {
	dist  int
	doors doorSet
}

func generateLocToKeyRequirements(w *world) map[int]map[int]keyRequirement {
	reqs := make(map[int]map[int]keyRequirement)
	reqs[ENTRANCE] = findKeyRequirements(w, w.entrance)
	for k1 := range w.keys {
		reqs[k1] = findKeyRequirements(w, w.keys[k1])
	}
	return reqs
}

type keyToKeyState struct {
	doors doorSet
	loc   location.Point
	steps int
}

func (k keyToKeyState) Clone() keyToKeyState {
	doors := doorSet{}
	for d := range k.doors {
		doors[d] = struct{}{}
	}
	return keyToKeyState{
		doors: doors,
		loc:   k.loc,
		steps: k.steps,
	}
}

func findKeyRequirements(w *world, start location.Point) map[int]keyRequirement {
	keyRequirements := make(map[int]keyRequirement)

	q := list.New()
	seen := make(map[location.Point]struct{})
	q.PushBack(keyToKeyState{
		doors: doorSet{},
		loc:   start,
		steps: 0,
	})

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(keyToKeyState)

		// Processing that needs to be done at the location.
		obj := w.m[current.loc]

		if isDoor(obj) {
			current.doors[obj] = struct{}{}
		}

		// Continuation/Completion checks.
		if _, ok := seen[current.loc]; ok {
			continue
		}

		if isKey(obj) && current.loc != start {
			clone := current.Clone()
			keyRequirements[obj] = keyRequirement{
				dist:  clone.steps,
				doors: clone.doors,
			}
		}

		// Generate next states.
		x, y := current.loc.X, current.loc.Y

		nextLoc := location.Point{X: x + 1, Y: y}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x - 1, Y: y}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x, Y: y + 1}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x, Y: y - 1}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		// Current state has been fully explored.
		seen[current.loc] = struct{}{}
	}

	return keyRequirements
}

func canMoveToSpace(w *world, nextLoc location.Point) bool {
	if obj, ok := w.m[nextLoc]; !ok || obj == WALL {
		return false
	}
	return true
}

func findAllKeys2(w *world) int {
	keyRequirements := generateAllKeyRequirements(w)
	cache := make(map[string]int)
	totalSteps := findDistanceToCollectKeys(w, w.entrances[:], w.getKeySet(), keyRequirements, cache)
	return totalSteps
}

func generateAllKeyRequirements(w *world) map[location.Point]map[location.Point]keyRequirement {
	keyRequirements := make(map[location.Point]map[location.Point]keyRequirement)
	for _, e := range w.entrances {
		keyRequirements[e] = generateKeyRequirements(w, e)
	}
	for k := range w.getKeySet() {
		loc := w.keys[k]
		keyRequirements[loc] = generateKeyRequirements(w, loc)
	}
	return keyRequirements
}

type keyRequirementsSearchState struct {
	doors doorSet
	loc   location.Point
	steps int
}

func (k keyRequirementsSearchState) Clone() keyRequirementsSearchState {
	doors := doorSet{}
	for door := range k.doors {
		doors[door] = struct{}{}
	}
	return keyRequirementsSearchState{
		doors: doors,
		loc:   k.loc,
		steps: k.steps,
	}
}

func generateKeyRequirements(w *world, origin location.Point) map[location.Point]keyRequirement {
	keyRequirements := make(map[location.Point]keyRequirement)

	q := list.New()
	seen := make(map[location.Point]struct{})
	q.PushBack(keyRequirementsSearchState{
		doors: doorSet{},
		loc:   origin,
		steps: 0,
	})

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(keyRequirementsSearchState)

		// Ensure we haven't already visited this state.
		if _, ok := seen[current.loc]; ok {
			continue
		}

		// Evaluate things at the current location.
		obj := w.m[current.loc]
		if isDoor(obj) {
			current.doors[obj] = struct{}{}
		} else if isKey(obj) && current.loc != origin {
			clone := current.Clone()
			keyRequirements[current.loc] = keyRequirement{
				dist:  clone.steps,
				doors: clone.doors,
			}
		}

		// Generate next states.
		x, y := current.loc.X, current.loc.Y

		nextLoc := location.Point{X: x + 1, Y: y}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x - 1, Y: y}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x, Y: y + 1}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		nextLoc = location.Point{X: x, Y: y - 1}
		if canMoveToSpace(w, nextLoc) {
			next := current.Clone()
			next.loc = nextLoc
			next.steps += 1
			q.PushBack(next)
		}

		// Current state has been fully explored.
		seen[current.loc] = struct{}{}
	}

	return keyRequirements
}

func findDistanceToCollectKeys(w *world, currentLocs []location.Point, unfoundKeys keySet,
	keyRequirements map[location.Point]map[location.Point]keyRequirement, cache map[string]int) int {

	if len(unfoundKeys) == 0 {
		return 0
	}

	cacheKey := generateCacheKey(currentLocs, unfoundKeys)

	if steps, ok := cache[cacheKey]; ok {
		return steps
	}

	steps := math.MaxInt32
	for i, loc := range currentLocs {
		reachableKeys := computeReachableKeys(loc, unfoundKeys, w, keyRequirements)
		for key := range reachableKeys {
			keyLoc := w.keys[key]

			stepsToKey := keyRequirements[loc][keyLoc].dist

			unfound := unfoundKeys.Clone()
			delete(unfound, key)

			nextLocs := make([]location.Point, len(currentLocs))
			for j, l := range currentLocs {
				if i == j {
					nextLocs[j] = keyLoc
				} else {
					nextLocs[j] = l
				}
			}

			d := stepsToKey + findDistanceToCollectKeys(w, nextLocs, unfound, keyRequirements, cache)

			steps = int(math.Min(float64(steps), float64(d)))
		}
	}

	cache[cacheKey] = steps
	return steps
}

func generateCacheKey(currentLocs []location.Point, unfoundKeys keySet) string {
	return fmt.Sprintf("%v | %s", currentLocs, unfoundKeys.String())
}

func computeReachableKeys(origin location.Point, unfoundKeys keySet, w *world,
	allKeyRequirements map[location.Point]map[location.Point]keyRequirement) keySet {

	reachableKeys := keySet{}
	keyRequirements := allKeyRequirements[origin]
	for key := range unfoundKeys {
		keyLoc := w.keys[key]

		keyReachable := false
		for k := range keyRequirements {
			if k == keyLoc {
				keyReachable = true
			}
		}
		if !keyReachable {
			continue
		}

		missingKeyToGetKey := false
		req := keyRequirements[keyLoc]
		for door := range req.doors {
			if _, ok := unfoundKeys[doorToKey(door)]; ok {
				missingKeyToGetKey = true
				break
			}
		}
		if missingKeyToGetKey {
			continue
		}

		reachableKeys[key] = struct{}{}
	}
	return reachableKeys
}

func dumpAllKeyRequirements(w *world, allReqs map[location.Point]map[location.Point]keyRequirement) {
	for k1, reqs := range allReqs {
		fmt.Printf("%c %v:\n", rune(w.m[k1]), k1)
		for k2, req := range reqs {
			fmt.Printf("\t%c %v --> %v\n", rune(w.m[k2]), k2, req)
		}
	}
}
