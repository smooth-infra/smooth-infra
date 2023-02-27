package yaml

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

var data = `
---
version: 1
input:
  terraform:
    outputs_file: output.vars
  tests:
    - name: Verify that requesting ${input.terraform.address} is giving a 200 OK
      type: http/request
      params:
        address: ${input.terraform.address}
        secure: true
      expects:
        status_code: 200
`

func TestProcessSimpleYaml(t *testing.T) {
	structure := BaseStructure{}

	err := yaml.Unmarshal([]byte(data), &structure)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	assert.Equal(t, "output.vars", structure.Input.Terraform.OutputsFile, "The outputs_file is not what is expected")
}
