# frozen_string_literal: true

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
