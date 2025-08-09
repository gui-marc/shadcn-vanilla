// cli/internal/cmd/add.go
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gui-marc/shadcn-vanilla/internal/adapter"
	"github.com/gui-marc/shadcn-vanilla/internal/config"
	"github.com/gui-marc/shadcn-vanilla/internal/registry"
)

type AddCmd struct {
	Component string `arg:"" help:"Component name to add (e.g. button)"`
	Engine    string `short:"e" help:"Template engine to generate (e.g. vanilla, gotempl)" default:"gotempl"`
	Branch    string `help:"GitHub branch for registry" default:"main"`
}

func (a *AddCmd) Run() error {
	ctx := context.Background()

	// Load global config
	globalConfig, err := config.ParseGlobalConfig("components.yaml")
	if err != nil {
		return fmt.Errorf("failed to read global config: %w", err)
	}

	// Use engine from flag if provided, else fallback to global config default engine
	engine := a.Engine
	if engine == "" {
		engine = globalConfig.DefaultEngine
		if engine == "" {
			engine = "vanilla" // fallback final default
		}
	}

	// Setup GitHub registry client (replace with your repo)
	ghRegistry := registry.NewGitHubRegistry("gui-marc", "shadcn-vanilla", a.Branch)

	basePath := filepath.Join(globalConfig.ComponentsFolder, a.Component)
	// Fetch component YAML config
	yamlPath := filepath.Join(basePath, fmt.Sprintf("%s.yaml", a.Component))
	yamlData, err := ghRegistry.FetchFile(ctx, yamlPath)
	if err != nil {
		return fmt.Errorf("failed to fetch component YAML: %w", err)
	}

	componentConfig, err := config.ParseComponentConfig(yamlData)
	if err != nil {
		return fmt.Errorf("failed to parse component YAML: %w", err)
	}

	// Fetch component HTML template
	htmlPath := filepath.Join(basePath, fmt.Sprintf("%s.html", a.Component))
	htmlData, err := ghRegistry.FetchFile(ctx, htmlPath)
	if err != nil {
		return fmt.Errorf("failed to fetch component HTML template: %w", err)
	}

	// Get adapter for engine
	componentAdapter, err := adapter.GetAdapter(engine)
	if err != nil {
		return err
	}

	// Generate component code from template + config
	generated, err := componentAdapter.GenerateComponent(string(htmlData), &componentConfig)
	if err != nil {
		return fmt.Errorf("failed to generate component: %w", err)
	}

	// Write output component file to user project folder
	if err := os.MkdirAll(globalConfig.ComponentsFolder, 0755); err != nil {
		return err
	}

	// Save file with engine-specific extension
	outFile := filepath.Join(globalConfig.ComponentsFolder, fmt.Sprintf("%s.%s", a.Component, componentAdapter.GetFileExtension()))
	if err := os.WriteFile(outFile, []byte(generated), 0644); err != nil {
		return err
	}

	// Get component CSS and JS
	cssPath := filepath.Join(basePath, fmt.Sprintf("%s.css", a.Component))
	cssData, err := ghRegistry.FetchFile(ctx, cssPath)
	if err == nil {
		// Save css file in assets if exists
		cssFile := filepath.Join(globalConfig.AssetsFolder, fmt.Sprintf("%s.%s", a.Component, "css"))
		if err := os.WriteFile(cssFile, []byte(cssData), 0644); err != nil {
			return err
		}

		fmt.Printf("CSS file saved to %s\n", cssFile)
	}

	jsPath := filepath.Join(basePath, fmt.Sprintf("%s.js", a.Component))
	jsData, err := ghRegistry.FetchFile(ctx, jsPath)
	if err == nil {
		// Save js file in assets if exists
		jsFile := filepath.Join(globalConfig.AssetsFolder, fmt.Sprintf("%s.%s", a.Component, "js"))
		if err := os.WriteFile(jsFile, []byte(jsData), 0644); err != nil {
			return err
		}

		fmt.Printf("JS file saved to %s\n", jsFile)
	}

	fmt.Printf("Component %s added successfully to %s\n", a.Component, outFile)

	return nil
}
