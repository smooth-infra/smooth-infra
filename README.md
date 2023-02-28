# smooth-infra

[![Go Report Card](https://goreportcard.com/badge/github.com/smooth-infra/smooth-infra)](https://goreportcard.com/report/github.com/smooth-infra/smooth-infra) [![codecov](https://codecov.io/gh/smooth-infra/smooth-infra/branch/main/graph/badge.svg?token=KVB6AVHPI5)](https://codecov.io/gh/smooth-infra/smooth-infra) [![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

## Usage

```golang
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

	errors := core.RunTests(t, file)
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

This project is extremely new and it is also my first open-source serious project. I am extremely excited to begin such a project as I find infrastructural testing to be really important to make sure that our infrastructure works as expected.

If you have any idea for any functionality, or if you find a bug in the code, please create an [Issue](https://github.com/smooth-infra/smooth-infra/issues/new) or a [Pull Request](https://github.com/smooth-infra/smooth-infra/compare) and I'll make sure to check it as soon as possible!

## License

This project uses [Mozilla Public License 2.0](/LICENSE).