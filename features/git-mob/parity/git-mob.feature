Feature: git-mob.spec

  Background:
    Given I have installed go-git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory
    And I successfully run `git config --global user.name "Jane Doe"`
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

  Scenario: -h prints help
    Given I successfully run `git mob -h`
    Then the output should contain "Usage"
    And the output should contain "Flags"
    And the output should contain "Examples"

  #@announce-stderr
  @not-windows
  Scenario: --help prints help
    # --help is intercepted by the git plugin launcher which returns a 404 (help not found)
    When I run `git mob --help`
    Then the output should contain "No manual entry for git-mob"

  Scenario: -v prints version
    When I successfully run `git mob -v`
    Then the output should match /\d.\d.\d/

  Scenario: --version prints version
    When I successfully run `git mob --version`
    Then the output should match /\d.\d.\d/

  Scenario: --list prints a list of avaialable co-authors
    When I successfully run `git mob --list`
    Then the output should contain:
      """
      ad Amy Doe amy@example.com
      bd Bob Doe bob@example.com
      """

  Scenario: prints only primary author when there is no mob
    When I successfully run `git mob`
    Then the output should contain:
      """
      Jane Doe <jane@example.com>
      """

  Scenario: prints current mob
    Given I cd to "example"
    And I successfully run `git mob ad bd`
    When I successfully run `git mob`
    Then the output should contain:
      """
      Jane Doe <jane@example.com>
      Amy Doe <amy@example.com>
      Bob Doe <bob@example.com>
      """

  Scenario: sets mob when co-author initials found
    Given I cd to "example"
    When I successfully run `git mob ad`
    Then the output should contain:
      """
      Jane Doe <jane@example.com>
      Amy Doe <amy@example.com>
      """

  @pending
  # issue #7 https://github.com/davidalpert/go-git-mob/issues/7
  # NOTE: the -o flag is already in use for --output; will need
  #       to either pick a different shortcut for the override
  #       command or disable output formatting; I would prefer
  #       to use a different flag peraps even --o
  Scenario: sets mob and override coauthor
    Given I cd to "example"
    When I successfully run `git mob -o ad bd`
    Then the output should contain:
      """
      Amy Doe <amy@example.com>
      Bob Doe <bob@example.com>
      """

  Scenario: errors when co-author initials not found
    Given I cd to "example"
    When I run `git mob zz`
    Then the output should contain:
      """
      author with initials 'zz' not found
      """

  Scenario: overwrites old mob when setting a new mob
    Given a simple git repo at "example"
    And I cd to "example"
    And a file named ".git/.gitmessage" with:
      """
      A commit title

      A commit body that goes into more detail.
      """
    And I successfully run `git mob ad`
    When I successfully run `git mob bd`
    Then the output should contain:
      """
      Jane Doe <jane@example.com>
      Bob Doe <bob@example.com>
      """
    And a file named ".git/.gitmessage" should contain:
      """
      A commit title

      A commit body that goes into more detail.

      Co-authored-by: Bob Doe <bob@example.com>
      """

  Scenario: appends co-authors to an existing commit template
    Given a simple git repo at "example"
    And I cd to "example"
    And a file named ".git/.gitmessage" with:
      """
      A commit title

      A commit body that goes into more detail.
      """
    When I successfully run `git mob ad bd`
    And a file named ".git/.gitmessage" should contain:
      """
      A commit title

      A commit body that goes into more detail.

      Co-authored-by: Amy Doe <amy@example.com>
      Co-authored-by: Bob Doe <bob@example.com>
      """

  Scenario: appends co-authors to a new commit template
    Given a simple git repo at "example"
    And I cd to "example"
    And a file named ".git/.gitmessage" does not exist
    When I successfully run `git mob ad bd`
    And a file named ".git/.gitmessage" should contain:
      """

      Co-authored-by: Amy Doe <amy@example.com>
      Co-authored-by: Bob Doe <bob@example.com>
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

# @wip
# @announce-stdout
