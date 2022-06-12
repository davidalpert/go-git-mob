@pending
# issue #8 https://github.com/davidalpert/go-git-mob/issues/8
Feature: git-add-coauthor.spec

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

  Scenario: adds coauthor to coauthors file
    When I successfully run `git add-coauthor tb "Barry Butterworth" barry@butterworth.org`
    Then the file named "~/.git-coauthors" should include these coauthors:
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
        },
        "tb": {
          "name": "Barry Butterworth",
          "email": "barry@butterworth.org"
        }
      }
      """

  Scenario: does not add coauthor to coauthors file if email invalid
    When I run `git add-coauthor tb "Barry Butterworth" barry.org`
    Then the exit status should be 1
    And the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "tb": {
          "name": "Barry Butterworth",
          "email": "barry.org"
        }
      }
      """

  Scenario: does not add coauthor to coauthors file if wrong amount of parameters
    When I run `git add-coauthor tb "Barry Butterworth"`
    Then the exit status should be 1
    And the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "tb": {
          "name": "Barry Butterworth"
        }
      }
      """

  Scenario: does not add coauthor to coauthors file if key already exists
    When I run `git add-coauthor ea "Emily Anderson" "emily@anderson.org"`
    Then the exit status should be 1
    And the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "ea": {
          "name": "Emily Anderson",
          "email": "emily@anderson.org"
        }
      }
      """

  Scenario: -h prints help
    When I successfully run `git add-coauthor -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"
