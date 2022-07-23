Feature: check-author.spec

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    Given a simple git repo at "example"

  Scenario: does not print warning when config present
    Given I cd to "example"
    And I successfully run `git config --global user.name "Jane Doe"`
    And I successfully run `git config --global user.email "jane@example.com"`
    When I successfully run `git mob`
    Then the stdout from "git mob" should contain:
      """
      Jane Doe <jane@example.com>
      """

  Scenario: prints warning and missing config when one argument is missing
    Given I cd to "example"
    And I successfully run `git config --global user.name "Jane Doe"`
    And I run `git config --global --unset user.email`
    When I run `git mob`
    Then the exit status should be 1
    And the stderr from "git mob" should contain:
      """
      warning: Missing information for the primary author. Set with:

      $ git config --global user.email "jane@example.com"
      """

  Scenario: prints warning and missing config when both arguments are missing
    Given I cd to "example"
    And I run `git config --global --unset user.name`
    And I run `git config --global --unset user.email`
    When I run `git mob`
    Then the exit status should be 1
    And the stderr from "git mob" should contain:
      """
      warning: Missing information for the primary author. Set with:

      $ git config --global user.name "Jane Doe"
      $ git config --global user.email "jane@example.com"
      """