require "rspec/expectations/version"

RSpec::Matchers.define :match_string do |expected|
  match do |actual|
    actual.force_encoding("UTF-8")
    @expected = Regexp.new(unescape_text(expected), Regexp::MULTILINE)
    @actual   = sanitize_text(actual)

    values_match? @expected, @actual
  end

  diffable

  description { "string matches: #{description_of expected}" }
end