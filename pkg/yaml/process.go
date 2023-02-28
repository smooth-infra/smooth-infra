package yaml

import (
	"fmt"
	"regexp"

	"github.com/smooth-infra/smooth-infra/pkg/terraform"
	goyaml "gopkg.in/yaml.v3"
)

type BaseStructure struct {
	Version int
	Input   map[string]Input `yaml:"input"`
	Tests   []Test
}

type Input struct {
	OutputsFile string `yaml:"outputs_file"`
}

type Test struct {
	Name    string
	Type    string
	Params  map[string]interface{}
	Expects map[string]interface{}
}

func Process(yaml string) (*BaseStructure, error) {
	structure := &BaseStructure{}

	err := goyaml.Unmarshal([]byte(yaml), &structure)
	if err != nil {
		return nil, err
	}
	return structure, nil
}

func (b *BaseStructure) GetInputName() string {
	var inputName string
	for k := range b.Input {
		inputName = k
		break
	}
	return inputName
}

func (b *BaseStructure) ProcessTests(inputName string, outputs *terraform.Data) {
	for testKey, test := range b.Tests {
		test.Name = replaceStringWithOutputValues(test.Name, inputName, outputs)

		result := make(map[string]interface{})
		for kParam, vParam := range test.Params {
			switch v := vParam.(type) {
			case string:
				result[kParam] = replaceStringWithOutputValues(v, inputName, outputs)
			default:
				result[kParam] = vParam
			}
			test.Params = result
		}

		b.Tests[testKey] = test
	}
}

func replaceStringWithOutputValues(str string, inputName string, outputs *terraform.Data) string {
	re := regexp.MustCompile(fmt.Sprintf(`\${input.%s\.(\w+)}`, inputName))

	matches := re.FindAllStringSubmatch(str, -1)

	params := make(map[string]string)
	for _, match := range matches {
		key := match[1] // Extract the captured group
		params[key] = fmt.Sprint(outputs.Variables[match[1]].Value)
	}

	replacer := func(match string) string {
		// Extract the key from the match
		keyMatches := re.FindStringSubmatch(match)
		if len(keyMatches) != 2 {
			// Key not found, return the original string
			return match
		}
		key := keyMatches[1]

		// Look up the replacement value in the map
		value, ok := params[key]
		if !ok {
			// Key not found, return the original string
			return match
		}

		return value
	}

	return re.ReplaceAllStringFunc(str, replacer)
}
