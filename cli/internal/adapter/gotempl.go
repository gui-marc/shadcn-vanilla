package adapter

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/gui-marc/shadcn-vanilla/internal/config"
)

type GoTemplAdapter struct{}

func NewGoTemplAdapter() Adapter {
	return &GoTemplAdapter{}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// GenerateComponent generates a GoTempl component file content by executing the Go template in htmlTemplate
// using the componentConfig as data.
func (a *GoTemplAdapter) GenerateComponent(htmlTemplate string, componentConfig *config.ComponentConfig) (string, error) {
	// Parse the incoming htmlTemplate as a Go template
	tmpl, err := template.New("component").Parse(htmlTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing html template: %w", err)
	}

	// Execute the template with dummy values for placeholders that appear in the template.
	// We want to produce a string that still contains GoTempl placeholders like { variant.X }
	// So we inject these placeholders as data values so that they appear verbatim.

	// Create a map for the template execution with keys matching the placeholders:
	execData := make(map[string]any)

	// Provide `.As` with the tag name
	execData["As"] = componentConfig.DefaultAs

	// Provide `.Slot` with the placeholder for content
	// Since `.Slot` in the template becomes { children... } in final GoTempl syntax
	execData["Slot"] = "{ children... }"

	execData["Class"] = "{ props.Class }"                                                                 // Always include Class
	execData["Attributes"] = "{ for key, value := range props.Attributes { \"+key+\"+\"=\"+\"value\" } }" // Placeholder for additional attributes

	// Provide each variant field with a placeholder string like `{ props.VariantName }`
	for _, variant := range componentConfig.Variants {
		execData[capitalize(variant.Name)] = fmt.Sprintf("{ props.%s }", capitalize(variant.Name))
	}

	// Execute the template into a buffer
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, execData)
	if err != nil {
		return "", fmt.Errorf("executing html template: %w", err)
	}

	// Build the variant struct fields for the Go struct definition
	var structFields []string
	for _, variant := range componentConfig.Variants {
		field := fmt.Sprintf("\t%s string", capitalize(variant.Name))
		structFields = append(structFields, field)
	}

	structFields = append(structFields, "\tAttributes map[string]any") // Always include Attributes
	structFields = append(structFields, "\tClass string")              // Always include Class

	packageDefinition := "package components\n\n"

	structDef := fmt.Sprintf("type %sProps struct {\n%s\n}\n\n", capitalize(componentConfig.Name), strings.Join(structFields, "\n"))

	// Build the templ block with the executed HTML
	templFunc := fmt.Sprintf("templ %s(props %sProps) {\n", capitalize(componentConfig.Name), capitalize(componentConfig.Name))

	// Indent HTML lines with a tab
	var indentedLines []string
	for _, line := range strings.Split(buf.String(), "\n") {
		if line == "" {
			continue // skip empty lines
		}
		indentedLines = append(indentedLines, "\t"+line)
	}
	htmlWithIndent := strings.Join(indentedLines, "\n")

	templFunc += htmlWithIndent + "\n}\n"

	return packageDefinition + structDef + templFunc, nil
}

// GetFileExtension implements Adapter.
func (a *GoTemplAdapter) GetFileExtension() string {
	return "templ"
}
