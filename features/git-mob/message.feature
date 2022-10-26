Feature: message

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

  # @announce-stdout
  Scenario: starting update ~/.gitmessage but does not configure commit.template
    Given I cd to "example"
    When I run `git mob ad`
    Then the output should contain:
      """
      Jane Doe <jane@example.com>
      Amy Doe <amy@findmypast.com>
      """
    And the file "~/.gitmessage" should not contain "Co-authored-by:"
    And the file ".git/config" should not contain:
      """
      template = ~/.gitmessage
      """
    And the file "~/.gitconfig" should not contain:
      """
      template = ~/.gitmessage
      """

  # @announce-stdout @announce-stderr
  Scenario: git mob should warn about local commit.template
    Given a file named "example/.git/config" with:
      """
      [commit]
      template = ~/.commit-message-template
      """
    And I cd to "example"
    When I successfully run `git mob`
    Then the stdout from "git mob" should contain:
      """
      Warning: local commit.template value detected

      Using local commit.template could mean your template does not have selected co-authors appended after switching projects.

      If you do not use commit.template (e.g. it was added by an earlier version of go-git-mob) you can safely remove it:

          git config --local --unset commit.template

      If your team or project uses a local commit.template value you can silence this warning for this repo with:

          git config --local git-mob.use-local-template true

      Happy Mobbing!

      """

  Scenario: git mob should not warn about local commit.template when local commit.template is unset
    Given a file named "example/.git/config" with:
      """
      [commit]
      template = ~/.commit-message-template
      """
    And I cd to "example"
    And I successfully run `git config --local --unset commit.template`
    When I successfully run `git mob`
    Then the stdout from "git mob" should not contain:
      """
      Warning: local commit.template value detected
      """

  Scenario: git mob should not warn about local commit.template when use-local-template is set locally
    Given a file named "example/.git/config" with:
      """
      [commit]
      template = ~/.commit-message-template
      """
    And I cd to "example"
    And I successfully run `git config --local git-mob.use-local-template true`
    When I successfully run `git mob`
    Then the stdout from "git mob" should not contain:
      """
      Warning: local commit.template value detected
      """

  # TOOD: this is not yet supported
  # Scenario: git mob should not warn about local commit.template when use-local-template is set globally
  #   Given a file named "example/.git/config" with:
  #     """
  #     [commit]
  #     template = ~/.commit-message-template
  #     """
  #   And I cd to "example"
  #   And I successfully run `git config --global git-mob.use-local-template true`
  #   When I successfully run `git mob`
  #   Then the stdout from "git mob" should not contain:
  #     """
  #     Warning: local commit.template value detected
  #     """
