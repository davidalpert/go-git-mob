Feature: mob

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
    And a simple git repo at "example"

  Scenario: start mob with one coauthor
    Given I cd to "example"
    When I run `git mob ad`
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

  Scenario: start mob with a nonexistant coauthor
    Given I cd to "example"
    When I run `git mob be`
    Then the output should contain:
      """
      author with initials 'be' not found; run 'git mob --list' to see a list of available co-authors
      """

  Scenario: start mob with two coauthors
    Given I cd to "example"
    When I run `git mob ad bd`
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
    Given I cd to "example"
    When I run `git mob ad`
    And I run `git mob bd`
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

  Scenario: mob with no args prints the current mob
    Given I cd to "example"
    And I run `git mob bd`
    When I run `git mob`
    Then the stdout from "git mob" should contain:
      """
      Jane Doe <jane@example.com>
      Bob Doe <bob@findmypast.com>
      """

  Scenario: mob with no args outside a working tree prints the current mob
    Given I successfully run `git mob bd`
    When I successfully run `git mob`
    Then the stdout from "git mob" should contain:
      """
      Jane Doe <jane@example.com>
      Bob Doe <bob@findmypast.com>
      """

  Scenario: no git-coauthors file
    Given I remove the file "~/.git-coauthors"
    And I cd to "example"
    When I run `git mob ad`
    Then the output should contain:
      """
      author with initials 'ad' not found; run 'git mob --list' to see a list of available co-authors
      """
    And the file "~/.git-coauthors" should exist
    And the file "~/.git-coauthors" should contain:
      """
      {
      	"coauthors": {}
      }
      """
