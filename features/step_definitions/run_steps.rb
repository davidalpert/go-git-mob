# frozen_string_literal: true

Given('I clear the screen') do
  aruba.command_monitor.clear
end

When(/^I run git mob `(.*?)`(?: for up to ([\d.]+) seconds)?$/) do |cmd, secs|
  cmd = "#{git_mob_bin} #{sanitize_text(cmd)}"
  run_command_and_stop(cmd, fail_on_error: false, exit_timeout: secs && secs.to_f)
end

When(/^I successfully run git mob `(.*?)`(?: for up to ([\d.]+) seconds)?$/) do |cmd, secs|
  cmd = "#{git_mob_bin} #{sanitize_text(cmd)}"
  run_command_and_stop(cmd, fail_on_error: true, exit_timeout: secs && secs.to_f)
end

Then('the {channel} should show usage') do |_channel|
  # pending # Write code here that turns the phrase above into concrete actions
  all_output_includes('Usage:')
end

When('I sleep for {int} seconds') do |n_seconds|
  sleep(n_seconds)
end

When('I prepare to edit the commit message with sed:') do |sed_string|
  set_environment_variable 'EDITOR', "sed -i -e '#{sed_string}'" 
end

When('I prepare to edit the sequence of commit messages with sed:') do |sed_string|
  set_environment_variable 'GIT_SEQUENCE_EDITOR', "sed -i -e '#{sed_string}'" 
end
