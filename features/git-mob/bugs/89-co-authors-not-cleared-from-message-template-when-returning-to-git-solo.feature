Feature: 🐛 co-authors not cleared from message template when returning to git solo

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com
      """

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

  Scenario: 
    Given I successfully run `git mob ad`
    And the file "~/.gitmessage" should not contain "Co-authored-by"
    When I successfully run `git solo`
    And the file "~/.gitmessage" should not contain "Co-authored-by"
