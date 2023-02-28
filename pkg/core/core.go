package core

import (
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/http"
	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	"github.com/smooth-infra/smooth-infra/pkg/yaml"

	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

type TestFunction func(t *testing.T, params map[string]interface{}, expects map[string]interface{}) error

var availableTests = map[string]TestFunction{
	"http/request": http.Request,
}

func RunTests(t *testing.T, yamlFile string) map[string]error {
	config, err := yaml.Process(yamlFile)
	require.Nil(t, err)

	outputValues, err := terraform.GetOutputValues(config.Input[config.GetInputName()].OutputsFile)
	require.Nil(t, err)

	errors := make(map[string]error)

	config.ProcessTests(config.GetInputName(), outputValues)

	for _, test := range config.Tests {
		if function, ok := availableTests[test.Type]; ok {
			t.Logf("Running \"%s\" (%s) test...", test.Name, test.Type)
			err = function(t, test.Params, test.Expects)
			if err != nil {
				slug := slug.Make(test.Name)
				errors[slug] = err
			}
		}
	}
	return errors
}
