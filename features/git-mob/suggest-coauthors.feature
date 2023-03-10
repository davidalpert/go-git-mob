Feature: Suggest co-authors from commit history

  Background: 
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    And a file named "~/.git-coauthors" with:
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
    When I run `git suggest-coauthors`
    Then the output should contain:
      """
      The following authors from your coauthors file have contributed to this repository:

      - ad "Amy Doe" amy@findmypast.com
      - bd "Bob Doe" bob@findmypast.com

      Here are some suggestions for coauthors based on existing authors of this repository:

      git add-coauthor jd "Jane Doe" jane@example.com

      Paste any line above.
      """

  Scenario: suggest all co-authors as text
    Given I cd to "example"
    When I run `git suggest-coauthors --all`
    Then the output should contain:
      """
      Here are some suggestions for coauthors based on existing authors of this repository:

      git add-coauthor ad "Amy Doe" amy@findmypast.com
      git add-coauthor bd "Bob Doe" bob@findmypast.com
      git add-coauthor jd "Jane Doe" jane@example.com

      Paste any line above.
      """

  Scenario: suggest co-authors as a table
    Given I cd to "example"
    When I run `git suggest-coauthors -otable`
    Then the output should contain:
      """
      +----------+----------+------------------+
      | INITIALS |   NAME   |      EMAIL       |
      +----------+----------+------------------+
      | jd       | Jane Doe | jane@example.com |
      +----------+----------+------------------+
      suggested co-authors
      """

  Scenario: suggest all co-authors as a table
    Given I cd to "example"
    When I run `git suggest-coauthors --all -otable`
    Then the output should contain:
      """
      +----------+----------+--------------------+
      | INITIALS |   NAME   |       EMAIL        |
      +----------+----------+--------------------+
      | ad       | Amy Doe  | amy@findmypast.com |
      | bd       | Bob Doe  | bob@findmypast.com |
      | jd       | Jane Doe | jane@example.com   |
      +----------+----------+--------------------+
      suggested co-authors
      """

  Scenario: suggest co-authors as yaml
    Given I cd to "example"
    When I run `git suggest-coauthors -oyaml`
    Then the output should contain:
      """
      - initials: jd
        name: Jane Doe
        email: jane@example.com
      """

  Scenario: suggest all co-authors as yaml
    Given I cd to "example"
    When I run `git suggest-coauthors --all -oyaml`
    Then the output should contain:
      """
      - initials: ad
        name: Amy Doe
        email: amy@findmypast.com
      - initials: bd
        name: Bob Doe
        email: bob@findmypast.com
      - initials: jd
        name: Jane Doe
        email: jane@example.com
      """

  Scenario: suggest co-authors as json
    Given I cd to "example"
    When I run `git suggest-coauthors -ojson`
    Then the output should contain:
      """
      [
        {
          "initials": "jd",
          "name": "Jane Doe",
          "email": "jane@example.com"
        }
      ]
      """

  Scenario: suggest all co-authors as json
    Given I cd to "example"
    When I run `git suggest-coauthors --all -ojson`
    Then the output should contain:
      """
      [
        {
          "initials": "ad",
          "name": "Amy Doe",
          "email": "amy@findmypast.com"
        },
        {
          "initials": "bd",
          "name": "Bob Doe",
          "email": "bob@findmypast.com"
        },
        {
          "initials": "jd",
          "name": "Jane Doe",
          "email": "jane@example.com"
        }
      ]
      """
