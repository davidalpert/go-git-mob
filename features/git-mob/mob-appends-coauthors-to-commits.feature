Feature: git mob appends coauthors to commits

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

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
    And I successfully run `git commit --allow-empty -m "empty mobbed commit"`
    Then the most recent commit log should contain:
      """
      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """

  @announce-git-log @wip @announce-stdout
  Scenario: add coauthor when rebasing
    Given I set the environment variables to:
      | variable         | value        |
      | GITMOB_LOG_LEVEL | debug        |
      | GITMOB_LOG_FILE  | .git/mob.log |

    Given I cd to "example"
    And I successfully run `git mob init`
    And I successfully run `git mob ad`
    And I successfully run `git commit --allow-empty -m "empty mobbed commit"`
    And I successfully run `cat .git/mob.log`
    And I successfully run `git mob bd`
    And the most recent commit log should contain:
      """
      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """

    And I successfully run `git commit --allow-empty --amend -m "another empty commit"`
    And I successfully run `cat .git/mob.log`
    Then the most recent commit log should contain:
      """
      Co-Authored-By: Amy Doe <amy@findmypast.com>
      Co-Authored-By: Bob Doe <bob@findmypast.com>
      """