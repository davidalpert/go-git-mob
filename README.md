<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![License: MIT v3][license-shield]][license-url]

<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<!-- vale Google.Headings = NO -->
<h1 align="center">go-git-mob</h1>
<!-- vale Google.Headings = YES -->

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
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [About the project](#about-the-project)
  - [Built with](#built-with)
- [Getting started](#getting-started)
  - [Installation](#installation)
    - [Install a binary release](#install-a-binary-release)
    - [Install using go install](#install-using-go-install)
  - [Post-install steps](#post-install-steps)
  - [Uninstall](#uninstall)
- [Usage](#usage)
  - [Utility commands](#utility-commands)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

</details>

<!-- ABOUT THE PROJECT -->
## About the project

`go-git-mob` helps manage git co-authors when collaborating in real-time.

Like the original nodejs `git-mob` plugin, this golang port tool differs from and complements [`mob.sh`](https://mob.sh/) in two key ways:
- whereas `mob.sh` detects co-authors from commit messages alone, `git-mob` and `go-git-mob` understand that not all co-authors have their hands on the keyboard each session.
- whereas `mob.sh` squashes each feature branch into a single commit, `git-mob` and `go-git-mob` leave this decision up to you, complimenting your workflow by injecting conventional `Co-authored-by:` comments into each commit message through the use of a git `prepare-commit-message` hook.

As the original authors of the node `git-mob` tool wrote:

<!-- vale off -->
> Documenting code collaboration
>
> When we work together, we should document that. It’s more than just giving credit to others, it’s also informing everyone about who was responsible for introducing changes to a codebase. Keeping a clear record of your work has lasting value - ever stumbled on an old commit only to be left with unanswered questions? In addition to explaining in your commits why a change was introduced, it also helps to document who implemented the change in case there needs to be a follow up.
<!-- vale on -->

`go-git-mob` provides a single-file binary version of the node `git-mob` tool which installs globally and operates independent of any specific nodejs version in use for a given git repository.

### Built with

* [Golang 1.18](https://golang.org/)
* [go-git](https://github.com/go-git/go-git)

<!-- GETTING STARTED -->
## Getting started

To get a local copy up and running follow these simple steps.

### Installation

#### Install a binary release

- Download an appropriate package for your `GOOS` and `GOARCH` from the [Releases](https://github.com/davidalpert/go-git-mob/releases) tab;
- Uncompress it and put the binary in your path;
- Optionally, review [Post-install steps](#post-install-steps) to explode the convenience methods;

For example:

```
export VERSION=0.4.0
mkdir -p ~/bin/go-git-mob/v${VERSION}
mv ~/Downloads/go-git-mob_${VERSION}_Darwin_arm64.tar.gz ~/bin/go-git-mob/v${VERSION}/
cd ~/bin/go-git-mob/v${VERSION}
gunzip -c go-git-mob_${VERSION}_Darwin_arm64.tar.gz | tar xopf -
ln -f -s ~/bin/go-git-mob/v${VERSION}/go-git-mob ~/bin/git-mob
git mob version
```

#### Install using go install

With a working golang setup at version >= 1.18 you can use `go install`:

```
go install github.com/davidalpert/go-git-mob@latest
```

Unfortunately, due to a naming conflict with the original `git-mob` repo, this one installs as `$($GOBIN || $GOPATH/bin/)/go-git-mob` and goes unrecognized as a git plugin.

To enable it you can symlink `git-mob` to it like this:

```
ln -s "$(which go-git-mob)" "$(dirname $(which go-git-mob))/git-mob"
```

### Post-install steps

As a single file executable `go-git-mob` ships with one git plugin:
```
git mob
```
which offers several sub-commands:
```
git mob coauthors suggest
git mob print
git mob solo
git mob version
```

For easier use and compatibility with `git-mob`, `go-git-mob` includes an `install` sub-command:
```
git mob install
```
 which sets up the following top-level git plugins as aliases to the matching `git mob` sub-commands:
```
git mob-print
git mob-version
git solo
git suggest-coauthors
```

### Uninstall

`go-git-mob` ships with an `implode` sub-command:

```
git mob implode
```

which cleans up and removes the top-level mob plugins and deletes itself.

<!-- USAGE EXAMPLES -->
## Usage

- TODO; coming as the project nears v1.0

### Utility commands

The `git-mob` binary also ships with several utility commands which you can explore using the `--help` flag:

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

<!-- vale Google.Parens = NO -->
See [open issues](https://github.com/davidalpert/go-git-mob/issues) and specifically the [v1.0 - feature parity](https://github.com/davidalpert/go-git-mob/projects/1) project board for a list of known issues and up-for-grabs tasks.
<!-- vale Google.Parens = YES -->

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
[license-shield]: https://img.shields.io/badge/License-MIT-yellow.svg
[license-url]: https://opensource.org/licenses/MIT

