<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- vale Google.Acronyms = NO -->
[![License: MIT v3][license-shield]][license-url]
<!-- vale Google.Acronyms = YES -->

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
  - [Why port the nodejs version to Golang?](#why-port-the-nodejs-version-to-golang)
  - [What about mob.sh?](#what-about-mobsh)
  - [Built with](#built-with)
- [Getting started](#getting-started)
  - [Install](#install)
    - [`go install`](#go-install)
    - [Pre-compiled binaries](#pre-compiled-binaries)
  - [Verify your installation](#verify-your-installation)
  - [Post-install steps](#post-install-steps)
  - [Uninstall](#uninstall)
- [Usage](#usage)
  - [Sub-command help](#sub-command-help)
- [Troubleshooting](#troubleshooting)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

</details>

<!-- ABOUT THE PROJECT -->
## About the project

`git-mob` helps manage git co-authors when collaborating in real-time.

As the original authors of the node `git-mob` tool wrote:

<!-- vale off -->
> Documenting code collaboration
>
> When we work together, we should document that. It’s more than just giving credit to others, it’s also informing everyone about who was responsible for introducing changes to a codebase. Keeping a clear record of your work has lasting value - ever stumbled on an old commit only to be left with unanswered questions? In addition to explaining in your commits why a change was introduced, it also helps to document who implemented the change in case there needs to be a follow up.
<!-- vale on -->

### Why port the nodejs version to Golang?

People working with nodejs commonly use a version manager like [`nodenv`](https://github.com/nodenv/nodenv), [`nvm`](https://github.com/nvm-sh/nvm), or [`asdf`](https://asdf-vm.com/) to manage several versions of nodejs side-by-side.

These tools install global packages per node version which means you have to install the node `git-mob` plugin once per node version.

In contrast Golang offers the ability to build source code into single-file executables which truly install globally, independent of any versioning tools.

A Golang version of `git-mob` simplifies the install and update story making this plugin more manageable.

### What about mob.sh?

Like the original nodejs `git-mob` plugin, this golang port tool differs from and complements [`mob.sh`](https://mob.sh/) in two key ways:
- whereas `mob.sh` detects co-authors from commit messages alone, `git-mob` and `go-git-mob` understand that not all co-authors have their hands on the keyboard each session.
- whereas `mob.sh` squashes each feature branch into a single commit, `git-mob` and `go-git-mob` leave this decision up to you, complimenting your workflow by injecting conventional `Co-authored-by:` comments into each commit message through the use of a git `prepare-commit-message` hook.

### Built with

* [Golang 1.18](https://golang.org/)
* [go-releaser](https://goreleaser.com/)

<!-- GETTING STARTED -->
## Getting started

### Install

> :warning: The install process changed in `v0.6.0` If you have a version of `go-git-mob` older to `v0.6.0` you must first uninstall the current version. See [Uninstall](#uninstall) for instructions.

#### `go install`

With a working golang installation at version >= 1.18 you can install or update with:

```
go install github.com/davidalpert/go-git-mob/cmd/git-mob@latest
```

#### Pre-compiled binaries

Visit the [Releases](https://github.com/davidalpert/go-git-mob/releases) page to find binary packages pre-compiled for a variety of `GOOS` and `GOARCH` combinations:
1. Download an appropriate package for your `GOOS` and `GOARCH`;
1. Unzip it and put the binary in your path;

### Verify your installation

1. Confirm that git recognizes the `git-mob` plugin:
    ```
    git mob version
    ```

    With `git-mob` installed that command displays the plugin version:
    ```
    git-mob 0.5.1+f5536c2
    ```

### Post-install steps

1. Install helper plugins [once per machine]:
    ```
    git mob rehash
    ```

    `git-mob` ships as a single-file executable. The `rehash` sub-command generates simple shell scripts to make the following helper plugins available:
    ```
    git mob-print
    git mob-version
    git solo
    git suggest-coauthors
    ```

1. Initialize `prepare-commit-msg` hook script [once per repository]:

    ```
    git mob init
    ```

### Uninstall

- `git-mob` ships with an `uninstall` sub-command which cleans up and removes the top-level mob plugins and deletes itself:

    ```
    git mob uninstall
    ```

<!-- USAGE EXAMPLES -->
## Usage

- TODO; coming as the project nears v1.0

### Sub-command help

`git-mob` contains help for the various sub-commands:

```
git mob -h
```

> :warning: When requesting help make sure to use the short `-h` flag as `git` may intercept the full `--help` flag 

```
$ git mob -h
A git plugin to help manage git coauthors.

Examples:
   $ git mob jd                                      # Set John as co-authors
   $ git solo                                        # Return to working by yourself (i.e. unset all co-authors)
   $ git mob -l                                      # Show a list of all co-authors, John Doe should be there

Usage:
  git mob [flags]
  git mob [command]

Use "git-mob [command] -h" for more information about a command.
```

<!-- Troubleshooting -->
## Troubleshooting

If you run into trouble you can ask `go-git-mob` to write some diagnostics to a log file by setting the following environment variables:

| Variable          | Default   | Description                                                      |
| ----------------- | --------- | ---------------------------------------------------------------- |
| GITMOB_LOG_LEVEL  | `"fatal"` | `"fatal"`, `"error"`, `"warning"`, `"warn"`, `"info"`, `"debug"` |
| GITMOB_LOG_FORMAT | `"text"`  | `"text"` or `"json"`                                             |
| GITMOB_LOG_FILE   | `""`      | path to a log file; when empty logs go to STDOUT                 |

Dial up log levels to show more detail:
```
GITMOB_LOG_LEVEL=debug git commit -m "my log message"
```

Capture log messages to a file:
```
GITMOB_LOG_FILE=./mob.log GITMOB_LOG_LEVEL=debug git commit -m "my log message"
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

