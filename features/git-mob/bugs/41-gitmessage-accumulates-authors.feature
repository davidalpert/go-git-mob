Feature: 🐛 gitmessage accumulates authors over time

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com

      [git-mob]
      use-local-template = false
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
    And a simple git repo at "example"

  # @announce-gitmessage
  Scenario: starting update ~/.gitmessage but does not configure commit.template
    Given I cd to "example"
    And I run `git mob ad`
    And the file "~/.gitmessage" should contain:
      """

      Co-authored-by: Amy Doe <amy@findmypast.com>
      """
    When I successfully run `git mob bd`
    Then the file "~/.gitmessage" should contain:
      """

      Co-authored-by: Bob Doe <bob@findmypast.com>
      """
    And the file "~/.gitmessage" should not contain:
      """

      Co-authored-by: Amy Doe <amy@findmypast.com>
      """
