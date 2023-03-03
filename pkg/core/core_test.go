package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIfItCanRunASimpleSuccessfulTest(t *testing.T) {
	outputs := `
{
  "address": {
      "sensitive": false,
      "type": "string",
      "value": "google.com"
  }
}
`

	tempFile, cleanup := utilities.CreateTestFile(t, outputs)
	defer cleanup()

	yamlConfig := getTestConfig(tempFile.Name())

	errors := Run(t, strings.NewReader(yamlConfig))

	require.Empty(t, errors)
}

func TestIfItCanRunASimpleFailingTest(t *testing.T) {
	outputs := `
{
  "address": {
      "sensitive": false,
      "type": "string",
      "value": "somefailingwebsitethatdoesnotexist.com"
  }
}
`

	tempFile, cleanup := utilities.CreateTestFile(t, outputs)
	defer cleanup()
	yamlConfig := getTestConfig(tempFile.Name())

	errors := Run(t, strings.NewReader(yamlConfig))

	require.NotEmpty(t, errors)
	for k := range errors {
		errorName := k
		assert.Equal(t, "verify-that-requesting-somefailingwebsitethatdoesnotexist-com-is-giving-a-200-ok", errorName)
		break
	}
}

func getTestConfig(outputsFilePath string) string {
	return fmt.Sprintf(`
---
version: 1
input:
  terraform:
    outputs_file: %s
tests:
  - name: "Verify that requesting ${input.terraform.address} is giving a 200 OK"
    type: http/request
    params:
      address: ${input.terraform.address}
      secure: true
    expects:
      status_code: 200
    `, outputsFilePath,
	)
}
