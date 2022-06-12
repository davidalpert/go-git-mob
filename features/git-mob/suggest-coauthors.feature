Feature: Suggest co-authors from commit history

  Background:
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    Given a file named "~/.git-coauthors" with:
    """
    {
      "coauthors": {
        "ad": {
          "name": "Amy Doe",
          "email": "amy@findmypast.com"
        },
        "bd": {
          "name": "Bob Doe",
          "email": "bob@findmypast.com"
        }
      }
    }
    """
    And a simple git repo at "example" with the following empty commits:
      | Name     | Email              | Commit_Message       |
      | Amy Doe  | amy@findmypast.com | Amy's empty commit   |
      | Bob Doe  | bob@findmypast.com | Bob's empty commit   |
      | Jane Doe | jane@example.com   | initial empty commit |

  Scenario: suggest co-authors as text
    Given I cd to "example"
    When I run git mob `coauthors suggest`
    Then the output should contain:
    """
    Here are some suggestions for coauthors based on existing authors of this repository:
    git mob add-coauthor AD Amy Doe amy@findmypast.com
    git mob add-coauthor BD Bob Doe bob@findmypast.com
    git mob add-coauthor JD Jane Doe jane@example.com
    """

  Scenario: suggest co-authors as a table
    Given I cd to "example"
    When I run git mob `coauthors suggest -otable`
    Then the output should contain:
    """
    +----------+----------+--------------------+
    | INITIALS |   NAME   |       EMAIL        |
    +----------+----------+--------------------+
    | AD       | Amy Doe  | amy@findmypast.com |
    | BD       | Bob Doe  | bob@findmypast.com |
    | JD       | Jane Doe | jane@example.com   |
    +----------+----------+--------------------+
    suggested co-authors
    """

  Scenario: suggest co-authors as yaml
    Given I cd to "example"
    When I run git mob `coauthors suggest -oyaml`
    Then the output should contain:
    """
    - initials: AD
      name: Amy Doe
      email: amy@findmypast.com
    - initials: BD
      name: Bob Doe
      email: bob@findmypast.com
    - initials: JD
      name: Jane Doe
      email: jane@example.com
    """

  Scenario: suggest co-authors as json
    Given I cd to "example"
    When I run git mob `coauthors suggest -ojson`
    Then the output should contain:
    """
    [
      {
        "initials": "AD",
        "name": "Amy Doe",
        "email": "amy@findmypast.com"
      },
      {
        "initials": "BD",
        "name": "Bob Doe",
        "email": "bob@findmypast.com"
      },
      {
        "initials": "JD",
        "name": "Jane Doe",
        "email": "jane@example.com"
      }
    ]
    """

