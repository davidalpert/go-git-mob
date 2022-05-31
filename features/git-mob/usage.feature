Feature: usage

  # @announce-stdout @announce-stderr
  Scenario: help
    When I run git mob `--help`
    Then the output should show usage