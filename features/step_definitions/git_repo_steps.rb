# frozen_string_literal: true

Given('a simple git repo at {string}') do |path|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text('git init .'), fail_on_error: true)

    # ensure that we have a valid local user for our empty commit
    run_command_and_stop('git config --add user.name "Mal Reynolds"', fail_on_error: true)
    run_command_and_stop('git config --add user.email "mal@serenity.com"', fail_on_error: true)

    # make a first commit to the repo because git functions more
    # predictably when a repo has at least one commit
    run_command_and_stop('git commit --allow-empty -m "initial, empty root commit"', fail_on_error: true)

    run_command_and_stop('git config --unset user.name', fail_on_error: true)
    run_command_and_stop('git config --unset user.email', fail_on_error: true)
  end
end

Given('a simple git repo at {string} with the following empty commits:') do |path, table|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text('git init .'), fail_on_error: true)

    data = table.raw.drop(1) # first row is headings
    data.reverse.each do |cols|
      run_command_and_stop("git config user.name \"#{cols[0]}\"", fail_on_error: true)
      run_command_and_stop("git config user.email \"#{cols[1]}\"", fail_on_error: true)
      run_command_and_stop("git commit --allow-empty -m \"#{cols[2]}\"", fail_on_error: true)
      run_command_and_stop('git config --remove-section user', fail_on_error: true)
    end
  end
end
