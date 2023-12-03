package aoc_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Read the contents of a file into a string.
func readFile(t *testing.T, filename string) string {
	t.Helper()

	inputFile, err := os.Open(filename)
	require.NoError(t, err)
	defer func() {
		if closeErr := inputFile.Close(); err != nil {
			panic(closeErr)
		}
	}()

	// Find the length of the file.
	fileInfo, err := inputFile.Stat()
	require.NoError(t, err)
	fileSize := fileInfo.Size()

	// Make a buffer of this length and read into it.
	buf := make([]byte, fileSize)
	readLen, err := inputFile.Read(buf)
	require.NoError(t, err)
	assert.EqualValues(t, fileSize, readLen)

	// Trim any surrounding whitespace before returning.
	return strings.TrimSpace(string(buf))
}
