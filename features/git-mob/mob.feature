Feature: mob

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

  # @announce-stdout @announce-stderr
  Scenario: start mob with one coauthor
    When I run git mob `ad`
    Then the output should contain:
    """
    Jane Doe <jane@example.com>
    Amy Doe <amy@findmypast.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Amy Doe <amy@findmypast.com>
    """

  # @announce-stdout @announce-stderr
  Scenario: start mob with two coauthors
    When I run git mob `ad bd`
    Then the output should contain:
    """
    Jane Doe <jane@example.com>
    Amy Doe <amy@findmypast.com>
    Bob Doe <bob@findmypast.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Amy Doe <amy@findmypast.com>
    \tco-author = Bob Doe <bob@findmypast.com>
    """

  Scenario: change coauthors
    When I run git mob `ad`
    And I run git mob `bd`
    Then the output should contain:
    """
    Jane Doe <jane@example.com>
    Bob Doe <bob@findmypast.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Bob Doe <bob@findmypast.com>
    """
