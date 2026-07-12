package config

import "os"

type Config struct {
	Port        string
	DBPath      string
	OllamaURL   string
	OllamaModel string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8084"),
		DBPath:      getEnv("DB_PATH", "agent.db"),
		OllamaURL:   getEnv("OLLAMA_URL", "http://localhost:11434"),
		OllamaModel: getEnv("OLLAMA_MODEL", "phi3:mini"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	return v
}
