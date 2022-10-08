Feature: mob-init

  Background:
    Given I have installed git-mob into "local_bin" within the current directory
    And I look for executables in "local_bin" within the current directory

    Given a file named "~/.gitconfig" with:
      """
      [user]
      name = Jane Doe
      email = jane@example.com

      [git-mob]
      co-author = Amy Doe <amy@findmypast.com>
      """
    Given a simple git repo at "example"

  Scenario: commit made with a coauthor should contain coauthor's details
    Given I cd to "example"
    And I run "git mod init"
    When I run `git mob ad`
    And I make changes to the content of repo "example"
    And I successfully run "git commit -m 'My commit message'"
    And I run "git log -1 --pretty=%B"
    Then the output should contain:
    """
    My commit message

    Co-Authored-By: Jane Doe <jane@example.com>
    """
    
  Scenario: commit made without a coauthor should contain no coauthor's details
    Given I cd to "example"
    And I run "git mod init"
    When I run `git mob solo`
    And I make changes to the content of repo "example"
    And I successfully run "git commit -m 'My solo commit message'"
    And I run "git log -1 --pretty=%B"
    Then the output should contain:
    """
    My solo commit message
    """
