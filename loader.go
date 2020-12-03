package adventofcode

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func Test(t *testing.T, path string, fn func(reader io.Reader) (string, error)) string {
	t.Helper()
	f, err := os.Open(path)
	require.NoError(t, err)

	defer f.Close()
	ans, err := fn(f)
	require.NoError(t, err)
	fmt.Printf("answer = %+v\n", ans)

	return ans
}
