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
<!-- vale Google.Headings = NO -->
<h1 align="center">go-git-mob</h1>
<!-- vale Google.Headings = YES -->

<p align="center">
  A golang port of the nodejs <a href="https://github.com/rkotze/git-mob">git-mob</a> tool,
  <code>go-git-mob</code> assists with managing co-authors when mob programming.
  <br />
  <a href="./README.md">README</a>
  ·
  <a href="./CHANGELOG.md"><strong>CHANGELOG</strong></a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-git-mob">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-git-mob/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-git-mob/issues">Request Feature</a>
</p>

## Changelog


<a name="v0.2.0"></a>
## [v0.2.0] - 2022-05-31
### Bug Fixes
- doctor would constantly regenerate chglog config

### Build
- go mod tidy
- add depgraph target
- ignore doctor.sh fixes

### Code Refactoring
- replicate git-authors specs from git-mob
- replace firefly names with git-mob examples
- move Author into the cfg package
- extract env helpers to an env package

### Features
- update .git/.gitmessage when the mob changes
- mob solo
- git mob
- git mob print -i
- mob print

### Test Coverage
- add step def to create a git repo
- add aruba/cucumber specs


<a name="v0.1.0"></a>
## v0.1.0 - 2022-05-31
### Build
- add a makefile to tie it all together
- add baseline tool config
- add initial doctor.sh script
- add vscode workspace

### Docs
- release notes for v0.1.0
- add initial project README

### Features
- add basic CLI framework


[Unreleased]: https://github.com/davidalpert/go-git-mob/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/davidalpert/go-git-mob/compare/v0.1.0...v0.2.0
