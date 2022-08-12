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
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-git-mob">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-git-mob/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-git-mob/issues">Request Feature</a>
</p>

## Changelog


<a name="v0.6.1"></a>
## [v0.6.1] - 2022-08-12
### Bug Fixes
- commit templates are stored globally
- git mob with no args should print mob

### Build
- don't remove branchName when workspace is dirty

### Chore
- go mod tidy with go1.18 found the missing dependencies

### Code Refactoring
- format .git-coauthors file
- reorganize code to better match latest git-mob src

### Test Coverage
- ensure that git mob with no args prints the mob
- ensure that coauthors file exists

### Pull Requests
- Merge pull request [#52](https://github.com/davidalpert/go-git-mob/issues/52) from davidalpert/update-parity


<a name="v0.6.0"></a>
## [v0.6.0] - 2022-08-07
### Build
- add a git-mob cmd wrapper
- remove build-all target
- align formatting
- inject release notes commit summary into version details

### Docs
- update readme to simplify install instructions

### Pull Requests
- Merge pull request [#51](https://github.com/davidalpert/go-git-mob/issues/51) from davidalpert/rename-installed-cmd
- Merge pull request [#50](https://github.com/davidalpert/go-git-mob/issues/50) from davidalpert/commit-version


<a name="v0.5.1"></a>
## [v0.5.1] - 2022-08-06
### Build
- build only once


<a name="v0.5.0"></a>
## [v0.5.0] - 2022-08-06
### Build
- ensure that version tags are available for PR builds
- use git-mob as the project name
- allow makefile to override VERSION
- make version detail dependent on source files
- update PR workflow to fail on valedation failures
- clean up diagnostic output when overriding the branch name
- print provided version as a string, not a byte array

### Code Refactoring
- update go-printers and streamline printing
- factor out go-printers

### Docs
- clean up RELEASE_NOTES
- clean up language to vale standards
- update binary install instructions

### Features
- git mob init-all --dry-run
- git mob init-all [base-path]

### Pull Requests
- Merge pull request [#48](https://github.com/davidalpert/go-git-mob/issues/48) from davidalpert/[GH-43](https://github.com/davidalpert/go-git-mob/issues/43)-fix-packaged-version
- Merge pull request [#49](https://github.com/davidalpert/go-git-mob/issues/49) from davidalpert/refactor-to-use-go-printers
- Merge pull request [#40](https://github.com/davidalpert/go-git-mob/issues/40) from davidalpert/init-all
- Merge pull request [#38](https://github.com/davidalpert/go-git-mob/issues/38) from davidalpert/use-go-printers
- Merge pull request [#37](https://github.com/davidalpert/go-git-mob/issues/37) from davidalpert/errata


<a name="v0.4.1"></a>
## [v0.4.1] - 2022-07-17
### Build
- don't need to use branchName when it's 'main'
- assume 'main' branch when rev-list bewteen origin/main and HEAD is 0
- ensure that local main tracks origin/main in a tag build
- dump job context (with event) when building a release
- move version_gen.go into .tools
- simplify release notes commit message
- add test target to document testing push events with act

### Code Refactoring
- rename github action workflows

### Docs
- update installation instructions


<a name="v0.4.0"></a>
## [v0.4.0] - 2022-07-16
### Bug Fixes
- generate release notes before running goreleaser
- build before releasing to update the version number
- update url for the conform tool
- git mob outside a working tree should fail

### Build
- enable preview of release notes
- tune goreleaser workflow
- add goreleaser github action
- tune .goreleaser configuration
- goreleaser init
- commit generated version detail when tagging releases
- enhance version_gen.go to accept version params
- add PR workflow
- feature-flag vale validation
- format bundle install output
- show output when installing the bundle
- disable tput and colors in GITHUB_ACTIONS

### Chore
- go mod tidy

### Code Refactoring
- calculate semantic version at generate, not runtime
- use git CLI instead of go-git for revParse commands
- show better errors when executing setMob
- move SilentRun into a new shell package
- import git step defs from my blog post sample

### Docs
- release notes for v0.4.0
- update installation instructions
- update license to match original MIT

### Features
- update version command to print semver format
- allow FAST=1 releases

### Test Coverage
- fix up failing features
- set GIT_CEILING_DIRECTORIES before all features

### Pull Requests
- Merge pull request [#36](https://github.com/davidalpert/go-git-mob/issues/36) from davidalpert/32-improve-install-story
- Merge pull request [#34](https://github.com/davidalpert/go-git-mob/issues/34) from davidalpert/32-improve-install-story
- Merge pull request [#35](https://github.com/davidalpert/go-git-mob/issues/35) from davidalpert/add-license-1
- Merge pull request [#33](https://github.com/davidalpert/go-git-mob/issues/33) from davidalpert/4-add-pull-request-validation-and-feedback


<a name="v0.3.0"></a>
## [v0.3.0] - 2022-06-14
### Bug Fixes
- remove debugging output
- change initials-only to initials
- update path to doctor.sh
- app version is always 0.0.0
- given a simple repo with commits do not create an initial commit
- propagate helper args to the subcommands
- git mob overwrites template instead of appending coauthors
- re-add the 'install' alias
- remove the test .gitmessage path from the version command
- match subcommands on aliases as well as names
- start mob with coauthor who didn't exist would work

### Build
- ignore the actual deploy scripts
- build-all before deploying all :rollsafe:
- add a list-ignored target
- add an implode make target
- add explicit gen dependency to build targets
- fix the install target and add an uninstall target
- add valedation into the ci target
- add guard-process to auto-run valedation on md file change
- add valedation as a make target
- add vale config and a set of initial .styles
- add vale as a build requirement
- remove irrelevant scopes

### Code Refactoring
- don't need to load the primary user more than once
- reorganize subcommands
- standardize & improve usage text
- replace go-git with git CLI
- collect git paths into a revParse module
- remove duplication in explode.feature
- rearrange suggest subcommand as a coauthors subcommand

### Docs
- release notes for v0.3.0
- add reference for BDD via cucumber/aruba
- explain the install/uninstall story
- correct the vale target's help text
- be more assertive about imploding the binary itself
- update branch naming convention
- sketch out initial architectural guidance in README.md
- resolve vale warnings in README.md
- resolve vale warnings in CONTRIBUTING.md
- add link to github project v1.0
- improve setup and contributing guidance

### Features
- set .git/.gitmessage only when inside a working tree
- set .git/.gitmessage as the commit template
- write diagnostic logs to file
- append coauthor to a commit with the message flag
- initialize a local repo with a prepare-commit-msg script
- allow format "text" by the format printers
- add a prepare-commit-msg subcommand
- print commit SHA and a dirty flag (when built from a dirty repo)
- implode/uninstall
- explode/rehash
- suggest co-authors from git history
- list all coauthors

### Test Coverage
- add/delete/edit-coauthors.spec
- git-suggest-coauthors.spec
- check-author.spec shows warning if primary git author is not set
- git-mob.spec document we don't need to be in a working tree
- git-solo.spec ignores positional arguments
- git-solo.spec removes co-authors from commit template
- git-solo.spec sets the current mob to the primary author
- git-mob.spec document we don't need to be in a working tree
- refactor setup; ensure git mob runs inside an example git repo
- git-mob.spec appends co-authors to a new commit template
- appends co-authors to an existing commit template
- git-mob.spec overwrites old mob when setting a new mob
- git-mob.spec errors when co-author initials not found
- git-mob.spec sets mob and override coauthor
- add a [@pending](https://github.com/pending) tag
- git-mob.specs sets mob when co-author initials found
- refactor common .git-coauthor setup
- git-mob.specs prints current mob
- git-mob.specs prints only primary author when there is no mob
- git-mob.specs --list print a list of available co-authors
- --version prints version
- git-mob.specs -v prints version
- git-mob.specs --help is an error
- git-mob.spec -h prints help
- prove that amending commits adds coauthors
- ignore featues with an [@ignore](https://github.com/ignore) flag

### Pull Requests
- Merge pull request [#31](https://github.com/davidalpert/go-git-mob/issues/31) from HugeIRL/29-fix-initials-flag
- Merge pull request [#30](https://github.com/davidalpert/go-git-mob/issues/30) from HugeIRL/patch-1
- Merge pull request [#28](https://github.com/davidalpert/go-git-mob/issues/28) from davidalpert/parity-specs
- Merge pull request [#27](https://github.com/davidalpert/go-git-mob/issues/27) from davidalpert/24-feat-support-logging-to-a-file
- Merge pull request [#26](https://github.com/davidalpert/go-git-mob/issues/26) from davidalpert/23-feat-append-to-commit-message
- Merge pull request [#22](https://github.com/davidalpert/go-git-mob/issues/22) from davidalpert/21-refactor-go-git-to-git-cli
- Merge pull request [#20](https://github.com/davidalpert/go-git-mob/issues/20) from davidalpert/17-feature-implode
- Merge pull request [#19](https://github.com/davidalpert/go-git-mob/issues/19) from davidalpert/16-feature-exploderehash
- Merge pull request [#18](https://github.com/davidalpert/go-git-mob/issues/18) from davidalpert/12-document-architecture-and-intended-usage
- Merge pull request [#15](https://github.com/davidalpert/go-git-mob/issues/15) from davidalpert/[GH-5](https://github.com/davidalpert/go-git-mob/issues/5)-integrate-valedation
- Merge pull request [#14](https://github.com/davidalpert/go-git-mob/issues/14) from davidalpert/[GH-11](https://github.com/davidalpert/go-git-mob/issues/11)-suggest-co-authors
- Merge pull request [#13](https://github.com/davidalpert/go-git-mob/issues/13) from davidalpert/[GH-6](https://github.com/davidalpert/go-git-mob/issues/6)-list-all-coauthors
- Merge pull request [#3](https://github.com/davidalpert/go-git-mob/issues/3) from davidalpert/[GH-2](https://github.com/davidalpert/go-git-mob/issues/2)-enhance-contributor-guidelines


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

### Docs
- release notes for v0.2.0

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


[Unreleased]: https://github.com/davidalpert/go-git-mob/compare/v0.6.1...HEAD
[v0.6.1]: https://github.com/davidalpert/go-git-mob/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/davidalpert/go-git-mob/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/davidalpert/go-git-mob/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/davidalpert/go-git-mob/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/davidalpert/go-git-mob/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/davidalpert/go-git-mob/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/davidalpert/go-git-mob/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/davidalpert/go-git-mob/compare/v0.1.0...v0.2.0
