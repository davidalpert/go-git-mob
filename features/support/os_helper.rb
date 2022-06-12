module GitMob
  # syntax helpers for OS platform detection
  # adapted from: https://stackoverflow.com/a/171011/8997
  module OS
    def OS.windows?
      (/cygwin|mswin|mingw|bccwin|wince|emx/ =~ RUBY_PLATFORM) != nil
    end

    def OS.mac?
      (/darwin/ =~ RUBY_PLATFORM) != nil
    end

    def OS.unix?
      !OS.windows?
    end

    def OS.linux?
      OS.unix? and not OS.mac?
    end

    def OS.jruby?
      RUBY_ENGINE == 'jruby'
    end
  end
end

# rubocop:disable Style/MixinUsage
# this extend is at the global scope deliberately, to make these helpers available in step definitions
include GitMob::OS
# rubocop:enable Style/MixinUsage
