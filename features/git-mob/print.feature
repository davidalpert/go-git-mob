Feature: print

  Scenario: no coauthors configured
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
    """
    When I run git mob `print`
    Then the output should not contain:
    """
    Co-authored-by:
    """

  # @announce-stdout @announce-stderr
  Scenario: one coauthor configured
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
        co-author = Mal Reynolds <mal@serenity.com>
    """
    When I run git mob `print`
    Then the output should contain:
    """
    Co-authored-by: Mal Reynolds <mal@serenity.com>
    """

  # @announce-stdout @announce-stderr
  Scenario: two coauthors configured
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
        co-author = Mal Reynolds <mal@serenity.com>
        co-author = Zoe Washburne <zoe@serenity.com>
    """
    When I run git mob `print`
    Then the output should contain:
    """
    Co-authored-by: Mal Reynolds <mal@serenity.com>
    Co-authored-by: Zoe Washburne <zoe@serenity.com>
    """
