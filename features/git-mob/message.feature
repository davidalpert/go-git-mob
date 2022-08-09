Feature: message

    Background:
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
      And a simple git repo at "example"

      # @announce-stdout
      Scenario: start mobbing with one
        When I cd to "example"
        And I run git mob `ad`
        Then the output should contain:
    """
    Jane Doe <jane@example.com>
    Amy Doe <amy@findmypast.com>
    """
        And the file "~/.gitmessage" should exist
        And the file "~/.gitmessage" should contain:
    """

    Co-authored-by: Amy Doe <amy@findmypast.com>
    """


