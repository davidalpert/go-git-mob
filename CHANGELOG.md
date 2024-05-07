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


<a name="v0.11.1"></a>
## [v0.11.1] - 2024-05-06
### Bug Fixes
- prepare-commit-msg script fails when cannot find git mob

### Build
- wip up the [@not](https://github.com/not)-windows as it's failing (er, passing)
- fix issue where installing godepgraph failed with invalid go version

### Pull Requests
- Merge pull request [#147](https://github.com/davidalpert/go-git-mob/issues/147) from davidalpert/[GH-141](https://github.com/davidalpert/go-git-mob/issues/141)-hook-scripts-fail-without-login-shell


<a name="v0.11.0"></a>
## [v0.11.0] - 2024-02-02
### Features
- suggest-coauthors accepts filters
- git mob -p prints location of co-authors file

### Build
- add preconditions for the install task
- fix typo in tools.go
- set custom sort order for commit groups
- read tool versions from lockfiles
- upgrade to go1.20
- upgrade goreleaser flags
- upgrade goreleaser syntax
- upgrade to go1.19
- add .tool-versions to support asdf-managed golang
- **deps:** bump ruby/setup-ruby from 1.150.0 to 1.171.0
- **deps:** bump github.com/onsi/gomega from 1.20.2 to 1.31.1
- **deps:** bump ruby/setup-ruby from 1.139.0 to 1.150.0
- **deps:** bump github.com/spf13/cobra from 1.6.1 to 1.7.0
- **deps:** bump actions/setup-go from 3 to 4

### Pull Requests
- Merge pull request [#132](https://github.com/davidalpert/go-git-mob/issues/132) from davidalpert/sort-release-notes-commit-groups
- Merge pull request [#131](https://github.com/davidalpert/go-git-mob/issues/131) from davidalpert/126-filter-suggested-coauthors
- Merge pull request [#129](https://github.com/davidalpert/go-git-mob/issues/129) from davidalpert/dependabot/go_modules/github.com/onsi/gomega-1.31.1
- Merge pull request [#128](https://github.com/davidalpert/go-git-mob/issues/128) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.171.0
- Merge pull request [#127](https://github.com/davidalpert/go-git-mob/issues/127) from davidalpert/125-print-coauthors-file-location
- Merge pull request [#130](https://github.com/davidalpert/go-git-mob/issues/130) from davidalpert/read-tool-versions-from-lockfiles
- Merge pull request [#118](https://github.com/davidalpert/go-git-mob/issues/118) from davidalpert/dependabot/go_modules/github.com/spf13/cobra-1.7.0
- Merge pull request [#115](https://github.com/davidalpert/go-git-mob/issues/115) from davidalpert/dependabot/github_actions/actions/setup-go-4
- Merge pull request [#120](https://github.com/davidalpert/go-git-mob/issues/120) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.150.0


<a name="v0.10.0"></a>
## [v0.10.0] - 2023-03-10
### Features
- better handling of anonymous github email addresses
- show a nice message when you know all the author suggestions
- coauthor suggestions now filters out known coauthors
- suggest coauthor initials as lowercase

### Test Coverage
- co-author suggestion feature uses long form

### Bug Fixes
- build the correct cmd/git-mob package on release
- goreleaser script had old make targets
- co-author suggestion format isn't copy/pasteable
- co-author suggestion uses the wrong command
- uninstall can break hook scripts test is non-deterministic

### Build
- replace makefile with taskfile
- **deps:** bump ruby/setup-ruby from 1.133.0 to 1.139.0
- **deps:** bump github.com/stretchr/testify from 1.8.1 to 1.8.2
- **deps:** bump ruby/setup-ruby from 1.126.0 to 1.133.0
- **deps:** bump goreleaser/goreleaser-action from 3 to 4
- **deps:** bump ruby/setup-ruby from 1.120.0 to 1.126.0
- **deps:** bump github.com/stretchr/testify from 1.8.0 to 1.8.1
- **deps:** bump github.com/spf13/cobra from 1.5.0 to 1.6.1
- **deps:** bump ruby/setup-ruby from 1.117.0 to 1.120.0

### Pull Requests
- Merge pull request [#113](https://github.com/davidalpert/go-git-mob/issues/113) from davidalpert/106-suggesting-coauthors-generates-output-which-is-not-copypastable
- Merge pull request [#112](https://github.com/davidalpert/go-git-mob/issues/112) from davidalpert/78-explore-task-and-taskfiles-as-an-alternative-to-makemakefiles
- Merge pull request [#109](https://github.com/davidalpert/go-git-mob/issues/109) from davidalpert/dependabot/go_modules/github.com/stretchr/testify-1.8.2
- Merge pull request [#111](https://github.com/davidalpert/go-git-mob/issues/111) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.139.0
- Merge pull request [#103](https://github.com/davidalpert/go-git-mob/issues/103) from davidalpert/dependabot/github_actions/goreleaser/goreleaser-action-4
- Merge pull request [#104](https://github.com/davidalpert/go-git-mob/issues/104) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.133.0
- Merge pull request [#101](https://github.com/davidalpert/go-git-mob/issues/101) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.126.0
- Merge pull request [#97](https://github.com/davidalpert/go-git-mob/issues/97) from davidalpert/dependabot/go_modules/github.com/spf13/cobra-1.6.1
- Merge pull request [#98](https://github.com/davidalpert/go-git-mob/issues/98) from davidalpert/dependabot/go_modules/github.com/stretchr/testify-1.8.1
- Merge pull request [#96](https://github.com/davidalpert/go-git-mob/issues/96) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.120.0


<a name="v0.9.2"></a>
## [v0.9.2] - 2022-10-25
### Test Coverage
- fix feature specs to align with [#89](https://github.com/davidalpert/go-git-mob/issues/89)
- example bug feature

### Features
- rename feature file
- added feature spec for expected git-mob behaviour

### Bug Fixes
- [#88](https://github.com/davidalpert/go-git-mob/issues/88) co-authors not cleared from message template when returning to git solo
- colons break square brackets, added them before instead
- lint, add period and colon consistency to README.md
- removed instances when .gitmessage is written to during a mob command
- colons break square brackets, added them before instead
- lint, add period and colon consistency to README.md

### Code Refactoring
- rename a feature file

### Build
- don't install goreleaser as [@latest](https://github.com/latest) requires go1.17
- doctor script exit code should reflect doctor status
- Features PR step could pass with errors
- reviewdog token needs write permissions on issues
- **deps:** bump ruby/setup-ruby from 1.115.3 to 1.117.0

### Pull Requests
- Merge pull request [#94](https://github.com/davidalpert/go-git-mob/issues/94) from davidalpert/88-bug-co-authors-not-cleared-from-message-template-when-returning-to-git-solo
- Merge pull request [#89](https://github.com/davidalpert/go-git-mob/issues/89) from teezzan/fix/co-authors_not_cleared_from_message_template_on_switch
- Merge pull request [#93](https://github.com/davidalpert/go-git-mob/issues/93) from HugeIRL/fix-readme
- Merge pull request [#91](https://github.com/davidalpert/go-git-mob/issues/91) from davidalpert/90-pr-build-is-passing-with-failing-steps
- Merge pull request [#87](https://github.com/davidalpert/go-git-mob/issues/87) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.117.0


<a name="v0.9.1"></a>
## [v0.9.1] - 2022-09-30
### Bug Fixes
- managing coauthors does not need to be in a working tree

### Pull Requests
- Merge pull request [#86](https://github.com/davidalpert/go-git-mob/issues/86) from davidalpert/bug-coauthor-subcommands-require-being-in-a-repo


<a name="v0.9.0"></a>
## [v0.9.0] - 2022-09-27
### Docs
- improve post-install docs
- add shell one-liner to verify that GOPATH bin is in your PATH

### Features
- git mob now works outside a working tree

### Build
- **deps:** bump ruby/setup-ruby from 1.114.0 to 1.115.3
- **deps:** bump github.com/onsi/gomega from 1.20.0 to 1.20.2

### Pull Requests
- Merge pull request [#85](https://github.com/davidalpert/go-git-mob/issues/85) from davidalpert/docs
- Merge pull request [#84](https://github.com/davidalpert/go-git-mob/issues/84) from davidalpert/81-go-install-doesnt-put-git-mob-in-my-path
- Merge pull request [#83](https://github.com/davidalpert/go-git-mob/issues/83) from davidalpert/82-remove-requirement-for-git-mob-to-be-inside-a-working-copy
- Merge pull request [#80](https://github.com/davidalpert/go-git-mob/issues/80) from davidalpert/dependabot/github_actions/ruby/setup-ruby-1.115.3
- Merge pull request [#79](https://github.com/davidalpert/go-git-mob/issues/79) from davidalpert/dependabot/go_modules/github.com/onsi/gomega-1.20.2


<a name="v0.8.0"></a>
## [v0.8.0] - 2022-08-23
### Test Coverage
- reproduce issue [#41](https://github.com/davidalpert/go-git-mob/issues/41)
- add support for an [@announce](https://github.com/announce)-gitmessage tag
- verify duplicate coauthor initials thows error

### Features
- write duplicates initials to stderr
- raise error on duplicate coauthors initials

### Bug Fixes
- [#41](https://github.com/davidalpert/go-git-mob/issues/41) coauthors accumulate in gitmessage

### Code Refactoring
- extract method to put replaceCoauthors under test
- sentinal error to custom error

### Chore
- remove [@wip](https://github.com/wip) tag and clean up whitespace

### Pull Requests
- Merge pull request [#77](https://github.com/davidalpert/go-git-mob/issues/77) from davidalpert/41-bug-gitgitmessage-accumulates-authors
- Merge pull request [#76](https://github.com/davidalpert/go-git-mob/issues/76) from davidalpert/45-add-validation-to-git-coauthors-file-to-preventalert-on-duplicate-keys


<a name="v0.7.3"></a>
## [v0.7.3] - 2022-08-22
### Test Coverage
- uninstall command can break hook scripts
- fix typo in announced paths
- remove {project}/bin from aruba command_search_paths

### Docs
- add how-to for showing mob in your shell prompt

### Code Refactoring
- replace explode/implode with rehash/uninstall

### Build
- ignore windows/arm64
- downgrade dependencies to go1.16

### Pull Requests
- Merge pull request [#75](https://github.com/davidalpert/go-git-mob/issues/75) from davidalpert/71-feat-add-more-documentation-to-feature-specs
- Merge pull request [#72](https://github.com/davidalpert/go-git-mob/issues/72) from davidalpert/70-support-install-with-earlier-golang-versions-eg-117-or-116


<a name="v0.7.2"></a>
## [v0.7.2] - 2022-08-16
### Test Coverage
- reproduce [#67](https://github.com/davidalpert/go-git-mob/issues/67) and [#68](https://github.com/davidalpert/go-git-mob/issues/68)

### Docs
- add issue templates

### Bug Fixes
- ignore missing config keys when getting the current mob
- expose better error detail from underlying git config calls
- return the actual exit code from utils.SilentRun

### Pull Requests
- Merge pull request [#69](https://github.com/davidalpert/go-git-mob/issues/69) from davidalpert/67-getallglobalgit-mobco-author-nonzero-exit-code-1-when-soloing


<a name="v0.7.1"></a>
## [v0.7.1] - 2022-08-15
### Docs
- fix broken license badges
- add troubleshooting guidance

### Test Coverage
- enable [@announce](https://github.com/announce)-gitmob-log

### Bug Fixes
- log better error detail when SilentRun fails
- print log destination to diagnostics.Log instead of STDOUT
- expose missing output when subprocess exit code is non-zero

### Build
- configure dependabot.yml
- **deps:** bump actions/setup-go from 2 to 3
- **deps:** bump github.com/onsi/gomega from 1.10.1 to 1.20.0
- **deps:** bump github.com/stretchr/testify from 1.7.1 to 1.8.0
- **deps:** bump github.com/spf13/cobra from 1.4.0 to 1.5.0
- **deps:** bump ruby/setup-ruby from 1.110.0 to 1.114.0
- **deps:** bump actions/checkout from 2 to 3
- **deps:** bump goreleaser/goreleaser-action from 2 to 3

### Pull Requests
- Merge pull request [#61](https://github.com/davidalpert/go-git-mob/issues/61) from davidalpert/dependabot/github_actions/actions/checkout-3
- Merge pull request [#60](https://github.com/davidalpert/go-git-mob/issues/60) from davidalpert/dependabot/github_actions/goreleaser/goreleaser-action-3


<a name="v0.7.0"></a>
## [v0.7.0] - 2022-08-14
### Features
- --override-author overrides git author for current and future commands
- edit-coauthor lets you edit an existing coauthor
- delete-coauthor removes an existing coauthor by initials
- add-coauthor adds a new coauthor to ~/.git-coauthors
- print added co-authors to the terminal

### Docs
- improve warning message around commit.template

### Bug Fixes
- error message reversed when adding a coauthor

### Code Refactoring
- remove commented code
- extract the map of shims to it's own file
- standardize feature specs

### Build
- improve release notes content

### Pull Requests
- Merge pull request [#58](https://github.com/davidalpert/go-git-mob/issues/58) from davidalpert/improve-release-notes
- Merge pull request [#57](https://github.com/davidalpert/go-git-mob/issues/57) from davidalpert/7-feature-overwrite-the-main-author
- Merge pull request [#56](https://github.com/davidalpert/go-git-mob/issues/56) from davidalpert/10-feature-edit-co-author
- Merge pull request [#55](https://github.com/davidalpert/go-git-mob/issues/55) from davidalpert/9-feature-delete-co-author
- Merge pull request [#54](https://github.com/davidalpert/go-git-mob/issues/54) from davidalpert/[GH-8](https://github.com/davidalpert/go-git-mob/issues/8)-add-co-author
- Merge pull request [#53](https://github.com/davidalpert/go-git-mob/issues/53) from davidalpert/[GH-47](https://github.com/davidalpert/go-git-mob/issues/47)


<a name="v0.6.1"></a>
## [v0.6.1] - 2022-08-12
### Test Coverage
- ensure that git mob with no args prints the mob
- ensure that coauthors file exists

### Bug Fixes
- commit templates are stored globally
- git mob with no args should print mob

### Code Refactoring
- format .git-coauthors file
- reorganize code to better match latest git-mob src

### Build
- don't remove branchName when workspace is dirty

### Chore
- go mod tidy with go1.18 found the missing dependencies

### Pull Requests
- Merge pull request [#52](https://github.com/davidalpert/go-git-mob/issues/52) from davidalpert/update-parity


<a name="v0.6.0"></a>
## [v0.6.0] - 2022-08-07
### Docs
- update readme to simplify install instructions

### Build
- add a git-mob cmd wrapper
- remove build-all target
- align formatting
- inject release notes commit summary into version details

### Pull Requests
- Merge pull request [#51](https://github.com/davidalpert/go-git-mob/issues/51) from davidalpert/rename-installed-cmd
- Merge pull request [#50](https://github.com/davidalpert/go-git-mob/issues/50) from davidalpert/commit-version


<a name="v0.5.1"></a>
## [v0.5.1] - 2022-08-06
### Build
- build only once


<a name="v0.5.0"></a>
## [v0.5.0] - 2022-08-06
### Features
- git mob init-all --dry-run
- git mob init-all [base-path]

### Docs
- clean up RELEASE_NOTES
- clean up language to vale standards
- update binary install instructions

### Code Refactoring
- update go-printers and streamline printing
- factor out go-printers

### Build
- ensure that version tags are available for PR builds
- use git-mob as the project name
- allow makefile to override VERSION
- make version detail dependent on source files
- update PR workflow to fail on valedation failures
- clean up diagnostic output when overriding the branch name
- print provided version as a string, not a byte array

### Pull Requests
- Merge pull request [#48](https://github.com/davidalpert/go-git-mob/issues/48) from davidalpert/[GH-43](https://github.com/davidalpert/go-git-mob/issues/43)-fix-packaged-version
- Merge pull request [#49](https://github.com/davidalpert/go-git-mob/issues/49) from davidalpert/refactor-to-use-go-printers
- Merge pull request [#40](https://github.com/davidalpert/go-git-mob/issues/40) from davidalpert/init-all
- Merge pull request [#38](https://github.com/davidalpert/go-git-mob/issues/38) from davidalpert/use-go-printers
- Merge pull request [#37](https://github.com/davidalpert/go-git-mob/issues/37) from davidalpert/errata


<a name="v0.4.1"></a>
## [v0.4.1] - 2022-07-17
### Docs
- update installation instructions

### Code Refactoring
- rename github action workflows

### Build
- don't need to use branchName when it's 'main'
- assume 'main' branch when rev-list bewteen origin/main and HEAD is 0
- ensure that local main tracks origin/main in a tag build
- dump job context (with event) when building a release
- move version_gen.go into .tools
- simplify release notes commit message
- add test target to document testing push events with act


<a name="v0.4.0"></a>
## [v0.4.0] - 2022-07-16
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

### Bug Fixes
- generate release notes before running goreleaser
- build before releasing to update the version number
- update url for the conform tool
- git mob outside a working tree should fail

### Code Refactoring
- calculate semantic version at generate, not runtime
- use git CLI instead of go-git for revParse commands
- show better errors when executing setMob
- move SilentRun into a new shell package
- import git step defs from my blog post sample

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

### Pull Requests
- Merge pull request [#36](https://github.com/davidalpert/go-git-mob/issues/36) from davidalpert/32-improve-install-story
- Merge pull request [#34](https://github.com/davidalpert/go-git-mob/issues/34) from davidalpert/32-improve-install-story
- Merge pull request [#35](https://github.com/davidalpert/go-git-mob/issues/35) from davidalpert/add-license-1
- Merge pull request [#33](https://github.com/davidalpert/go-git-mob/issues/33) from davidalpert/4-add-pull-request-validation-and-feedback


<a name="v0.3.0"></a>
## [v0.3.0] - 2022-06-14
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

### Code Refactoring
- don't need to load the primary user more than once
- reorganize subcommands
- standardize & improve usage text
- replace go-git with git CLI
- collect git paths into a revParse module
- remove duplication in explode.feature
- rearrange suggest subcommand as a coauthors subcommand

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

### Bug Fixes
- doctor would constantly regenerate chglog config

### Code Refactoring
- replicate git-authors specs from git-mob
- replace firefly names with git-mob examples
- move Author into the cfg package
- extract env helpers to an env package

### Build
- go mod tidy
- add depgraph target
- ignore doctor.sh fixes


<a name="v0.1.0"></a>
## v0.1.0 - 2022-05-31
### Docs
- release notes for v0.1.0
- add initial project README

### Features
- add basic CLI framework

### Build
- add a makefile to tie it all together
- add baseline tool config
- add initial doctor.sh script
- add vscode workspace


[Unreleased]: https://github.com/davidalpert/go-git-mob/compare/v0.11.1...HEAD
[v0.11.1]: https://github.com/davidalpert/go-git-mob/compare/v0.11.0...v0.11.1
[v0.11.0]: https://github.com/davidalpert/go-git-mob/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/davidalpert/go-git-mob/compare/v0.9.2...v0.10.0
[v0.9.2]: https://github.com/davidalpert/go-git-mob/compare/v0.9.1...v0.9.2
[v0.9.1]: https://github.com/davidalpert/go-git-mob/compare/v0.9.0...v0.9.1
[v0.9.0]: https://github.com/davidalpert/go-git-mob/compare/v0.8.0...v0.9.0
[v0.8.0]: https://github.com/davidalpert/go-git-mob/compare/v0.7.3...v0.8.0
[v0.7.3]: https://github.com/davidalpert/go-git-mob/compare/v0.7.2...v0.7.3
[v0.7.2]: https://github.com/davidalpert/go-git-mob/compare/v0.7.1...v0.7.2
[v0.7.1]: https://github.com/davidalpert/go-git-mob/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/davidalpert/go-git-mob/compare/v0.6.1...v0.7.0
[v0.6.1]: https://github.com/davidalpert/go-git-mob/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/davidalpert/go-git-mob/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/davidalpert/go-git-mob/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/davidalpert/go-git-mob/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/davidalpert/go-git-mob/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/davidalpert/go-git-mob/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/davidalpert/go-git-mob/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/davidalpert/go-git-mob/compare/v0.1.0...v0.2.0
[license-shield]: https://img.shields.io/badge/License-MIT-yellow.svg
[license-url]: https://opensource.org/licenses/MIT
