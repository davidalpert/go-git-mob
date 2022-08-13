# frozen_string_literal: true

Given('I clear the screen') do
  aruba.command_monitor.clear
end

# Then(/(the )?(\w+) should (not )?show usage/) do |_, channel, negated|
Then(/the (\w+) should show usage/) do |channel|
  negated = false
  validate_channel(
    channel: channel,
    negated: negated,
    match_as_regex: false,
    content: "Usage:"
  )
end

When('I sleep for {int} seconds') do |n_seconds|
  sleep(n_seconds)
end
