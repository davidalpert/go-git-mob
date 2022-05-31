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
    [user]
      name = Pavan Kumar Sunkara
      email = pavan.sss1991@gmail.com

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

  Scenario: two coauthors configured, list initials
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
      co-author = Mal Reynolds <mal@serenity.com>
      co-author = Zoe Washburne <zoe@serenity.com>
    """
    And a file named "~/.git-coauthors" with:
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
    When I run git mob `print -i`
    Then the output should match:
    """
    mr zw
    """
