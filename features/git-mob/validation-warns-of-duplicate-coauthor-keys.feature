Feature: Validation warns of duplicate coauthor keys

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    And I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`

  Scenario: list fails with duplicate intinitials
    Given a file named "~/.git-coauthors" with:
      """
      {
        "coauthors": {
          "ad": {
            "name": "Amy Doe",
            "email": "amy@example.com"
          },
          "ad": {
            "name": "Arnold Doe",
            "email": "arnold@example.com"
          }
        }
      }
      """
    When I run `git mob --list`
    Then the exit status should be 1
    And the stderr should contain:
      """
      ad Amy Doe amy@example.com
      ad Arnold Doe arnold@example.com

      duplicate coauthor initials found
      """
