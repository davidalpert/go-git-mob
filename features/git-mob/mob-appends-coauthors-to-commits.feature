Feature: git mob appends coauthors to commits

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com

      [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      """
    And a simple git repo at "example"

  # @announce-git-log
  Scenario: append coauthor to a commit with the message flag
    Given I cd to "example"
    And I successfully run `git mob init`
    And I successfully run `git commit --allow-empty -m "empty mobbed commit"`
    Then the most recent commit log should contain:
      """
      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """
