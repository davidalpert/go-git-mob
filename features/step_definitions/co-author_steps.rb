# Then('the file named {string} should include these coauthors:')
Then(/^(?:a|the) file(?: named)? "([^"]*)" should (not )?include these coauthors:$/) do |file, negated, doc_string|
  expanded_path = expand_path(file)

  expect(file).to be_an_existing_file

  @actual   = JSON.parse(File.read(expanded_path).chomp)
  @expected = JSON.parse(doc_string.chomp)

  # puts "ACTUAL\n#{@actual}\n"
  # puts "EXPECTED\n#{@expected}\n"

  if negated
    expect(@actual["coauthors"]).to_not include(@expected)
  else
    expect(@actual["coauthors"]).to include(@expected)
  end
end

Then('the current git author should be {string} {string}') do |name, email|
  run_command_and_validate_channel(
    cmd: 'git config user.name',
    fail_on_error: true,
    channel: 'stdout',
    negated: false,
    match_as_regex: false,
    content: name
  )

  run_command_and_validate_channel(
    cmd: 'git config user.email',
    fail_on_error: true,
    channel: 'stdout',
    negated: false,
    match_as_regex: false,
    content: email
  )
end