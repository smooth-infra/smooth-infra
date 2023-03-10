# smooth-infra

[![Go Report Card](https://goreportcard.com/badge/github.com/smooth-infra/smooth-infra)](https://goreportcard.com/report/github.com/smooth-infra/smooth-infra) [![codecov](https://codecov.io/gh/smooth-infra/smooth-infra/branch/main/graph/badge.svg?token=KVB6AVHPI5)](https://codecov.io/gh/smooth-infra/smooth-infra) [![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

## Usage

After applying your Terraform code, let's verify that the endpoint provided by Terraform responds 200 OK.

First, let's create a configuration file called `test_terraform_endpoint.yml`

```yaml
version: 1
input:
  terraform:
    outputs_file: terraform_outputs.json
tests:
  - name: Verify that requesting ${input.terraform.address} is giving a 200 OK
    type: http/request
    params:
      address: ${input.terraform.address}
      secure: true
    expects:
      status_code: 200
```

Then let's execute the tests in a Go test file: 

```golang
package main

import (
	"os"
	"testing"

	"github.com/smooth-infra/smooth-infra/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestInfra(t *testing.T) {
	file, err := os.Open("test_terraform_endpoint.yml")
	require.Nil(t, err)

	defer file.Close()

	errors := core.Run(t, file)
	require.NotNil(t, errors)
}
```

## Background

You've just designed your perfect cloud infrastructure.

But is it really perfect?

There's a very interesting quote from Bruce Eckel:
> If it's not tested, it's broken.

And yes, there are many tools out there that we you can use to test your infrastructure, usually by using programming languages such as Golang, or frameworks like [Terratest](https://github.com/gruntwork-io/terratest).

But it's time consuming.

...and hard.

..and takes a lot of effort.

Furthermore, what if you don't know Golang?

**smooth-infra** makes all of this easy.

How?

By abstracting your tests in very simple human-readable YAML files, easy to scale, easy to write, easy to read.

## Why smooth-infra can make your life easier

As a DevOps Engineer and Cloud Architect, I always find myself dealing with infrastructural testing. For some it is fun, for some it is not, but it's a necessary pain to make sure that what you're building works as expected. I came up with this project because I'm myself sometimes overwhelmed by the infrastructural tests complexity, without even counting how long it takes to write these tests.

Because of these issues, sometimes infrastructural tests are overlooked and skipped, which may cause problems in the long-run in your precious production environments. This is why I came up with **smooth-infra**.

## Contributing

You can find more information in the [CONTRIBUTING.md](/CONTRIBUTING.md) file.

## License

This project uses [Apache-2.0 license](/LICENSE).