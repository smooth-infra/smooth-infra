# smooth-infra

## Background

As a DevOps Engineer and Cloud Architect, I always find myself dealing with infrastructural testing. For some it is fun, for some it is not, but it's a necessary pain to make sure that what you're building works as expected. I came up with this project because I'm myself sometimes overwhelmed by the infrastructural tests complexity, without even counting how long it takes to write these tests.

Because of these issues, sometimes infrastructural tests are overlooked and skipped, which may cause problems in the long-run in your precious production environments. This is why I came up with **smooth-infra**.

## Why smooth-infra?

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

Here's a very basic example:
```yaml
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
```

## Contributing

This project is extremely new and it is also my first open-source serious project. I am extremely excited to begin such a project as I find infrastructural testing to be really important to make sure that our infrastructure works as expected.

If you have any idea for any functionality, or if you find a bug in the code, please create an [Issue](https://github.com/smooth-infra/smooth-infra/issues/new) or a [Pull Request](https://github.com/smooth-infra/smooth-infra/compare) and I'll make sure to check it as soon as possible!

## License

This project uses [Mozilla Public License 2.0](/LICENSE).