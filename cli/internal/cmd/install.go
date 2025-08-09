package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gui-marc/shadcn-vanilla/internal/config"
	"gopkg.in/yaml.v3"
)

type InstallCmd struct{}

func (i *InstallCmd) Run() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter components folder (default: components): ")
	componentsFolder, _ := reader.ReadString('\n')
	componentsFolder = strings.TrimSpace(componentsFolder)
	if componentsFolder == "" {
		componentsFolder = "components"
	}

	fmt.Print("Enter assets folder (default: assets): ")
	assetsFolder, _ := reader.ReadString('\n')
	assetsFolder = strings.TrimSpace(assetsFolder)
	if assetsFolder == "" {
		assetsFolder = "assets"
	}

	fmt.Print("Enter default template engine (default: gotempl): ")
	defaultEngine, _ := reader.ReadString('\n')
	defaultEngine = strings.TrimSpace(defaultEngine)
	if defaultEngine == "" {
		defaultEngine = "gotempl"
	}

	globalConfig := config.GlobalConfig{
		ComponentsFolder: componentsFolder,
		AssetsFolder:     assetsFolder,
		RegistryURL:      "https://github.com/gui-marc/shadcn-vanilla", // fixed for now
		DefaultEngine:    defaultEngine,
	}

	data, err := yaml.Marshal(globalConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal global config: %w", err)
	}

	if err := os.WriteFile("components.yaml", data, 0644); err != nil {
		return fmt.Errorf("failed to write components.yaml: %w", err)
	}

	fmt.Println("Global configuration saved to components.yaml")
	fmt.Println("You can edit this file later to customize the setup.")

	return nil
}
