package terraform

import (
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanFetchTerraformOutputValuesFromFile(t *testing.T) {
	t.Parallel()

	outputs := `
    {
      "test_string": {
        "sensitive": false,
        "type": "string",
        "value": "this_is_a_test_string"
      },
      "test_bool": {
        "sensitive": false,
        "type": "bool",
        "value": true
      },
      "test_list": {
        "sensitive": false,
        "type": "list",
        "value": ["item_1", "item_2"]
      },
      "test_map": {
        "sensitive": false,
        "type": "map",
        "value": {
          "key": "value"
        }
      },
      "test_complex": {
        "sensitive": false,
        "type": "map",
        "value": {
          "object": {
            "nested1a": {
              "nested2a_key": "nested2value",
              "nested2b_key": true
            },
            "nested1b": true,
            "nested1c": ["value1", "value2"]
          }
        }
      }
    }
    `

	tempFile, cleanup := utilities.CreateTestFile(t, outputs)
	defer cleanup()

	errMsg := "The value does not match the expected one"

	input, err := GetOutputValues(tempFile.Name())
	require.Nil(t, err, "There were some errors")
	assert.Equal(t, "this_is_a_test_string", input.Variables["test_string"].Value, errMsg)
	assert.Equal(t, true, input.Variables["test_bool"].Value, errMsg)
	assert.Equal(t, []interface{}{"item_1", "item_2"}, input.Variables["test_list"].Value, errMsg)
	assert.Equal(t, map[string]interface{}{"key": "value"}, input.Variables["test_map"].Value, errMsg)
	assert.Equal(t, map[string]interface{}{
		"object": map[string]interface{}{
			"nested1a": map[string]interface{}{
				"nested2a_key": "nested2value", "nested2b_key": true,
			},
			"nested1b": true,
			"nested1c": []interface{}{"value1", "value2"},
		},
	}, input.Variables["test_complex"].Value)
}
