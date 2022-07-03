Feature: explode

  Background:
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    And I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

  Scenario: creates helper plugins
    When I run `git mob explode`
    Then a file named "local_bin/git-mob" should exist
    And a file named "local_bin/git-mob-print" should exist
    And a file named "local_bin/git-mob-version" should exist
    And a file named "local_bin/git-solo" should exist
    And a file named "local_bin/git-suggest-coauthors" should exist

  Scenario: helper-plugins work: mob-version
    When I run `git mob explode`
    And I run `git mob-version`
    Then the output should contain "git-mob"

  Scenario: helper-plugins work: suggest-coauthors
    Given a simple git repo at "example" with the following empty commits:
      | Name     | Email              | Commit_Message       |
      | Jane Doe | jane@example.com   | initial empty commit |
    When I run `git mob explode`
    And I cd to "example"
    And I run `git suggest-coauthors`
    Then the output should contain:
      """
      Here are some suggestions for coauthors based on existing authors of this repository:

      git mob add-coauthor JD Jane Doe jane@example.com
      """