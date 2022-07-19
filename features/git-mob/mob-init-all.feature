Feature: mob-init-all

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
    And a simple git repo at "example2"
    And a simple git repo at "example3"

  Scenario: initialize git-mob in each subfolder of the provided folder
    When I successfully run `git mob init-all .`
    Then the file "example/.git/hooks/prepare-commit-msg" should exist
    And the file "example/.git/hooks/prepare-commit-msg" should contain:
      """
      git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
      """
    And the file "example2/.git/hooks/prepare-commit-msg" should exist
    And the file "example2/.git/hooks/prepare-commit-msg" should contain:
      """
      git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
      """
    And the file "example3/.git/hooks/prepare-commit-msg" should exist
    And the file "example3/.git/hooks/prepare-commit-msg" should contain:
      """
      git mob hooks prepare-commit-msg "$COMMIT_MSG_FILE" $COMMIT_SOURCE $SHA1
      """
    And the output should contain:
      """
      initialized prepare-commit-msg git hooks in:
      - './example/'
      - './example2/'
      - './example3/'
      git-mob will now append coauthor annotations to commit messages in those repos

      the following folders were not initialized:
      - './local_bin/' ('./local_bin' does not appear to be a valid git repo)

      happy mobbing!
      """

