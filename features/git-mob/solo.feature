Feature: solo

  Failing Scenarios:
  cucumber features/git-mob/message.feature:24 # Scenario: start mobbing with one
  cucumber features/git-mob/mob-appends-coauthors-to-commits.feature:28 # Scenario: append coauthor to a commit with the message flag
  cucumber features/git-mob/mob-appends-coauthors-to-commits.feature:42 # Scenario: add coauthor when rebasing
  cucumber features/git-mob/mob-init.feature:16 # Scenario: initialize git-mob inside a given repo
  cucumber features/git-mob/parity/git-mob.feature:105 # Scenario: overwrites old mob when setting a new mob
  cucumber features/git-mob/parity/git-mob.feature:130 # Scenario: appends co-authors to an existing commit template
  cucumber features/git-mob/parity/git-mob.feature:150 # Scenario: appends co-authors to a new commit template


  Background:
    Given I run `git config --global user.name "Jane Doe"`
    And I run `git config --global user.email "jane@example.com"`
    Given a file named "~/.git-coauthors" with:
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

  Scenario: clear out co-authors
    When I run git mob `solo`
    Then the output should contain:
    """
    Jane Doe <jane@example.com>
    """
    And the file "~/.gitconfig" should not contain:
    """
    co-author
    """
