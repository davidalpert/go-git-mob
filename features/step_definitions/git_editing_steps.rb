# frozen_string_literal: true

When('I prepare to edit the commit message with sed:') do |sed_string|
  set_environment_variable 'EDITOR', "sed -i -e '#{sed_string}'"
end

When('I prepare to edit the sequence of commit messages with sed:') do |sed_string|
  set_environment_variable 'GIT_SEQUENCE_EDITOR', "sed -i -e '#{sed_string}'"
end
