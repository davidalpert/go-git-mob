require_relative 'os_helper'

Before do
end

After do
end

Before('@announce-git-log') do
  aruba.announcer.activate :git_log
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