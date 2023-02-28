package http

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIfCanMakeRequestGivenParamsAndExpectedStatusCode(t *testing.T) {
	params := map[string]interface{}{"address": "httpstat.us/200", "secure": false}
	expects := map[string]interface{}{"status_code": 200}

	err := Request(t, params, expects)
	require.Nil(t, err)
}
