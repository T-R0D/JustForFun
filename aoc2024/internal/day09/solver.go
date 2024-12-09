package day09

import (
	"fmt"
	"slices"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	compactRepr, err := parseCompactedDiskRepresentation(input)
	if err != nil {
		return "", err
	}

	looseRepr := compactRepresentationToLooseRepresentation(compactRepr)
	compactedLooseRepr, blocksUsed := compactDisk(looseRepr)
	checksum := checksumLooseRepresentation(compactedLooseRepr[:blocksUsed])

	return fmt.Sprintf("%d", checksum), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	compactRepr, err := parseCompactedDiskRepresentation(input)
	if err != nil {
		return "", err
	}

	compactedLooseRepr := compactDiskWithoutFragmentingFiles(compactRepr)
	checksum := checksumLooseRepresentation(compactedLooseRepr)

	return fmt.Sprintf("%d", checksum), nil
}

func parseCompactedDiskRepresentation(input string) ([]int64, error) {
	repr := make([]int64, len(input))
	for i, r := range input {
		digitValue := int64(r - '0')
		if digitValue < 0 || 9 < digitValue {
			return []int64{}, fmt.Errorf("item %d was not a digit ('%c')", i, r)
		}
		repr[i] = digitValue
	}

	return repr, nil
}

func compactRepresentationToLooseRepresentation(compactRepr []int64) []int64 {
	looseRepr := make([]int64, 0, len(compactRepr))
	for i, nBlocks := range compactRepr {
		for range nBlocks {
			id := int64(i / 2)
			if i&0x01 == 1 {
				looseRepr = append(looseRepr, -1)
			} else {
				looseRepr = append(looseRepr, id)
			}
		}
	}

	return looseRepr
}

func compactDisk(looseRepr []int64) ([]int64, int) {
	compactedLooseRepr := slices.Clone(looseRepr)
	i, j := 0, len(compactedLooseRepr)-1
	for ; compactedLooseRepr[i] != -1 && i < j; i += 1 {
	}
	for ; compactedLooseRepr[j] == -1 && i < j; j -= 1 {
	}
	for i < j {
		compactedLooseRepr[i], compactedLooseRepr[j] = compactedLooseRepr[j], compactedLooseRepr[i]

		for ; compactedLooseRepr[i] != -1 && i < j; i += 1 {
		}
		for ; compactedLooseRepr[j] == -1 && i < j; j -= 1 {
		}
	}

	for ; i < len(compactedLooseRepr) && compactedLooseRepr[i] != -1; i += 1 {
	}

	return compactedLooseRepr, i
}

func compactDiskWithoutFragmentingFiles(compactRepr []int64) []int64 {
	type freeSpaceSpec struct {
		start int
		size  int
	}

	freeSpaces := make([]*freeSpaceSpec, 0, len(compactRepr)/2)

	for i, offset := 0, 0; i < len(compactRepr); i += 1 {
		objectSize := compactRepr[i]
		if i&0x01 == 1 {
			freeSpaces = append(freeSpaces, &freeSpaceSpec{start: offset, size: int(objectSize)})
		}
		offset += int(objectSize)
	}

	looseRepr := compactRepresentationToLooseRepresentation(compactRepr)

	i := len(looseRepr) - 1
	for ; i >= 0 && looseRepr[i] == -1; i -= 1 {
	}
COMPACTION_SCAN:
	for i >= 0 {
		currentId := looseRepr[i]
		if currentId == -1 {
			continue
		}

		fileSize := 0
		for j := i; j >= 0 && currentId == looseRepr[j]; j, fileSize = j-1, fileSize+1 {
		}

		var freeSpace *freeSpaceSpec = nil
	FREE_SPACE_SEARCH:
		for _, space := range freeSpaces {
			if space.start >= i - fileSize {
				break FREE_SPACE_SEARCH
			}

			if space.size >= fileSize {
				freeSpace = space
				break FREE_SPACE_SEARCH
			}
		}

		if freeSpace == nil {
			for ; i >= 0 && (looseRepr[i] == -1 || looseRepr[i] == currentId); i -= 1 {
			}
			continue COMPACTION_SCAN
		}

		for j := range fileSize {
			looseRepr[j+freeSpace.start] = currentId
			looseRepr[i-j] = -1
		}

		freeSpace.size -= fileSize
		freeSpace.start += fileSize

		for ; i >= 0 && (looseRepr[i] == -1 || looseRepr[i] == currentId); i -= 1 {
		}
	}

	return looseRepr
}

func checksumLooseRepresentation(looseRepr []int64) int64 {
	checksum := int64(0)
	for i, id := range looseRepr {
		if id == -1 {
			continue
		}

		checksum += int64(i) * id
	}

	return checksum
}
