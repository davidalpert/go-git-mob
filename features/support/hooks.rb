Before do
end

After do
end

Before('@announce-git-log') do
  aruba.announcer.activate :git_log
end
