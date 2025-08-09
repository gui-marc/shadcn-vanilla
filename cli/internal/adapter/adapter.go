package adapter

import (
	"fmt"

	"github.com/gui-marc/shadcn-vanilla/internal/config"
)

// Adapter is the interface for transforming component templates
type Adapter interface {
	// GenerateComponent converts a canonical HTML template into
	// the specific syntax required by the target templating engine
	GenerateComponent(template string, componentConfig *config.ComponentConfig) (string, error)

	GetFileExtension() string // Returns the file extension for this adapter
}

// Factory method to get the Adapter based on engine name
func GetAdapter(engine string) (Adapter, error) {
	switch engine {
	case "gotempl":
		return NewGoTemplAdapter(), nil
	default:
		return nil, fmt.Errorf("unsupported engine: %s", engine)
	}
}
