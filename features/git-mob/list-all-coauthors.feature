Feature: List all co-authors

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
            "email": "amy@example.com"
          },
          "bd": {
            "name": "Bob Doe",
            "email": "bob@example.com"
          }
        }
      }
      """

  Scenario: list co-authors as text
    When I run `git mob --list`
    Then the output should contain:
      """
      ad Amy Doe amy@example.com
      bd Bob Doe bob@example.com
      """

  Scenario: list co-authors as table
    When I run `git mob --list -otable`
    Then the output should contain:
      """
      +----------+---------+-----------------+
      | INITIALS |  NAME   |      EMAIL      |
      +----------+---------+-----------------+
      | ad       | Amy Doe | amy@example.com |
      | bd       | Bob Doe | bob@example.com |
      +----------+---------+-----------------+
      """

  Scenario: list co-authors as yaml
    When I run `git mob --list -oyaml`
    Then the output should contain:
      """
      - initials: ad
        name: Amy Doe
        email: amy@example.com
      - initials: bd
        name: Bob Doe
        email: bob@example.com
      """

  Scenario: list co-authors as json
    When I run `git mob --list -ojson`
    Then the output should contain:
      """
      [
        {
          "initials": "ad",
          "name": "Amy Doe",
          "email": "amy@example.com"
        },
        {
          "initials": "bd",
          "name": "Bob Doe",
          "email": "bob@example.com"
        }
      ]
      """
