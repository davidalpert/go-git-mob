Feature: git-mob.spec

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I successfully run `git config --global user.name "Jane Doe"`
    And I successfully run `git config --global user.email "jane@example.com"`

  Scenario: -h prints help
    When I successfully run `git mob -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"
    And the output should contain "Examples"

