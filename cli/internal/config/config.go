package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ComponentVariant struct {
	Name    string   `yaml:"name"`
	Values  []string `yaml:"values"`
	Default string   `yaml:"default"`
}

type ComponentConfig struct {
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	DefaultAs   string             `yaml:"default-as"`
	Variants    []ComponentVariant `yaml:"variants"`
}

type GlobalConfig struct {
	ComponentsFolder string `yaml:"components_folder"`
	AssetsFolder     string `yaml:"assets_folder"`
	RegistryURL      string `yaml:"registry_url"`
	DefaultEngine    string `yaml:"default_engine"`
}

func ParseComponentConfig(yamlData []byte) (ComponentConfig, error) {
	var config ComponentConfig
	err := yaml.Unmarshal(yamlData, &config)
	return config, err
}

func ParseGlobalConfig(path string) (GlobalConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return GlobalConfig{}, err
	}
	var cfg GlobalConfig
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
