# issue #10 https://github.com/davidalpert/go-git-mob/issues/10
Feature: git-edit-coauthor.spec

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.git-coauthors" with:
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
    And a simple git repo at "example"
    And I cd to "example"

  Scenario: edits coauthors name in coauthors file
    When I successfully run `git edit-coauthor ea --name="emily aldershot"`
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
          "name": "emily aldershot",
          "email": "ealderson@findmypast.com"
        }
      }
      """

  Scenario: edits coauthors email in coauthors file
    When I successfully run `git edit-coauthor ea --email="emily@aldershot.com"`
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
          "email": "emily@aldershot.com"
        }
      }
      """

  Scenario: edits coauthors name and email in coauthors file
    When I successfully run `git edit-coauthor ea --email="emily@aldershot.com" --name="emily aldershot"`
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
          "name": "emily aldershot",
          "email": "emily@aldershot.com"
        }
      }
      """

  Scenario: does not update a random key input
    When I run `git edit-coauthor ea --ship="serenity"`
    Then the exit status should be 1
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
        }
      }
      """
    And the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "ea": {
          "name": "Elliot Alderson",
          "email": "emily@aldershot.com",
          "ship": "serenity"
        }
      }
      """

  Scenario: does not update if author does not already exist'
    When I run `git edit-coauthor bb --name="barry butterworth"`
    Then the exit status should be 1
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
        }
      }
      """
    And the file named "~/.git-coauthors" should not include these coauthors:
      """
      {
        "bb": {
          "name": "barry butterworth"
        }
      }
      """

  Scenario: -h prints help
    When I successfully run `git edit-coauthor -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"