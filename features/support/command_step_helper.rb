# frozen_string_literal: true

require 'aruba'
require 'aruba/cucumber'

# Aruba ------------------------------------------------

# aruba file    matchers: https://github.com/cucumber/aruba/blob/master/lib/aruba/cucumber/file.rb
# aruba command matchers: https://github.com/cucumber/aruba/blob/master/lib/aruba/cucumber/command.rb

# Step Helpers -----------------------------------------------

module Command
  # syntax wrappers around arub commands
  module StepHelpers

    # this helper provides some common logic around running a specific
    # command using Aruba's environment and path resolution, then
    # inspecting that output and matching it either as a regex match
    # or a string contains
    def run_command_and_validate_channel(cmd: '', fail_on_error: true, channel: 'stdout', negated: false, match_as_regex: false, content: '')
      run_command_and_stop(cmd, fail_on_error: fail_on_error)

      command = aruba.command_monitor.find(Aruba.platform.detect_ruby(cmd))

      matcher = case channel
                when 'output'; then :have_output
                when 'stderr'; then :have_output_on_stderr
                when 'stdout'; then :have_output_on_stdout
                end

      output_string_matcher = if match_as_regex
                                :an_output_string_matching
                              else
                                :an_output_string_including
                              end

      if negated
        expect(command).not_to send(matcher, send(output_string_matcher, content))
      else
        expect(command).to send(matcher, send(output_string_matcher, content))
      end
    end

    def validate_channel(channel: 'stdout', negated: false, match_as_regex: false, content: '')
      output = send("all_#{channel}")

      matcher = case channel
                when 'output'; then :have_output
                when 'stderr'; then :have_output_on_stderr
                when 'stdout'; then :have_output_on_stdout
                end

      output_string_matcher = if match_as_regex
                                :match_output_string
                              else
                                :include_output_string
                              end

      if negated
        expect(output).not_to send(output_string_matcher, content)
      else
        expect(output).to send(output_string_matcher, content)
      end
    end
  end
end

# rubocop:disable Style/MixinUsage
# this extend is at the global scope deliberately, to make these helpers available in step definitions
include Command::StepHelpers
# rubocop:enable Style/MixinUsage
