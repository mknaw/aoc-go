package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitTestCases(t *testing.T) {

	input := []byte(`This is some text.
>>> 42

Another line of text.
Yet another line.
>>> 123`)

	result, err := splitTestCases(input)
	require.NoError(t, err)
	require.Len(t, result, 2)
	assert.Equal(t, "This is some text.\n", result[0].Input)
	assert.Equal(t, 42, result[0].Expected)
	assert.Equal(t, "Another line of text.\nYet another line.\n", result[1].Input)
	assert.Equal(t, 123, result[1].Expected)
}

func TestSplitTestCases_EmptyString(t *testing.T) {

	input := []byte("")

	result, err := splitTestCases(input)
	require.NoError(t, err)
	require.Len(t, result, 0)
}
