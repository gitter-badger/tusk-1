# Contributing

All contributions are welcome and appreciated. Feel free to open issues or pull
requests for any fixes, changes, or new features, and if you are not sure about
anything, open it anyway. Issues and pull requests are a great forum for
discussion and a great opportunity to help improve the code as a whole.

## Opening Issues

A great way to contribute to the project is to open GitHub issues whenever you
encounter any issue, or if you have an idea on how Tusk can improve. Missing
or incorrect documentation are issues too, so feel free to open one whenever
you feel there is a chance to make Tusk better.

When reporting a bug, make sure to include the expected behavior and steps to 
reproduce. The more descriptive you can be, the faster the issue can be
resolved.

## Opening Pull Requests

Always feel free to open a pull request, whether it is a fix or a new addition.
For big or breaking changes, you might consider opening an issue first to check
interest, but it is absolutely not required to make a contribution.

Tests are run automatically on each PR, and 100% test and lint pass rate is
required to get the code merged in, although it is absolutely fine to have
work-in-progress pull requests open if you are trying to debug. details on how
to run the test suite can be found [here](#running-tests).

## Setting Up a Development Environment

For local development, you will need Go version 1.9+ installed.

To avoid issues with imports, Go projects must be placed on the 
[`GOPATH`][GOPATH], which defaults to `$HOME/go`:

```bash
git clone https://github.com/rliebz/tusk.git $(go env GOPATH)/src/github.com/rliebz/tusk
cd $(go env GOPATH)/src/github.com/rliebz/tusk
```

If it is not already on your path, you probably also want to have the `GOPATH`
binary directory available for projects installed by `go get` and `go install`.
To do so, add the following to your `.bash_profile` or `.zshrc`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Tusk currently uses [dep][dep] to manage its dependencies. To install
dependencies into the `vendor/` directory:

```bash
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

Finally, to install Tusk:

```bash
go install
```

If you have already installed tusk from another source, make sure you test
against the development version installed locally. If you do not get `dev` as
the version, you may need to move your Go binary path earlier in your `PATH`:

```bash
$ tusk --version
dev
```

Once `tusk` is on your path, make sure to run the `bootstrap` command to
install all other dependencies:

```bash
tusk bootstrap
```

## Making Changes

If you have not yet done so, make sure you fork the repository so you can push
your changes back to your own fork on GitHub.

When starting work on a new feature, create a new branch based off the `master`
branch. Pull requests should generally target the `master` branch, and releases
will be cut separately.

## Running Tests

To run the full test suite, along with `gometalinter`:

```bash
tusk test
```

If the gometalinter fails, execution will stop short and not actually run the
unit test suite. If there is a linter error that is a false-positive, or the 
violation is necessary for your contribution, you can disable a specific linter
for that line:

```golang
cmd := exec.Command("sh", "-c", command) // nolint: gas
```

For a slightly higher-fidelity test suite, you can [configure the `circleci`
command line tool][circleci-cli] along with docker to run your tests locally in
containers the same way they would be run against a pull request.

With docker running locally:

```bash
tusk circleci
```

[circleci-cli]: https://circleci.com/docs/2.0/local-jobs/#installing-the-cli-locally
[dep]: https://github.com/golang/dep
[GOPATH]: https://golang.org/doc/code.html#GOPATH
