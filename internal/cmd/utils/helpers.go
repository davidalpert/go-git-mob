package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetEnvOrDefaultString(envKey, defaultValue string) string {
	v := os.Getenv(envKey)
	if v != "" {
		return v
	}

	return defaultValue
}

func GetEnvOrDefaultInt(envKey string, defaultValue int) int {
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

func GetEnvOrDefaultDuration(envKey string, defaultValue time.Duration) time.Duration {
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

func GetEnvOrDefaultBool(envKey string, defaultValue bool) bool {
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

func GetEnvOrDefaultLogLevel(envKey string, defaultValue log.Level) log.Level {
	v := os.Getenv(envKey)
	if v != "" {
		level, err := log.ParseLevel(v)
		if err != nil {
			panic(fmt.Errorf("failed parsing '%s' as log.Level: %v", v, err))
		}
		return level
	}

	return defaultValue
}

func ExitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func StringPointer(s string) *string {
	return &s
}

func StringInSlice(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
