Feature: rehash

  `go-git-mob` ships as a single-file binary and implements all additional
  behaviors as subcommands (e.g. `git mob solo`)

  To better match the behavior of the nodejs `git-mob` tool, `go-git-mob` is
  able to generate shims in your path to expose some subcommands to git as
  additional plugins.

  These shims (e.g. `git-solo.sh`) are simple shell scripts which forward
  calls to the appropriate mob subcommand allowing them to be called as
  their own git plugins (e.g. `git solo`)

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`

  Scenario: aliases
    When I run `git mob rehash -h`
    Then the output should contain:
      """
      Aliases:
        rehash, install, explode
      """

  Scenario: creates additional plugin shims
    When I run `git mob rehash`
    Then a file named "local_bin/git-mob" should exist
    And a file named "local_bin/git-mob-print" should exist
    And a file named "local_bin/git-mob-version" should exist
    And a file named "local_bin/git-solo" should exist
    And a file named "local_bin/git-suggest-coauthors" should exist

  Scenario: plugin shims wrap mob subcommands
    When I run `git mob rehash`
    Then the file "local_bin/git-solo" should contain:
      """
      #!/bin/sh
      git-mob solo "$@"
      """

  Scenario: plugin shims work: suggest-coauthors
    Given a simple git repo at "example" with the following empty commits:
      | Name     | Email            | Commit_Message       |
      | Jane Doe | jane@example.com | initial empty commit |
    When I run `git mob rehash`
    And I cd to "example"
    And I run `git suggest-coauthors`
    Then the output should contain:
      """
      Here are some suggestions for coauthors based on existing authors of this repository:

      git add-coauthor jd "Jane Doe" jane@example.com
      """