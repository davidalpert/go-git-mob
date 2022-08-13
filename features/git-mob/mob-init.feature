Feature: mob-init

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com

      [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      """

  Scenario: initialize git-mob inside a given repo
    Given a simple git repo at "example"
    When I cd to "example"
    And I successfully run `git mob init`
    Then the file ".git/hooks/prepare-commit-msg" should exist
    And the file ".git/hooks/prepare-commit-msg" should contain:
      """
      #!/bin/sh

      COMMIT_MSG_FILE=$1
      COMMIT_SOURCE=$2
      SHA1=$3

      set -e

      git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
      """
    And the output should contain:
      """
      initialized local git hook: '.git/hooks/prepare-commit-msg'
      git-mob will now help prepare commit messages in this repo
      """
