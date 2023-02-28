package yaml

import (
	"fmt"
	"os"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIfCanDetermineInputName(t *testing.T) {
	structure, err := Process(getTestData("output.vars"))
	require.Nil(t, err)
	assert.Equal(t, "terraform", structure.GetInputName(), "The input name is not what is expected")
}

func TestIfCanReplaceTerraformOutputsToTests(t *testing.T) {
	testData := `
{
  "address": {
      "sensitive": false,
      "type": "string",
      "value": "google.com"
  }
}
`

	tempFile, err := os.CreateTemp("", "outputs.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	_, err = tempFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Error writing test data to file: %v", err)
	}
	tempFile.Close()

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

func TestProcessSimpleYaml(t *testing.T) {
	structure, err := Process(getTestData("output.vars"))
	require.Nil(t, err)
	assert.Equal(t, "output.vars", structure.Input["terraform"].OutputsFile, "The outputs_file is not what is expected")
}
