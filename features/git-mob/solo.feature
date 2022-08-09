Feature: solo

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

  Scenario: clear out co-authors
    When I run git mob `solo`
    Then the output should contain:
    """
    Jane Doe <jane@example.com>
    """
    And the file "~/.gitconfig" should not contain:
    """
    co-author
    """
