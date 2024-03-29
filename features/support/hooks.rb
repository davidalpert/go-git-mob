require_relative 'os_helper'

Before('@announce-paths') do
  aruba.announcer.activate :paths
end

Before('@announce-gitmob-log') do
  aruba.announcer.activate :gitmob_log
end

Before('@announce-gitmessage') do
  aruba.announcer.activate :gitmessage
end

Before('@windows-only') do
  pending unless GitMob::OS.windows?
end

Before('@not-windows') do
  pending if GitMob::OS.windows?
end

Before('@pending') do
  pending
end

Before do
  aruba.announcer.announce(:paths, "aruba - root_directory:    #{aruba.root_directory}")
  aruba.announcer.announce(:paths, "aruba - working_directory: #{File.join(aruba.root_directory, aruba.config.working_directory)}")
  aruba.announcer.announce(:paths, "aruba - current_directory: #{File.join(aruba.root_directory, aruba.current_directory)}")

  # ensure that git, when run by aruba, does not accidentally discover
  # any of the parent repo's git config
  ceiling_dir = File.dirname(expand_path('.'))
  set_environment_variable('GIT_CEILING_DIRECTORIES', ceiling_dir)
  aruba.announcer.announce(:paths, "git ceiling set to:    #{ceiling_dir}         (git won't look for config or .git/ in this folder or above)")
end

After do
  aruba.announcer.announce(:gitmob_log, '<<-GITMOB_LOG')
  if aruba.environment['GITMOB_LOG_FILE']
    aruba.announcer.announce(:gitmob_log, File.read(expand_path(aruba.environment['GITMOB_LOG_FILE'])))
  else
    aruba.announcer.announce(:gitmob_log, "git-mob logs are available as STDOUT; use @announce-stdout")
  end
  aruba.announcer.announce(:gitmob_log, 'GITMOB_LOG')

  aruba.announcer.announce(:gitmessage, '<<-GITMESSAGE')
  gitmessage_file = expand_path('~/.gitmessage')
  if File.exist?(gitmessage_file)
    aruba.announcer.announce(:gitmessage, File.read(gitmessage_file))
  else
    aruba.announcer.announce(:gitmessage, "#{gitmessage_file} does not exist")
  end
  aruba.announcer.announce(:gitmessage, 'GITMESSAGE')
end
