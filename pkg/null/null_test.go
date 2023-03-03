package null

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNull(t *testing.T) {
	t.Parallel()

	require.Nil(t, Null(t, map[string]interface{}{}, map[string]interface{}{}))
}
