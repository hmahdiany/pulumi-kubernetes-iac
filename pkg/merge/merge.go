package merge

import (
	"os"

	"gopkg.in/yaml.v3"
)

// MergeValues merges a base values.yaml file with an override file
func MergeValues(baseFile, overrideFile string) (map[string]interface{}, error) {
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

	merged := MergeMaps(baseMap, overrideMap) // Modify baseMap directly
	return merged, nil
}

// Merge two maps recursively (overriding base values with override values)
func MergeMaps(base, override map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{}) // Create a NEW map

	for k, v := range base { // Copy base values first
		merged[k] = v
	}

	for k, v := range override {
		if baseValue, ok := merged[k]; ok { // Check in the *merged* map
			baseMap, okBase := baseValue.(map[string]interface{})
			overrideMap, okOverride := v.(map[string]interface{})
			if okBase && okOverride {
				merged[k] = MergeMaps(baseMap, overrideMap) // Recursive merge
				continue                                    // Important: Skip the direct assignment below
			}
		}
		merged[k] = v // Override or add new value
	}
	return merged
}
