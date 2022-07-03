# frozen_string_literal: true

Given('I set the local user config to {string} {string}') do |user_name, user_email|
  run_command_and_stop("git config --add user.name \"#{user_name}\"", fail_on_error: true)
  run_command_and_stop("git config --add user.email \"#{user_email}\"", fail_on_error: true)
end

Given('I set the global user config to {string} {string}') do |user_name, user_email|
  run_command_and_stop("git config --global --add user.name \"#{user_name}\"", fail_on_error: true)
  run_command_and_stop("git config --global --add user.email \"#{user_email}\"", fail_on_error: true)
end

Given('I set the repository user config to {string} {string}') do |user_name, user_email|
  repo_config_file = File.join(aruba.root_directory, '.git', 'config')
  run_command_and_stop("git config --file \"#{repo_config_file}\" --add user.name \"#{user_name}\"",
                       fail_on_error: true)
  run_command_and_stop("git config --file \"#{repo_config_file}\" --add user.email \"#{user_email}\"",
                       fail_on_error: true)
end

# GIT_CEILING_DIRECTORIES controls the behavior of searching for a .git directory.
# If you access directories that are slow to load (such as those on a tape drive,
# or across a slow network connection), you may want to have Git stop trying earlier
# than it might otherwise, especially if Git is invoked when building your shell prompt.
Given(/git does (not )?search outside the (test|current) (directory|folder)/) do |negated, _, _|
  ceiling_dir = File.dirname(expand_path('.'))

  aruba.announcer.announce(:paths, "ceiling_dir:       #{ceiling_dir}")

  if negated
    set_environment_variable('GIT_CEILING_DIRECTORIES', ceiling_dir) # aruba.root_directory)
  else
    old_value =
      aruba.environment.to_h.fetch('GIT_CEILING_DIRECTORIES', '')

    new_value = old_value.split(File::PATH_SEPARATOR).reject { |p| p == ceiling_dir }.join(File::PATH_SEPARATOR)

    delete_environment_variable('GIT_CEILING_DIRECTORIES')
    set_environment_variable('GIT_CEILING_DIRECTORIES', new_value)
  end
end

Then(/git should (not )?detect a repository/) do |negated|
  run_command_and_validate_channel(
    cmd: 'git rev-parse --show-toplevel',
    fail_on_error: false,
    channel: 'stderr',
    negated: !negated, # should detect (negated == false) means this content should not be detected
    match_as_regex: false,
    content: 'not a git repository (or any of the parent directories)'
  )
end

Then(/GIT_DIR should (not )?contain "([^"]*)"/) do |negated, content|
  run_command_and_validate_channel(
    cmd: 'git rev-parse --show-toplevel',
    fail_on_error: true,
    channel: 'stdout',
    negated: negated,
    match_as_regex: false,
    content: content
  )
end

Then(/GIT_DIR should (not )?match "([^"]*)"/) do |negated, content|
  run_command_and_validate_channel(
    cmd: 'git rev-parse --show-toplevel',
    fail_on_error: true,
    channel: 'stdout',
    negated: negated,
    match_as_regex: true,
    content: content
  )
end

Then(/git configuration should (not )?contain "([^"]*)"/) do |negated, content|
  run_command_and_validate_channel(
    cmd: 'git config --list --show-origin',
    fail_on_error: true,
    channel: 'stdout',
    negated: negated,
    match_as_regex: false,
    content: content
  )
end

Then(/git configuration should (not )?match "([^"]*)"/) do |negated, content|
  run_command_and_validate_channel(
    cmd: 'git config --list --show-origin',
    fail_on_error: true,
    channel: 'stdout',
    negated: negated,
    match_as_regex: true,
    content: content
  )
end

Then(/the project directory's git config file should (not )?contain "([^"]*)"/) do |negated, content|
  repo_config_file = File.join(aruba.root_directory, '.git', 'config')
  # puts "repo_config_file: #{repo_config_file}"
  repo_config = File.read(repo_config_file)
  # puts "repo_config:\n---\n#{repo_config}\n---\n"

  if negated
    expect(repo_config).not_to file_content_including(content.chomp)
  else
    expect(repo_config).to file_content_including(content.chomp)
  end
end

Given(/the repository git config file has (a|no) user section/) do |presence|
  repo_config_file = File.join(aruba.root_directory, '.git', 'config')
  repo_config = File.read(repo_config_file)
  case presence
  when 'a'
    expect(repo_config).to file_content_including('[user]')
  when 'no'
    expect(repo_config).not_to file_content_including('[user]')
  end
end
