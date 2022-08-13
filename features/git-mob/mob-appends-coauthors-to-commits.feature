Feature: git mob appends coauthors to commits

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I successfully run `which sed`

    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    And a file named "~/.git-coauthors" with:
      """
      {
        "coauthors": {
          "ad": {
            "name": "Amy Doe",
            "email": "amy@findmypast.com"
          },
          "bd": {
            "name": "Bob Doe",
            "email": "bob@findmypast.com"
          }
        }
      }
      """
    And a simple git repo at "example"

  # @announce-git-log
  Scenario: append coauthor to a commit with the message flag
    Given I cd to "example"
    And I successfully run `git mob init`
    And I successfully run `git mob ad`

    When I successfully run `git commit --allow-empty -m "empty mobbed commit"`

    Then the most recent commit log should contain:
      """
      empty mobbed commit

      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """

  Scenario: add coauthor when rebasing
    Given I cd to "example"
    And I successfully run `git mob init`
    And I successfully run `git mob ad`
    And I successfully run `git commit --allow-empty -m "empty mobbed commit"`

    When I successfully run `git mob bd`
    And I prepare to edit the commit message with sed:
      """
      s/empty mobbed commit/ammended mobbed commit/
      """
    And I successfully run `git commit --allow-empty --amend`

    Then the most recent commit log should contain:
      """
      ammended mobbed commit

      Co-Authored-By: Amy Doe <amy@findmypast.com>
      Co-Authored-By: Bob Doe <bob@findmypast.com>
      """