Feature: usage
  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

  # @announce-stdout @announce-stderr
  Scenario: help
    When I run `git mob -h`
    Then the stdout should show usage