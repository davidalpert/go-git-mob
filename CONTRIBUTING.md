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
<h1 align="center"><code>go-git-mob</code></h1>

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
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [Review existing issues](#review-existing-issues)
  - [Up for grabs](#up-for-grabs)
- [Setup for local development](#setup-for-local-development)
  - [Install prerequisites](#install-prerequisites)
  - [Get the code](#get-the-code)
  - [Visit the doctor](#visit-the-doctor)
  - [Run locally](#run-locally)
  - [Install from source](#install-from-source)
- [Useful Make targets](#useful-make-targets)
- [Development workflow](#development-workflow)
  - [Branch names](#branch-names)
  - [Commit message guidelines](#commit-message-guidelines)
- [Behavior-driven development with gherkin, cucumber, and aruba](#behavior-driven-development-with-gherkin-cucumber-and-aruba)
  - [Use git-mob's log file in a feature spec](#use-git-mobs-log-file-in-a-feature-spec)
  - [Custom step definitions](#custom-step-definitions)

</details>

Contributions make the open source community an great place to learn, inspire, and create.

Please review this contribution guide to streamline your experience.

## Review existing issues

Please review existing [issues](https://github.com/davidalpert/go-git-mob/issues) before reporting bug reports or requesting new features.

A quick discussion to coordinate a proposed change before you start can save hours of rework. 

### Up for grabs

The [v1.0 - feature parity](https://github.com/davidalpert/go-git-mob/projects/1) project board tracks progress through the remaining work in this current release.

## Setup for local development

### Install prerequisites

* [Make](https://www.gnu.org/software/make/manual/html_node/index.html#Top)  - often comes bundled with C compiler tools
* [golang 1.18](https://golang.org/doc/manage-install)
  * with a working go installation:
    ```
    go install golang.org/dl/go1.18@latest
    go1.18 download
    ```
  * open a terminal with `go1.18` as the linked `go` binary

* ruby 3.0.2

  * this project uses ruby and cucumber/aruba for integration tests and includes a `.ruby-version` file which specifies the supported/required version of ruby
  * use a ruby version manager like [rbenv](https://github.com/rbenv) or [asdf](https://asdf-vm.com/); or
  * install directly from [ruby-lang.org](https://www.ruby-lang.org/en/documentation/installation/)

### Get the code

1. [Fork the repository on Github](https://github.com/davidalpert/go-git-mob/fork)

1. Clone your fork
   ```sh
   git clone https://github.com/your-github-name/go-git-mob.git
   ```

### Visit the doctor

This repository includes a `doctor.sh` script which validates development dependencies.

1. Verify dependencies
    ```sh
    ./.tools/doctor.sh
    ```

This script attempts to fix basic issues, for example by running `go get` or `bundle install`.

If `doctor.sh` reports an issue that it can't resolve you may need to help it by taking action.

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

### Install from source

```sh
make install
```
This target builds the app and copies it into your GOPATH and makes it executable from any terminal which includes GOPATH in its path.

To install you can run either
```
make uninstall
```
or 
```
git mob implode
```

## Useful Make targets

This repository includes a `Makefile` for help running common tasks.

Run `make` with no arguments to list the available targets:
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

This project follows a standard open source fork/pull-request workflow:

1. First, [fork the repository on Github](https://github.com/davidalpert/go-git-mob/fork)


1. Create your Feature Branch
   ```
   git checkout -b 123-amazing-feature
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
   git push origin 123-amazing-feature
   ```
4. Open a Pull Request

    https://github.com/davidalpert/go-git-mob/compare/123-amazing-feature

### Branch names

When working on a pull request to address or resolve a Github issue, prefix the branch name with the Github issue number.

In the preceding example, after picking up an issue with an id of 123, create a branch which starts with `GH-123` or just `123-` and a hyphenated description:

```
git checkout -b 123-amazing-feature
```

### Commit message guidelines

This project uses [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary) to generate [CHANGELOG](CHANGELOG.md).

Format of a conventional commit:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

List of supported commit type tags include:
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

Prefix your commits with one of these type tags to automatically include the commit description in the [CHANGELOG](CHANGELOG.md) for the next release.

## Behavior-driven development with gherkin, cucumber, and aruba

This project includes integration tests written in [gherkin](https://cucumber.io/docs/gherkin/), a domain-specific language designed to specify given-when-then style specifications.

These specs live in the `./features/git-mob` folder:

    Scenario: start mob with one coauthor
      Given I cd to "example"
      When I run git mob `ad`
      Then the output should contain:
      """
      Jane Doe <jane@example.com>
      Amy Doe <amy@findmypast.com>
      """

This project uses the ruby [cucumber](https://cucumber.io/docs/installation/ruby/) implementation and the [cucumber/aruba](https://relishapp.com/cucumber/aruba/docs) gem which provides step definitions to manipulate files and run command-line applications.


### Use git-mob's log file in a feature spec

To enable `git-mob`'s debug logs set the following environment variables:

    Given I set the environment variables to:
      | variable         | value        |
      | GITMOB_LOG_LEVEL | debug        |
      | GITMOB_LOG_FILE  | .git/mob.log |

Show the log later with a step like:

    And I successfully run `cat .git/mob.log`

### Custom step definitions

This project includes several useful domain-specific step definitions, including:

    Then the most recent commit log should contain:
      """
      empty mobbed commit

      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """

That step uses `git log -1 --format=full` to read the formatted log message of
the previous git commit (at `HEAD~1`) and compare it to the provided string.

When feature specs using this step misbehave, decorate your features or scenarios with following custom [announcer](https://relishapp.com/cucumber/aruba/v/0-11-0/docs/core/announce-output-during-test-run) to print the full log message to STDOUT while running the tests:

    @announce-git-log

