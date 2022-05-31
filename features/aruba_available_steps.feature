# Feature: Available Aruba Steps
#   # uncomment me to describe available aruba step phrases
#   # see step implementations: https://github.com/cucumber/aruba/blob/master/lib/aruba/cucumber
#   # see step documentation: https://relishapp.com/cucumber/aruba/v/0-11-0/docs/getting-started

#   @announce-stdout
#   Scenario: 'available aruba steps'
#     Given an executable named "bin/cli" with:
#       """
#       #!/bin/bash
#       git clone https://github.com/cucumber/aruba.git
#       cd aruba
#       grep -E "When|Given|Then" lib/aruba/cucumber/*.rb | awk -F ":" '{ $1 = ""; print $0}' |sort
#       """
#     When I run `bin/cli`
#     Then the exit status should be 0
