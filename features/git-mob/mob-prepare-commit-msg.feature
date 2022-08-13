Feature: mob prepare-commit-msg

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

  Scenario: one coauthor - message - empty message
    Given an empty file ".git/COMMIT_EDITMSG"
    And I successfully run `git mob hooks prepare-commit-msg .git/COMMIT_EDITMSG message`
    Then the file ".git/COMMIT_EDITMSG" should contain:
      """

      Co-Authored-By: Amy Doe <amy@findmypast.com>
      """

  Scenario: one coauthor - message - message with comments
    Given a file named ".git/COMMIT_EDITMSG" with:
      """
      Add something awesome

      # Please enter the commit message for your changes. Lines starting
      # with '#' will be ignored, and an empty message aborts the commit.
      #
      # On branch 23-feat-append-to-commit-message
      # Your branch is up to date with 'origin/23-feat-append-to-commit-message'.
      #
      """
    And I successfully run `git mob hooks prepare-commit-msg .git/COMMIT_EDITMSG message`
    Then the file ".git/COMMIT_EDITMSG" should contain:
      """
      Add something awesome

      Co-Authored-By: Amy Doe <amy@findmypast.com>

      # Please enter the commit message for your changes. Lines starting
      # with '#' will be ignored, and an empty message aborts the commit.
      #
      # On branch 23-feat-append-to-commit-message
      # Your branch is up to date with 'origin/23-feat-append-to-commit-message'.
      #
      """
