require 'fileutils'

Given('I have installed go-git-mob into {string} within the current directory') do |path|
  exe = File.join(aruba.root_directory, 'bin', 'git-mob')
  raise "'#{exe}' not found; did you run 'make build'?" unless File.exist?(exe)

  create_directory(path)
  FileUtils.cp(exe, File.join(aruba.current_directory, path))
  run_command_and_stop("git mob explode", fail_on_error: true)
  run_command_and_stop("ls -la #{path}", fail_on_error: true)
end
