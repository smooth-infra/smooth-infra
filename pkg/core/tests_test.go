package core

import (
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/yaml"
	"github.com/stretchr/testify/require"
)

func TestIfCanExecuteTests(t *testing.T) {
	yamlConfig := &yaml.BaseStructure{
		Version: 1,
		Tests: []yaml.Test{
			{
				Name:    "human friendly test name",
				Id:      "computer_friendly_test_name",
				Type:    "null/null",
				Params:  map[string]interface{}{},
				Expects: map[string]interface{}{},
			},
		},
	}

	errors := ExecuteTests(t, yamlConfig)
	require.Empty(t, errors)
}
