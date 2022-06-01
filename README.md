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
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
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
- [Usage](#usage)
  - [Utility Commands](#utility-commands)
- [Roadmap](#roadmap)
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

- TODO (we are still in development)

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

## Contributing

See the [CONTRIBUTING](CONTRIBUTING.md) guide for local development setup and contribution guidelines.

<!-- LICENSE -->
## License

Distributed under the GPU v3 License. See [LICENSE](LICENSE) for more information.

<!-- CONTACT -->
## Contact

David Alpert - [@davidalpert](https://twitter.com/davidalpert)

Project Link: [https://github.com/davidalpert/go-git-mob](https://github.com/davidalpert/go-git-mob)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Richard Kotze & Dennis Ideler](https://tech.findmypast.com/co-author-commits-with-git-mob/) for the [git-mob](https://github.com/rkotze/git-mob) nodejs implementation

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