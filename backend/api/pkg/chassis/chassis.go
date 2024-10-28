package chassis

import "os"

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing required environment variable: " + key)
	}

	return val
}
