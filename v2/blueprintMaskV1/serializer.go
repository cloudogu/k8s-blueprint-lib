package blueprintMaskV1

import (
	"encoding/json"
	"fmt"
)

// Serialize takes a BlueprintMask v1 and returns its string representation.
func Serialize(blueprintDTO BlueprintMaskV1) (string, error) {
	serializedMask, err := json.Marshal(blueprintDTO)
	if err != nil {
		return "", fmt.Errorf("cannot serialize blueprint mask: %w", err)
	}
	return string(serializedMask), nil
}

// Deserialize takes a string and returns its BlueprintMask v1 representation.
func Deserialize(rawBlueprintMask string) (BlueprintMaskV1, error) {
	blueprintMaskDTO := BlueprintMaskV1{}

	err := json.Unmarshal([]byte(rawBlueprintMask), &blueprintMaskDTO)

	if err != nil {
		return BlueprintMaskV1{}, fmt.Errorf("cannot deserialize blueprint mask: %w", err)
	}

	return blueprintMaskDTO, nil
}
