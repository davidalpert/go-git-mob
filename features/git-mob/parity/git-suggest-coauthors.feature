Feature: git-suggest-coauthors.spec

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given I successfully run `git config --global user.name "Jane Doe"`
    And I successfully run `git config --global user.email "jane@example.com"`

  Scenario: suggests potential coauthors
    Given a simple git repo at "example" with the following empty commits:
      | Name    | Email              | Commit_Message     |
      | Amy Doe | amy@findmypast.com | Amy's empty commit |
      | Bob Doe | bob@findmypast.com | Bob's empty commit |
    When I cd to "example"
    And I successfully run `git suggest-coauthors`
    Then the stdout from "git suggest-coauthors" should contain:
      """
      Here are some suggestions for coauthors based on existing authors of this repository:

      git add-coauthor AD "Amy Doe" amy@findmypast.com
      git add-coauthor BD "Bob Doe" bob@findmypast.com

      Paste any line above.
      """

  Scenario: shows error when there are no commits found
    Given a simple git repo at "example" with the following empty commits:
      | Name | Email | Commit_Message |
    When I cd to "example"
    And I run `git suggest-coauthors`
    Then the exit status should be 1
    Then the stderr from "git suggest-coauthors" should contain:
      """
      unable to find existing authors
      """

  Scenario: -h prints help
    Given a simple git repo at "example"
    When I cd to "example"
    And I successfully run `git suggest-coauthors -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"
