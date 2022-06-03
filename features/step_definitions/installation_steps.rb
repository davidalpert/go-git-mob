require 'fileutils'

Given('I have installed go-git-mob into my path') do
  bin = File.join(aruba.root_directory, 'bin', 'git-mob')
  raise "'#{bin}' not found; did you run 'make build'?" unless File.exist?(bin)

  local_bin_folder = 'bin'
  create_directory(local_bin_folder)
  FileUtils.cp(bin, File.join(aruba.current_directory, local_bin_folder))
  prepend_environment_variable(
    "PATH",
    expand_path(local_bin_folder) + File::PATH_SEPARATOR
  )

  run_command_and_stop('which git-mob', fail_on_error: true)
end

Given('I have installed go-git-mob into {string} within the current directory') do |path|
  exe = File.join(aruba.root_directory, 'bin', 'git-mob')
  raise "'#{exe}' not found; did you run 'make build'?" unless File.exist?(exe)

  create_directory(path)
  FileUtils.cp(exe, File.join(aruba.current_directory, path))
  run_command_and_stop("ls -la #{path}", fail_on_error: true)
end
