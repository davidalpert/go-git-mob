Feature: git-solo.spec

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I successfully run `git config --global user.name "Jane Doe"`
    And I successfully run `git config --global user.email "jane@example.com"`
    And a file named "~/.git-coauthors" with:
      """
      {
        "coauthors": {
          "ad": {
            "name": "Amy Doe",
            "email": "amy@example.com"
          },
          "bd": {
            "name": "Bob Doe",
            "email": "bob@example.com"
          }
        }
      }
      """
    Given a simple git repo at "example"

  Scenario: sets the current mob to the primary author
    Given I cd to "example"
    And I successfully run `git mob ad bd`
    When I successfully run `git solo`
    Then the stdout from "git solo" should contain:
      """
      Jane Doe <jane@example.com>
      """
    And the stdout from "git solo" should not contain:
      """
      Bob Doe <bob@example.com>
      """

# @wip
# @announce-stdout