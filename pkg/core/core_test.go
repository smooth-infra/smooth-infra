package core

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIfItCanRunASimpleSuccessfulTest(t *testing.T) {
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

	errors := RunTests(t, yamlTestFile)

	require.Empty(t, errors)
}

func TestIfItCanRunASimpleFailingTest(t *testing.T) {
	testData := `
{
  "address": {
      "sensitive": false,
      "type": "string",
      "value": "somefailingwebsitethatdoesnotexist.com"
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

	errors := RunTests(t, yamlTestFile)

	require.NotEmpty(t, errors)
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
