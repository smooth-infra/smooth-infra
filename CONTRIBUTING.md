# Contributing to smooth-infra

Thank you for your interest in contributing to smooth-infra!

We work hard to provide a high-quality and useful tool, and we greatly value
feedback and contributions from our community. Whether it's a bug report,
new feature, correction, or additional documentation, we welcome your issues
and pull requests. Please read through this document before submitting any
[issues] or [pull requests][pr] to ensure we have all the necessary information to
effectively respond to your bug report or contribution.

Jump To:

* [Bug Reports](#Bug-Reports)
* [Code Contributions](#Code-Contributions)

## How to contribute

*Before you send us a pull request, please be sure that:*

1. You're working from the latest source on the main branch.
2. You check existing open, and recently closed, pull requests to be sure 
   that someone else hasn't already addressed the problem.
3. You create an issue before working on a contribution that will take a 
   significant amount of your time.

*Creating a Pull Request*

1. Fork the repository.
2. In your fork, make your change in a branch that's based on this repo's main branch.
3. Commit the change to your fork, using a clear and descriptive commit message.
4. Create a pull request, answering any questions in the pull request form.

For contributions that will take a significant amount of time, open a new 
issue to pitch your idea before you get started. Explain the problem and 
describe the content you want to see added to the documentation. Let us know 
if you'll write it yourself or if you'd like us to help. We'll discuss your 
proposal with you and let you know whether we're likely to accept it.   

## Bug Reports

You can file bug reports against the tool on the [GitHub issues][issues] page.

If you are filing a report for a bug or regression in the SDK, it's extremely
helpful to provide as much information as possible when opening the original
issue. This helps us reproduce and investigate the possible bug without having
to wait for this extra information to be provided. Please read the following
guidelines prior to filing a bug report.

1. Search through existing [issues][] to ensure that your specific issue has
   not yet been reported. If it is a common issue, it is likely there is
   already a bug report for your problem.

2. Provide a minimal test case that reproduces your issue or any error
   information you related to your problem. We can provide feedback much
   more quickly if we know what operations you are calling in the SDK. If
   you cannot provide a full test case, provide as much code as you can
   to help us diagnose the problem. Any relevant information should be provided
   as well, like whether this is a persistent issue, or if it only occurs
   some of the time.

## Code Contributions

We are always happy to receive code and documentation contributions to the tool. 
Code contributions to the tool are done through [Pull Requests][pr]. The list below are guidelines to use when submitting pull requests. These are the 
same set of guidelines that the core contributors use when submitting changes, and we ask the same of all community contributions as well:

1. The tool is released under the [Apache 2.0 license][license]. Any code you submit
   will be released under that license.

2. If you would like to implement support for a significant feature that is not
   yet available in the tool, please talk to us beforehand to avoid any
   duplication of effort. 

3. Wherever possible, pull requests should contain tests as appropriate.
   Bugfixes should contain tests that exercise the corrected behavior (i.e., the
   test should fail without the bugfix and pass with it), and new features
   should be accompanied by tests exercising the feature.

4. Pull requests that contain failing tests will not be merged until the test
   failures are addressed. Pull requests that cause a significant drop in the
   SDK's test coverage percentage are unlikely to be merged until tests have
   been added.

### Testing

To run the tests locally,  use:

```
go test ./... -v -count=1
```

[issues]: https://github.com/smooth-infra/smooth-infra/issues
[pr]: https://github.com/smooth-infra/smooth-infra/pulls
[license]: https://github.com/smooth-infra/smooth-infra/blob/main/LICENSE
