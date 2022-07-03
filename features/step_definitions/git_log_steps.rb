# frozen_string_literal: true

Then('the most recent commit log should contain:') do |doc_string|
  msg = 'TBD'

  Dir.chdir(aruba.current_directory) do
    msg = `git log -1 --format=full`
  end

  aruba.announcer.announce(:git_log, msg)

  # `git log --format=full` formats the actual commit message indented
  # by four spaces:
  #
  # commit 5d9f8d0fa938735feb909c229c3e09c3dba4ec81
  # Author: Jane Doe <jane@example.com>
  # Commit: Jane Doe <jane@example.com>
  #
  #     empty mobbed commit
  #
  #     Co-Authored-By: Amy Doe <amy@findmypast.com>
  #
  # but the cucumber step formatter removes leading spaces from doc_strings
  # so let's add them back here, effectively right-shifting each line of
  # the doc_string by the same message indent `--format=full` uses:
  expected_commit_message_string = doc_string.gsub(/^/, '    ')

  expect(msg)
    .to match_output_string expected_commit_message_string
end
