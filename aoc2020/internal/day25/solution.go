package day25

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver saves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	publicKeys, err := readPublicKeys(input)
	if err != nil {
		return "", err
	}

	encryptionKey, err := findEncryptionKey(defaultSubjectNumber, publicKeys)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(encryptionKey), nil
}

// Part2 solves part 2of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	return "Merry Christmas!", nil
}

const (
	cardIndex = 0
	doorIndex = 1
)

const (
	defaultSubjectNumber = uint64(7)

	magicDivisor = uint64(20201227)
)

func readPublicKeys(input string) ([]uint64, error) {
	keyStrs := strings.Split(input, "\n")
	if len(keyStrs) != 2 {
		return nil, errors.Errorf("Got %d keys, not 2", len(keyStrs))
	}

	keys := make([]uint64, 2)
	for i, keyStr := range keyStrs {
		value, err := strconv.Atoi(keyStr)
		if err != nil {
			return nil, errors.Errorf("key %d was not an integer", i)
		}
		keys[i] = uint64(value)
	}
	return keys, nil
}

func findEncryptionKey(subjectNumber uint64, publicKeys []uint64) (uint64, error) {
	loopSizes := make([]uint64, len(publicKeys))
	for i, publicKey := range publicKeys {
		loopSize, err := findSecretLoopSize(subjectNumber, publicKey)
		if err != nil {
			return 0, errors.Wrapf(err, "finding secret loop size for key %d", i)
		}
		loopSizes[i] = loopSize
	}

	encryptionKeys := make([]uint64, len(publicKeys))
	for i := 0; i < len(publicKeys); i++ {
		otherDeviceIndex := (i + 1) % len(publicKeys)
		encryptionKeys[i] = transformNumber(publicKeys[otherDeviceIndex], loopSizes[i])
	}

	if encryptionKeys[cardIndex] != encryptionKeys[doorIndex] {
		return 0, errors.Errorf("encryption keys don't match! (%d != %d)", encryptionKeys[0], encryptionKeys[1])
	}

	return encryptionKeys[cardIndex], nil
}

func findSecretLoopSize(subjectNumber uint64, publicKey uint64) (uint64, error) {
	loopSize := uint64(0)
	for value := uint64(1); value != publicKey; loopSize++ {
		value *= subjectNumber
		value %= magicDivisor
	}
	return loopSize, nil
}

func transformNumber(subjectNumber uint64, loopSize uint64) uint64 {
	value := uint64(1)
	for i := uint64(0); i < loopSize; i++ {
		value *= subjectNumber
		value %= magicDivisor	
	}
	return value
}
