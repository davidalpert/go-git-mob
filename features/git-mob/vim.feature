Feature: animate vim interactively

  @wip @announce-stdout
  Scenario: test-vim
    Given an empty file named "tmp.txt"
    And I successfully run `which vim`
    When I run `vim tmp.txt` interactively
    And I type "i10\e:x\r"
    And I sleep for 10 seconds
    Then the file "tmp.txt" should contain "10"

  @ignore @wip @announce-stdout
  Scenario: test-vim-empty
    Given an empty file named "tmp.txt"
    And I successfully run `which vim`
    When I run `vim tmp.txt` interactively
    And I type ":x\r"
    And I sleep for 2 seconds
    Then the file "tmp.txt" should contain ""