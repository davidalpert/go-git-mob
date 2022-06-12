@pending
# issue #9 https://github.com/davidalpert/go-git-mob/issues/9
Feature: git-delete-coauthor.spec

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And a file named "~/.git-coauthors" with:
      """
      {
        "coauthors": {
          "jd": {
            "name": "Jane Doe",
            "email": "jane@findmypast.com"
          },
          "fb": {
            "name": "Frances Bar",
            "email": "frances-bar@findmypast.com"
          },
          "ea": {
            "name": "Elliot Alderson",
            "email": "ealderson@findmypast.com"
          }
        }
      }
      """

  Scenario: deletes coauthor from coauthors file
    When I successfully run `git delete-coauthor ea`
    Then the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "ea": {
          "name": "Elliot Alderson",
          "email": "ealderson@findmypast.com"
        }
      }
      """

  Scenario: does nothing if initial are not a key in coauthors file
    When I run `git delete-coauthor bb`
    Then the exit status should be 1
    And the file named "~/.git-coauthors" should include these coauthors:
      """
      {
        "jd": {
          "name": "Jane Doe",
          "email": "jane@findmypast.com"
        },
        "fb": {
          "name": "Frances Bar",
          "email": "frances-bar@findmypast.com"
        },
        "ea": {
          "name": "Elliot Alderson",
          "email": "ealderson@findmypast.com"
        }
      }
      """

  Scenario: -h prints help
    When I successfully run `git delete-coauthor -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"
