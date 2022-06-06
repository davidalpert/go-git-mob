# frozen_string_literal: true

require 'aruba'
require 'aruba/cucumber'
require_relative 'matchers/match_string'
require 'yaml'

# Aruba ------------------------------------------------

# aruba file    matchers: https://github.com/cucumber/aruba/blob/master/lib/aruba/cucumber/file.rb
# aruba command matchers: https://github.com/cucumber/aruba/blob/master/lib/aruba/cucumber/command.rb

Aruba.configure do |config|
  config.exit_timeout = 1 # seconds
end

# Step Helpers -----------------------------------------------

Given(/PENDING/) do
  pending
end

module GitMob
  # syntax wrappers around arub commands
  module StepHelpers
    def git_mob_bin
      bin = File.join(aruba.root_directory, 'bin', 'git-mob')
      raise "'#{bin}' not found; did you run 'make build'?" unless File.exist?(bin)

      bin
    end

    def home_dir
      relative_dir('~/')
    end

    def current_dir
      relative_dir('.')
    end

    def relative_dir(path)
      abs_dir(path).delete_prefix(aruba.root_directory)[1..-1]
    end

    def relative_dir_from_abs(path)
      path.delete_prefix(aruba.root_directory)[1..-1]
    end

    def abs_dir(path)
      expanded = path
      with_environment do
        expanded = expand_path(path)
      end
      expanded
    end

    # TODO: this syntax will break with aruba 2.1.0
    def all_output_includes(s, negated: false)
      if negated
        expect(all_commands)
          .to_not include_an_object have_output an_output_string_including(s)
      else
        expect(all_commands)
          .to include_an_object have_output an_output_string_including(s)
      end
    end

    # TODO: this syntax will break with aruba 2.1.0
    def all_output_matches(s, negated: false)
      if negated
        expect(all_commands)
          .to_not include_an_object have_output an_output_string_matching(s)
      else
        expect(all_commands)
          .to include_an_object have_output an_output_string_matching(s)
      end
    end
  end
end

# rubocop:disable Style/MixinUsage
# this extend is at the global scope deliberately, to make these helpers available in step definitions
include GitMob::StepHelpers
# rubocop:enable Style/MixinUsage
