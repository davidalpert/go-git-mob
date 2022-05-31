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
  <a href="./README.md"><strong>README</strong></a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-git-mob">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-git-mob/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-git-mob/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>

- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Installation](#installation)
    - [using `go install`](#using-go-install)
    - [build from source](#build-from-source)
- [Usage](#usage)
  - [Utility Commands](#utility-commands)
- [Roadmap](#roadmap)
- [Local Development](#local-development)
  - [Prerequisites](#prerequisites)
  - [Make targets](#make-targets)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

</details>

<!-- ABOUT THE PROJECT -->
## About The Project

`go-git-mob` is a tool to help manage git coauthors. It is a golang port of the nodejs `git-mob` tool, and differs from/complimenents [`mob.sh`]() in two key ways:
- whereas `mob.sh` assumes that all co-authors can be detected from commit messages, `git-mob` and `go-git-mob` understand that not all co-authors have their hands on the keyboard each sesson
- whereas `mob.sh` assumes that each feature branch will be squashed into a single commit, `git-mob` and `go-git-mob` leave this decision up to you, complimenting your workflow by injecting conventional `Co-authored-by:` comments into each commit message through the use of a git `prepare-commit-message` hook.

As the original authors of the node `git-mob` tool wrote:

> Documenting code collaboration
>
> When we work together, we should document that. It’s more than just giving credit to others, it’s also informing everyone about who was responsible for introducing changes to a codebase. Keeping a clear record of your work has lasting value - ever stumbled on an old commit only to be left with unanswered questions? In addition to explaining in your commits why a change was introduced, it also helps to document who implemented the change in case there needs to be a follow up.

`go-git-mob` was created to provide a single-file binary version of the node `git-mob` tool which could be installed globally and operate independent of any specific nodejs version in use for a given git repo.

### Built With

* [Golang 1.18](https://golang.org/)
* [go-git](https://github.com/go-git/go-git)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

#### using `go install`

1. Clone the repo
   ```sh
   git install github.com/davidalpert/go-git-mob
   ```
1. Confirm the version
    ```sh
    git mob version
    ```

#### build from source

1. Clone the repo
   ```sh
   git clone https://github.com/davidalpert/go-git-mob.git
   ```
1. Install dependencies
    ```sh
    ./.tools/doctor.sh
    ```
1. Build and install locally in GOPATH
    ```sh
    make install
    ```
1. Confirm the version
    ```sh
    git mob version
    ```

<!-- USAGE EXAMPLES -->
## Usage

- TODO

### Utility Commands

The `git-mob` binary also ships with a number of utility commands which you can explore using the `--help` flag:

```
$: git mob -h
A tool for managing git coauthors.

Usage:
  git-mob [command]

Available Commands:
  help        Help about any command
  version     Show version information

Flags:
  -h, --help   help for git-mob

Use "git-mob [command] --help" for more information about a command.
```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/davidalpert/go-git-mob/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Local Development

### Prerequisites

`go-git-mob` is built and destributed as a single-file binary so there are no prerequisites.

* [golang](https://golang.org/doc/manage-install)
  * with a working go installation:
    ```
    go install golang.org/dl/go1.18@latest
    go1.18 download
    ```
* [make](https://www.gnu.org/software/make/manual/html_node/index.html#Top) (often comes pre-installed or installed with other dev tooling)

### Make targets

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
doctor                         run doctor.sh to sort out development dependencies
gen                            invoke go generate
install                        build and install locally into GOPATH
test                           run tests
version                        show current version
----------                     ------------------
release-major                  release major version
release-minor                  release minor version
release-patch                  release patch version

```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the GPU v3 License. See [LICENSE](LICENSE) for more information.

<!-- CONTACT -->
## Contact

David Alpert - [@davidalpert](https://twitter.com/davidalpert)

Project Link: [https://github.com/davidalpert/go-git-mob](https://github.com/davidalpert/go-git-mob)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Richard Kotze & Dennis Ideler](https://tech.findmypast.com/co-author-commits-with-git-mob/) and for the [git-mob](https://github.com/rkotze/git-mob) nodejs implementation

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/davidalpert/go-git-mob
[contributors-image-url]: https://contrib.rocks/image?repo=davidalpert/go-git-mob
[forks-shield]: https://img.shields.io/github/forks/davidalpert/go-git-mob
[forks-url]: https://github.com/davidalpert/go-git-mob/network/members
[issues-shield]: https://img.shields.io/github/issues/davidalpert/go-git-mob
[issues-url]: https://github.com/davidalpert/go-git-mob/issues
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0