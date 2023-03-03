package core

import (
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIfCanExecuteFailingTests(t *testing.T) {
	t.Parallel()

	yamlConfig := &yaml.BaseStructure{
		Version: 1,
		Tests: []yaml.Test{
			{
				Name: "human friendly test name",
				Id:   "computer_friendly_test_name",
				Type: "http/request",
				Params: map[string]interface{}{
					"address": "fake_domain.com",
					"secure":  false,
				},
				Expects: map[string]interface{}{
					"status_code": 200,
				},
			},
		},
	}

	errors := ExecuteTests(t, yamlConfig)
	require.NotEmpty(t, errors)
	for k := range errors {
		errorName := k
		assert.Equal(t, "computer_friendly_test_name", errorName)
		break
	}
}
