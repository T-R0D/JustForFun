package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSecretLoopSize(t *testing.T) {
	testCases := []struct {
		subjectNumber    uint64
		publicKey        uint64
		expectedLoopSize uint64
	}{
		{
			subjectNumber:    defaultSubjectNumber,
			publicKey:        5764801,
			expectedLoopSize: 8,
		},
		{
			subjectNumber:    defaultSubjectNumber,
			publicKey:        17807724,
			expectedLoopSize: 11,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("public key %d and subject number %d has a secret loop size of %d",
			tc.publicKey, tc.subjectNumber, tc.expectedLoopSize)
		t.Run(name, func(t *testing.T) {
			loopSize, err := findSecretLoopSize(tc.subjectNumber, tc.publicKey)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedLoopSize, loopSize)
		})
	}
}

func TestFindEncryptionKey(t *testing.T) {
	testCases := []struct {
		subjectNumber         uint64
		publicKeys            []uint64
		expectedEncryptionKey uint64
	}{
		{
			subjectNumber:         defaultSubjectNumber,
			publicKeys:            []uint64{5764801, 17807724},
			expectedEncryptionKey: 14897079,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("subject number %d and public keys %v produce encryption key %d",
			tc.subjectNumber, tc.publicKeys, tc.expectedEncryptionKey)
			t.Run(name, func(t *testing.T){
				encryptionKey, err := findEncryptionKey(tc.subjectNumber, tc.publicKeys)

				assert.NoError(t, err)
				assert.Equal(t, tc.expectedEncryptionKey, encryptionKey)
			})
	}
}
