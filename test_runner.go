package adventofcode

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T, path string, answer string, fn func(reader io.Reader) (string, error)) string {
	t.Helper()
	f, err := os.Open(path)
	require.NoError(t, err)

	defer f.Close()
	gotAnswer, err := fn(f)
	assert.NoError(t, err)
	assert.EqualValues(t, answer, gotAnswer)

	return gotAnswer
}
