package yaml

import (
	"fmt"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	"github.com/smooth-infra/smooth-infra/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessSimpleYaml(t *testing.T) {
	config, err := Process(getTestData("output.vars"))
	require.Nil(t, err)
	assert.Equal(t, "output.vars", config.Input["terraform"].OutputsFile, "The outputs_file is not what is expected")
}

func TestProcessingOfInvalidYaml(t *testing.T) {
	_, err := Process("invalid yaml")
	require.NotNil(t, err)
}

func TestIfCanDetermineInputName(t *testing.T) {
	config, err := Process(getTestData("output.vars"))
	require.Nil(t, err)
	assert.Equal(t, "terraform", config.GetInputName(), "The input name is not what is expected")
}

func TestIfCanReplaceTerraformOutputsToTests(t *testing.T) {
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

	yamlTestFile := getTestData(tempFile.Name())

	outputValues, err := terraform.GetOutputValues(tempFile.Name())
	require.Nil(t, err)

	config, err := Process(yamlTestFile)
	require.Nil(t, err)

	config.ProcessTests(config.GetInputName(), outputValues)

	assert.Equal(t, "Verify that requesting google.com is giving a 200 OK", config.Tests[0].Name, "The name does not match the expected value")
	assert.Equal(t, "google.com", config.Tests[0].Params["address"], "The address does not match the expected value")
}

func getTestData(outputsFilePath string) string {
	return fmt.Sprintf(`
---
version: 1
input:
  terraform:
    outputs_file: %s
tests:
  - name: Verify that requesting ${input.terraform.address} is giving a 200 OK
    type: http/request
    params:
      address: ${input.terraform.address}
      secure: true
    expects:
      status_code: 200
    `, outputsFilePath,
	)
}
