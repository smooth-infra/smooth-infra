package core

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/http"
	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	"github.com/smooth-infra/smooth-infra/pkg/yaml"

	"path/filepath"

	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

type TestFunction func(t *testing.T, params map[string]interface{}, expects map[string]interface{}) error

var availableTests = map[string]TestFunction{
	"http/request": http.Request,
}

func RunTests(t *testing.T, yamlSource io.Reader) map[string]error {
	var filePath string
	if file, ok := yamlSource.(*os.File); ok {
		filePath = filepath.ToSlash(file.Name())
	}

	var directory string
	if filePath != "" {
		directory = path.Dir(filePath)
	}

	yamlContent, err := io.ReadAll(yamlSource)
	require.Nil(t, err)

	config, err := yaml.Process(string(yamlContent))
	require.Nil(t, err)

	// Currently only Terraform is supported
	if config.Input != nil {
		outputsFilePath := strings.TrimPrefix(filepath.ToSlash(config.Input[config.GetInputName()].OutputsFile), "//")
		outputPath := strings.TrimPrefix(fmt.Sprintf("%s/%s", directory, outputsFilePath), "/")
		outputValues, err := terraform.GetOutputValues(outputPath)
		require.Nil(t, err)

		config.ProcessTests(config.GetInputName(), outputValues)
	}

	errors := make(map[string]error)
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
