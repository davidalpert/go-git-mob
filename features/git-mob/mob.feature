Feature: mob

  Background:
    Given a file named "~/.gitconfig" with:
    """
    [user]
      name = Hoban Washburne
      email = wash@serenity.com
    """
    Given a file named "~/.git-coauthors" with:
    """
    {
      "coauthors": {
        "mr": {
          "name": "Mal Reynolds",
          "email": "mal@serenity.com"
        },
        "zw": {
          "name": "Zoe Washburne",
          "email": "zoe@serenity.com"
        }
      }
    }
    """

  # @announce-stdout @announce-stderr
  Scenario: start mob with one coauthor
    When I run git mob `mr`
    Then the output should contain:
    """
    Hoban Washburne <wash@serenity.com>
    Mal Reynolds <mal@serenity.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Mal Reynolds <mal@serenity.com>
    """

  # @announce-stdout @announce-stderr
  Scenario: start mob with two coauthors
    When I run git mob `mr zw`
    Then the output should contain:
    """
    Hoban Washburne <wash@serenity.com>
    Mal Reynolds <mal@serenity.com>
    Zoe Washburne <zoe@serenity.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Mal Reynolds <mal@serenity.com>
    \tco-author = Zoe Washburne <zoe@serenity.com>
    """

  Scenario: change coauthors
    When I run git mob `mr`
    And I run git mob `zw`
    Then the output should contain:
    """
    Hoban Washburne <wash@serenity.com>
    Zoe Washburne <zoe@serenity.com>
    """
    And the file "~/.gitconfig" should contain:
    """
    [git-mob]
    \tco-author = Zoe Washburne <zoe@serenity.com>
    """
