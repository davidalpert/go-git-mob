Feature: git-solo.spec

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given I successfully run `git config --global user.name "Jane Doe"`
    And I successfully run `git config --global user.email "jane@example.com"`
    And a file named "~/.git-coauthors" with:
      """
      {
        "coauthors": {
          "ad": {
            "name": "Amy Doe",
            "email": "amy@example.com"
          },
          "bd": {
            "name": "Bob Doe",
            "email": "bob@example.com"
          }
        }
      }
      """
    Given a simple git repo at "example"

  Scenario: sets the current mob to the primary author
    Given I cd to "example"
    And I successfully run `git mob ad bd`
    When I successfully run `git solo`
    Then the stdout from "git solo" should contain:
      """
      Jane Doe <jane@example.com>
      """
    And the stdout from "git solo" should not contain:
      """
      Bob Doe <bob@example.com>
      """

  Scenario: removes co-authors from commit template
    Given I cd to "example"
    And a file named ".git/.gitmessage" with:
      """
      A commit title

      A commit body that goes into more detail.
      """
    And I successfully run `git mob ad bd`
    When I successfully run `git solo`
    Then a file named ".git/.gitmessage" should contain:
      """
      A commit title

      A commit body that goes into more detail.
      """

  Scenario: ignores positional arguments
    Given I cd to "example"
    And I successfully run `git mob ad bd`
    When I successfully run `git solo yolo`
    Then the stdout from "git solo yolo" should contain:
      """
      Jane Doe <jane@example.com>
      """
    And the stdout from "git solo yolo" should not contain:
      """
      Bob Doe <bob@example.com>
      """

  @pending
  # allow running some mob commands outside a working tree
  # so that the mob follows across project folders
  Scenario: warns when used outside of a git repo
    # also hard to test not being in a working tree
    # because by default the aruba test folder is
    # still "inside" the core project's working tree
    When I run `git mob`
    Then the exit status should be 1
