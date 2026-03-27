// Package config holds application-level configuration constants and helpers.
package config

// Config holds runtime configuration for the supply-chain monitor server.
type Config struct {
	ServerAddress string // HTTP listen address
	DataFilePath  string // Path to the CSV dataset
	WorkerCount   int    // Number of concurrent processing workers
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *Config {
	return &Config{
		ServerAddress: ":8080",
		DataFilePath:  "data/supply_chain_dataset_full_500k.csv",
		WorkerCount:   10,
	}
}
