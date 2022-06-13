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
      co-author = Amy Doe <amy@findmypast.com>
    """
    When I run git mob `print`
    Then the output should contain:
    """
    Co-authored-by: Amy Doe <amy@findmypast.com>
    """

  # @announce-stdout @announce-stderr
  Scenario: two coauthors configured
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      co-author = Bob Doe <bob@findmypast.com>
    """
    When I run git mob `print`
    Then the output should contain:
    """
    Co-authored-by: Amy Doe <amy@findmypast.com>
    Co-authored-by: Bob Doe <bob@findmypast.com>
    """

  Scenario: two coauthors configured, list initials
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      co-author = Bob Doe <bob@findmypast.com>
    """
    And a file named "~/.git-coauthors" with:
    """
    {
      "coauthors": {
        "mr": {
          "name": "Amy Doe",
          "email": "amy@findmypast.com"
        },
        "zw": {
          "name": "Bob Doe",
          "email": "bob@findmypast.com"
        }
      }
    }
    """
    When I run git mob `print -i`
    Then the output should match:
    """
    mr zw
    """

Scenario: list initials with long flag
    Given a file named "~/.gitconfig" with:
    """
    [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      co-author = Bob Doe <bob@findmypast.com>
    """
    And a file named "~/.git-coauthors" with:
    """
    {
      "coauthors": {
        "mr": {
          "name": "Amy Doe",
          "email": "amy@findmypast.com"
        },
        "zw": {
          "name": "Bob Doe",
          "email": "bob@findmypast.com"
        }
      }
    }
    """
    When I run git mob `print --initials`
    Then the output should match:
    """
    mr zw
    """
