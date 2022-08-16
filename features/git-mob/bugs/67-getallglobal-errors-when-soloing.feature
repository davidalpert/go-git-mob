#@announce-gitmob-log
#@announce-stdout @announce-stderr
Feature: 🐛 GetAllGlobal(git-mob.co-author): nonzero exit code: 1: (when soloing)

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com
      """
    And a simple git repo at "example"
    And I successfully run `git solo`

  Scenario: #67 git mob print
    Given I cd to "example"
    And I successfully run `git mob init`
    When I run `git mob print`
    Then the exit status should be 0

  Scenario: #68 git commit
    Given I cd to "example"
    And I successfully run `git mob init`
    When I run `git commit --allow-empty -m "test message"`
    Then the exit status should be 0