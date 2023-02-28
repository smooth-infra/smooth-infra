package main

import (
	"os"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestInfra(t *testing.T) {
	file, err := os.Open("./../examples/http/request.yaml")
	require.Nil(t, err)

	defer file.Close()

	errors := core.Run(t, file)
	require.NotNil(t, errors)
}
