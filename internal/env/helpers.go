package env

import (
	"fmt"
	"github.com/apex/log"
	"github.com/mitchellh/go-homedir"
	"os"
	"strconv"
	"time"
)

var (
	HomeDir string
)

func GetValueOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetValueOrDefaultString(envKey, defaultValue string) string {
	v := os.Getenv(envKey)
	if v != "" {
		return v
	}

	return defaultValue
}

func GetValueOrDefaultInt(envKey string, defaultValue int) int {
	v := os.Getenv(envKey)
	if v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("failed converting '%s' to an int: %v", v, err))
		}
		return i
	}

	return defaultValue
}

func GetValueOrDefaultDuration(envKey string, defaultValue time.Duration) time.Duration {
	v := os.Getenv(envKey)
	if v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			panic(fmt.Errorf("failed parsing '%s' as a time.Duration: %v", v, err))
		}
		return d
	}

	return defaultValue
}

func GetValueOrDefaultBool(envKey string, defaultValue bool) bool {
	v := os.Getenv(envKey)
	if v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Errorf("failed parsing '%s' as a bool: %v", v, err))
		}
		return b
	}

	return defaultValue
}

func GetValueOrDefaultLogLevel(envKey string, defaultValue log.Level) log.Level {
	v := os.Getenv(envKey)
	if v != "" {
		return log.MustParseLevel(v)
	}

	return defaultValue
}

func init() {
	if d, err := homedir.Dir(); err != nil {
		panic(fmt.Errorf("could not find home directory: %v", err))
	} else {
		HomeDir = d
	}
}
