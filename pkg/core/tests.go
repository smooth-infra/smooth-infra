package core

import (
	"testing"

	"github.com/gosimple/slug"
	"github.com/smooth-infra/smooth-infra/pkg/http"
	"github.com/smooth-infra/smooth-infra/pkg/null"
	"github.com/smooth-infra/smooth-infra/pkg/yaml"
)

type TestFunction func(t *testing.T, params map[string]interface{}, expects map[string]interface{}) error

var availableTests = map[string]TestFunction{
	"null/null":    null.Null,
	"http/request": http.Request,
}

func ExecuteTests(t *testing.T, config *yaml.BaseStructure) map[string]error {
	errors := make(map[string]error)

	for _, test := range config.Tests {
		if function, ok := availableTests[test.Type]; ok {
			t.Logf("Running \"%s\" (%s) test...", test.Name, test.Type)
			err := function(t, test.Params, test.Expects)
			if err != nil {
				if test.Id != "" {
					errors[test.Id] = err
				} else {
					slug := slug.Make(test.Name)
					errors[slug] = err
				}
			}
		}
	}

	return errors
}
