Feature: implode

  Background:
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    And I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I run `git mob explode`
    And a file named "local_bin/git-mob" should exist
    And a file named "local_bin/git-mob-print" should exist
    And a file named "local_bin/git-mob-version" should exist
    And a file named "local_bin/git-solo" should exist
    And a file named "local_bin/git-suggest-coauthors" should exist

  Scenario: removes helper plugins and itself
    When I run `git mob implode`
    Then a file named "local_bin/git-mob-print" should not exist
    And a file named "local_bin/git-mob-version" should not exist
    And a file named "local_bin/git-solo" should not exist
    And a file named "local_bin/git-suggest-coauthors" should not exist
    And a file named "local_bin/git-mob" should not exist

  Scenario: aliased to uninstall
    When I run `git mob uninstall`
    Then a file named "local_bin/git-mob-print" should not exist
    And a file named "local_bin/git-mob-version" should not exist
    And a file named "local_bin/git-solo" should not exist
    And a file named "local_bin/git-suggest-coauthors" should not exist
    And a file named "local_bin/git-mob" should not exist