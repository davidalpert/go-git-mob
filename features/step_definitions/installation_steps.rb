require 'fileutils'

Given('I have installed git-mob into {string} within the current directory') do |path|
  exe = File.join(aruba.root_directory, 'bin', 'git-mob')
  raise "'#{exe}' not found; did you run 'make build'?" unless File.exist?(exe)

  expanded_path = expand_path(path)
  installed_exe = File.join(expanded_path, "git-mob")

  create_directory(path)
  FileUtils.cp(exe, expanded_path)

  run_command_and_stop("#{installed_exe} explode", fail_on_error: true)
  run_command_and_stop("ls -la #{path}", fail_on_error: true)
end
