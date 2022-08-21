Feature: uninstall

  `go-git-mob` ships as a single-file binary and implements all additional
  behaviors as subcommands (e.g. `git mob solo`)

  To better match the behavior of the nodejs `git-mob` tool, `go-git-mob` is
  able to generate shims in your path to expose some subcommands to git as
  additional plugins.

  The `uninstall` command is similar to the `npm implode` command in that it
  removes all it's shims and then removes itself.

  NOTE: because git mob doesn't know where you might have dependences on
  `git mob` in git hook scripts it doesn't touch those

  Background:
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`

    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I run `git mob install`
    And a file named "local_bin/git-mob" should exist
    And a file named "local_bin/git-mob-print" should exist
    And a file named "local_bin/git-mob-version" should exist
    And a file named "local_bin/git-solo" should exist
    And a file named "local_bin/git-suggest-coauthors" should exist

  Scenario: aliases
    When I run `git mob uninstall -h`
    Then the output should contain:
      """
      Aliases:
        uninstall, implode
      """

  Scenario: removes helper plugins and itself
    When I successfully run `git mob uninstall`
    Then a file named "local_bin/git-mob" should not exist
    And a file named "local_bin/git-mob-print" should not exist
    And a file named "local_bin/git-mob-version" should not exist
    And a file named "local_bin/git-solo" should not exist
    And a file named "local_bin/git-suggest-coauthors" should not exist

  Scenario: aliased to implode
    When I successfully run `git mob implode`
    Then a file named "local_bin/git-mob" should not exist
    And a file named "local_bin/git-mob-print" should not exist
    And a file named "local_bin/git-mob-version" should not exist
    And a file named "local_bin/git-solo" should not exist
    And a file named "local_bin/git-suggest-coauthors" should not exist
