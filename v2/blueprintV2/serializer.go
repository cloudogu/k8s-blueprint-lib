package blueprintV2

import (
	"encoding/json"
	"fmt"
)

// Serialize takes a Blueprint v2 and returns its string representation.
func Serialize(blueprintDTO BlueprintV2) (string, error) {
	serializedBlueprint, err := json.Marshal(blueprintDTO)
	if err != nil {
		return "", fmt.Errorf("cannot serialize blueprint: %w", err)
	}

	return string(serializedBlueprint), nil
}

// Deserialize takes a string and returns its Blueprint v2 representation.
func Deserialize(rawBlueprint string) (BlueprintV2, error) {
	blueprintDTO := BlueprintV2{}

	err := json.Unmarshal([]byte(rawBlueprint), &blueprintDTO)
	if err != nil {
		return BlueprintV2{}, fmt.Errorf("cannot deserialize blueprint: %w", err)
	}

	return blueprintDTO, nil
}
