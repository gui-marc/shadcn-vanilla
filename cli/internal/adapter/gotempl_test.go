package adapter_test

import (
	"strings"
	"testing"

	"github.com/gui-marc/shadcn-vanilla/internal/adapter"
	"github.com/gui-marc/shadcn-vanilla/internal/config"
	"github.com/stretchr/testify/require"
)

func TestGenerateComponentBasic(t *testing.T) {
	t.Helper()

	adapter := adapter.NewGoTemplAdapter()

	// Define a minimal component config
	compCfg := &config.ComponentConfig{
		Name:      "button",
		DefaultAs: "button",
		Variants: []config.ComponentVariant{
			{
				Name: "variant",
			},
			{
				Name: "size",
			},
		},
	}

	// Sample input HTML template with placeholders
	htmlTemplate := `
<{{ .As }} class="btn" data-variant="{{ .Variant }}" data-size="{{ .Size }}">
	{{ .Slot }}
</{{ .As }}>
`

	got, err := adapter.GenerateComponent(htmlTemplate, compCfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `package components

type ButtonVariant struct {
	Variant string
	Size string
}

templ Button(variant ButtonVariant) {
	<button class="btn" data-variant="{ variant.Variant }" data-size="{ variant.Size }">
		{ children... }
	</button>
}
`

	require.Equal(t, want, got)
}

func TestGenerateComponentEmptyVariants(t *testing.T) {
	t.Helper()

	adapter := adapter.NewGoTemplAdapter()

	compCfg := &config.ComponentConfig{
		Name:     "icon",
		Variants: []config.ComponentVariant{},
	}

	htmlTemplate := "<svg>{{ .Slot }}</svg>"

	got, err := adapter.GenerateComponent(htmlTemplate, compCfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantStruct := "type IconVariant struct {\n\n}"

	if !strings.Contains(got, wantStruct) {
		t.Errorf("expected struct declaration, got: %s", got)
	}

	if !strings.Contains(got, "<svg>") {
		t.Errorf("expected svg tag in output, got: %s", got)
	}
}
