package yaml

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

func TestProcessSimpleYaml(t *testing.T) {
	var data = `
---
version: 1
input:
  terraform:
    outputs_file: stubs/output.vars
  tests:
    - name: Verify that requesting ${input.terraform.address} is giving a 200 OK
      type: http/request
      params:
        address: ${input.terraform.address}
        secure: true
      expects:
        status_code: 200
`
	structure := BaseStructure{}

	err := yaml.Unmarshal([]byte(data), &structure)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	assert.Equal(t, "stubs/output.vars", structure.Input.Terraform.OutputsFile, "The outputs_file is not what is expected")

	var myOutputs map[string]string
	myOutputs, err = godotenv.Read(structure.Input.Terraform.OutputsFile)

	assert.Equal(t, "google.com", myOutputs["address"], "The address in outputs_file is not what is expected")
}
