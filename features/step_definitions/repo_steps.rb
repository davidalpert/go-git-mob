Given('a simple git repo at {string}') do |path|
  create_directory(path)
  cd(path) do
    run_command_and_stop('git config --global init.defaultBranch main', fail_on_error: true)
    run_command_and_stop(sanitize_text("git init ."), fail_on_error: true)
    run_command_and_stop('git commit --allow-empty -m "initial, empty root commit"', fail_on_error: true)
  end
end

