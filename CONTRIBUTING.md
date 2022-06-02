<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![License: GPL v3][license-shield]][license-url]
<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<h1 align="center">go-git-mob</h1>

<p align="center">
  A golang port of the nodejs <a href="https://github.com/rkotze/git-mob">git-mob</a> tool,
  <code>go-git-mob</code> assists with managing co-authors when mob programming.
  <br />
  <a href="./README.md">README</a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  .
  <a href="./CONTRIBUTING.md"><strong>CONTRIBUTING</strong></a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-git-mob">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-git-mob/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-git-mob/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>

- [Review Existing Issues](#review-existing-issues)
  - [Up for grabs](#up-for-grabs)
- [Setup for Local Development](#setup-for-local-development)
  - [Install Prerequisites](#install-prerequisites)
  - [Get the code](#get-the-code)
  - [Visit the doctor](#visit-the-doctor)
  - [Run locally](#run-locally)
- [Useful Make targets](#useful-make-targets)
- [Development workflow](#development-workflow)
  - [Branching guidelines](#branching-guidelines)
  - [Commit message guidelines](#commit-message-guidelines)

</details>

Contributions are what make the open source community such an amazing place to learn, inspire, and create.  Any contributions you make are **greatly appreciated**.

Please review this contribution guide to streamline your experience.

## Review Existing Issues

This project's github [issues](https://github.com/davidalpert/go-git-mob/issues) are a great place to report bugs and discuss proposed changes or new features.

A quick discussion to coordinate a proposed change can someimes save hours of rework. 

Please review existing issues before creating a new one or submitting a pull request.

### Up for grabs

Our current focus is on reaching feature parity and releasing v1.0.

- please review our [v1.0 - feature parity](https://github.com/davidalpert/go-git-mob/projects/1) project board which tracks that scope

## Setup for Local Development

### Install Prerequisites

* [make](https://www.gnu.org/software/make/manual/html_node/index.html#Top) (often comes pre-installed or installed with other dev tooling)
* [golang 1.18](https://golang.org/doc/manage-install)
  * with a working go installation:
    ```
    go install golang.org/dl/go1.18@latest
    go1.18 download
    ```
  * open a terminal with `go1.18` as the linked `go` binary

* ruby (for running acceptance tests)

  * this project uses ruby and cucumber/aruba for integration tests; we include a `.ruby-version` file which specifies the specific version of ruby that is expected/supported
  * use a ruby version manager like [rbenv](https://github.com/rbenv) or [asdf](https://asdf-vm.com/); or
  * install directly from [ruby-lang.org](https://www.ruby-lang.org/en/documentation/installation/)

### Get the code

1. [Fork the repo on Github](https://github.com/davidalpert/go-git-mob/fork)

1. Clone your fork
   ```sh
   git clone https://github.com/your-github-name/go-git-mob.git
   ```

### Visit the doctor

This repo includes a `doctor.sh` script which helps validate that development dependencies are configured correctly.

1. Verify dependencies
    ```sh
    ./.tools/doctor.sh
    ```

This script attempts to fix basic issues (running `go get` or `bundle install`)

If `doctor.sh` reports an issue that it cannot resolve you may need to help it by taking action.

Please log any issues with the doctor script by [reporting a bug](https://github.com/davidalpert/go-git-mob/issues).

### Run locally

1. Build and run the tests
    ```sh
    make cit
    ```
1. Run from source
    ```sh
    go run main.go version
    go run main.go --help
    ```

## Useful Make targets

This repo includes a `Makefile` for help running common tasks.

Run `make` with no args to list the available targets:
```
$ make

 go-git-mob 0.0.0 - available targets:

build-all                      build for all platforms
build                          build for current platform
changelog                      Generate/update CHANGELOG.md
clean                          clean build output
deploy-local                   deploy binaries locally (for testing)
deploy                         deploy binaries
doctor                         run .tools/doctor.sh to sort out development dependencies
gen                            invoke go generate
install                        build and install locally into GOPATH
test                           run tests
version                        show current version
----------                     ------------------
release-major                  release major version
release-minor                  release minor version
release-patch                  release patch version

```

## Development workflow

We use a standard open-source fork/pull-request workflow:

1. [Fork the repo on Github](https://github.com/davidalpert/go-git-mob/fork)
1. Create your Feature Branch
   ```
   git checkout -b GH-123-amazing-feature
   ```
1. Commit your Changes
   ```
   git commit -m 'Add some AmazingFeature'
   ```
1. Make sure the code builds and all tests pass
   ```
   make cit
   ```
3. Push to the Branch
   ```
   git push origin GH-123-amazing-feature
   ```
4. Open a Pull Request

    https://github.com/davidalpert/go-git-mob/compare/GH-123-amazing-feature

### Branching guidelines

When working on a pull request to address or resolve a github issue, prefix the branch name with the github issue key.

For example, suppose you are working on an issue with an id of 23

Create a branch which starts with `GH-23` and an optional hyphanated description.

### Commit message guidelines

This project uses [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary) to generate [CHANGELOG](CHANGELOG.md).

The format of a conventional commit is:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Our currently list of supported type tags include
```yaml
  - "build"    # Changes that affect the build system or external dependencies
  - "ci"       # Changes to our CI configuration files and scripts 
  - "docs"     # Documentation only changes
  - "feat"     # A new feature
  - "fix"      # A bug fix
  - "perf"     # A code change that improves performance
  - "refactor" # A code change that neither fixes a bug nor adds a feature
  - "test"     # Adding missing tests or correcting existing tests
```

Prefixing your commits with one of these type tags and the description will automatically be included in the [CHANGELOG](CHANGELOG.md) for the next release.
