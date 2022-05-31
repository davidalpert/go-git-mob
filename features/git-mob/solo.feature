Feature: solo

  Background:
    Given a file named "~/.gitconfig" with:
    """
    [user]
      name = Hoban Washburne
      email = wash@serenity.com
    [git-mob]
      co-author = Mal Reynolds <mal@serenity.com>
      co-author = Zoe Washburne <zoe@serenity.com>
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

    @wip
      @announce-stdout
  Scenario: clear out co-authors
    When I run git mob `solo`
    Then the output should contain:
    """
    Hoban Washburne <wash@serenity.com>
    """
    And the file "~/.gitconfig" should not contain:
    """
    co-author
    """
