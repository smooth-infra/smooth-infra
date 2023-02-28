package null

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNull(t *testing.T) {
	require.Nil(t, Null(t, map[string]interface{}{}, map[string]interface{}{}))
}
