package config

import (
	"os"
	"strconv"
	"time"
)

func GetEnv(envName, defaultValue string) string {
	if v, ok := os.LookupEnv(envName); ok && len(v) > 0 {
		return v
	}
	return defaultValue
}

func GetEnvBool(envName string, defaultValue bool) bool {
	val := GetEnv(envName, strconv.FormatBool(defaultValue))
	if b, err := strconv.ParseBool(val); err == nil {
		return b
	}
	return defaultValue
}

func GetEnvInt(envName string, defaultValue int) int {
	val := GetEnv(envName, strconv.Itoa(defaultValue))
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return defaultValue
}

func GetEnvInt64(envName string, defaultValue int64) int64 {
	val := GetEnv(envName, strconv.FormatInt(defaultValue, 16))
	if i, err := strconv.ParseInt(val, 10, 64); err == nil {
		return i
	}
	return defaultValue
}

func GetEnvFloat(envName string, defaultValue float64) float64 {
	val := GetEnv(envName, strconv.FormatFloat(defaultValue, 'E', -1, 64))
	if f, err := strconv.ParseFloat(val, 64); err == nil {
		return f
	}
	return defaultValue
}

func GetEnvDuration(envName string, defaultValue time.Duration) time.Duration {
	val := GetEnv(envName, defaultValue.String())
	if t, err := time.ParseDuration(val); err == nil {
		return t
	}
	return defaultValue
}
