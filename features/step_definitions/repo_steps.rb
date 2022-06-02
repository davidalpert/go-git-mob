Given('a simple git repo at {string}') do |path|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text("git init ."), fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "initial, empty root commit"', fail_on_error: true)
  end
end

Given('a simple git repo at {string} with the following empty commits') do |path|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text("git init ."), fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "initial, empty root commit"', fail_on_error: true)

    run_command_and_stop('git config user.name "Amy Doe"', fail_on_error: true)
    run_command_and_stop('git config user.email "amy@findmypast.com"', fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "an empty commit from Amy"', fail_on_error: true)

    run_command_and_stop('git config user.name "Bob Doe"', fail_on_error: true)
    run_command_and_stop('git config user.email "bob@findmypast.com"', fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "an empty commit from Bob"', fail_on_error: true)

    run_command_and_stop('git config --remove-section user', fail_on_error: true)
  end
end

Given('a simple git repo at {string} with the following empty commits:') do |path, table|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text("git init ."), fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "initial, empty root commit"', fail_on_error: true)

    data = table.raw
    data.drop(1).reverse.each do |cols|
       run_command_and_stop("git config user.name \"#{cols[0]}\"", fail_on_error: true)
       run_command_and_stop("git config user.email \"#{cols[1]}\"", fail_on_error: true)
       run_command_and_stop("git commit --allow-empty -m \"#{cols[2]}\"", fail_on_error: true)
    end

    run_command_and_stop('git config --remove-section user', fail_on_error: true)
  end
end
