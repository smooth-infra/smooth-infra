package core

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	"github.com/smooth-infra/smooth-infra/pkg/yaml"

	"path/filepath"

	"github.com/stretchr/testify/require"
)

func Run(t *testing.T, yamlSource io.Reader) map[string]error {
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

	errors := ExecuteTests(t, config)

	return errors
}
