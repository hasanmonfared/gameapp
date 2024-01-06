package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"
)

func Load() *Config {
	var k = koanf.New(".")
	// Load default values using the confmap provider.
	// We provide a flat map with the "." delimiter.
	// A nested map can be loaded by setting the delimiter to an empty string "".
	k.Load(confmap.Provider(defaultConfig, "."), nil)

	// Load YAML config and merge into the previously loaded config (because we can).
	k.Load(file.Provider("../config.yml"), yaml.Parser())

	k.Load(env.Provider("GAMEAPP_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "GAMEAPP_")), "_", ".", -1)
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
