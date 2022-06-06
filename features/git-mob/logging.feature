# @wip @announce-stdout @announce-git-log
Feature: logging support

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com

      [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      """

    And a simple git repo at "example"

    Given I set the environment variables to:
      | variable         | value        |
      | GITMOB_LOG_LEVEL | debug        |
      | GITMOB_LOG_FILE  | .git/mob.log |

  @announce-stdout
  Scenario: commit with logging
    Given I cd to "example"
    And an empty file ".git/COMMIT_EDITMSG"
    When I successfully run `git mob prepare-commit-msg .git/COMMIT_EDITMSG message`
    And I successfully run `cat .git/mob.log`
    Then the file ".git/COMMIT_EDITMSG" should contain:
      """

      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """
    And the file ".git/mob.log" should contain:
      """
      logging initialized
      """
    And the file ".git/mob.log" should contain:
      """
      > COMMIT_SOURCE = message
      """