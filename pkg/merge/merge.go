package merge

import (
	"os"

	"gopkg.in/yaml.v3"
)

// MergeValues merges a base values.yaml file with an override file
func MergeValues(baseFile string, overrideFile string) (map[string]interface{}, error) {
	// Read base values file
	baseContent, err := os.ReadFile(baseFile)
	if err != nil {
		return nil, err
	}

	// Read override file
	overrideContent, err := os.ReadFile(overrideFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal YAML
	baseMap := make(map[string]interface{})
	overrideMap := make(map[string]interface{})
	if err := yaml.Unmarshal(baseContent, &baseMap); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(overrideContent, &overrideMap); err != nil {
		return nil, err
	}

	// Merge override into base
	for k, v := range overrideMap {
		baseMap[k] = v
	}

	return MergeMaps(baseMap, overrideMap), nil
}

// Merge two maps recursively (overriding base values with override values)
func MergeMaps(base, override map[string]interface{}) map[string]interface{} {
	for key, value := range override {
		if baseValue, ok := base[key]; ok {
			// Merge nested maps recursively
			baseMap, okBase := baseValue.(map[string]interface{})
			overrideMap, okOverride := value.(map[string]interface{})
			if okBase && okOverride {
				base[key] = MergeMaps(baseMap, overrideMap)
				continue
			}
		}
		// Override base value with override value
		base[key] = value
	}
	return base
}
