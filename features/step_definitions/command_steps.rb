# frozen_string_literal: true

Given "I look for executables only in {string} within the current directory" do |directory|
  # this application also needs git in the path, until I can get my changes merged for go-git
  set_environment_variable "PATH", File.dirname(which('git')) + File::PATH_SEPARATOR

  prepend_environment_variable "PATH", expand_path(directory) + File::PATH_SEPARATOR
end

Then(/(the )?(\w+) from `([^`]*)` should (not )?contain "([^"]*)"/) do |_, channel, cmd, negated, content|
  run_command_and_validate_channel(
    cmd: cmd,
    fail_on_error: true,
    channel: channel,
    negated: negated,
    match_as_regex: false,
    content: content
  )
end

Then(/(the )?(\w+) from `([^`]*)` should (not )?match "([^"]*)"/) do |_, channel, cmd, negated, content|
  run_command_and_validate_channel(
    cmd: cmd,
    fail_on_error: true,
    channel: channel,
    negated: negated,
    match_as_regex: true,
    content: content
  )
end
