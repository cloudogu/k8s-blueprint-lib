package entities

import (
	"github.com/cloudogu/k8s-blueprint-lib/json/bpcore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"maps"
	"slices"
	"testing"
)

func TestDeployConfig_DeepCopy(t *testing.T) {
	t.Run("should copy empty config", func(t *testing.T) {
		t.Run("should copy empty config map", func(t *testing.T) {
			// given
			input := DeployConfig{}

			// when
			actual := input.DeepCopy()

			// then
			require.NotSame(t, &input, &actual)
			require.NotNil(t, actual)
			assert.Empty(t, actual)
		})
		t.Run("should copy single value map", func(t *testing.T) {
			// given
			const key = "a cool key"
			input := DeployConfig{key: "a cool value"}

			// when
			actual := input.DeepCopy()

			// then
			require.NotSame(t, &input, &actual)
			require.NotNil(t, actual)
			mappedActual := unaliasDeployConfig(t, *actual)
			mappedInput := unaliasDeployConfig(t, input)
			assert.Equal(t, mappedInput[key], mappedActual[key])
			assert.Equal(t, slices.Collect(maps.Keys(mappedInput)), slices.Collect(maps.Keys(mappedActual)))
			assert.NotSame(t, &mappedInput, &mappedActual)
		})
		t.Run("mutation of one DeployConfig should not mutate a copy", func(t *testing.T) {
			// given
			input := &DeployConfig{"redmine": CombinedDoguConfig{Config: DoguConfig{
				Absent: []string{"redmineKeyToBeDeleted"},
			}}}

			// when
			actual := input.DeepCopy()
			mappedInput := unaliasDeployConfig(t, *input)
			mappedInput["redmine"] = CombinedDoguConfig{Config: DoguConfig{}} // overwrite ALL the things \o/

			// then
			_ = unaliasDeployConfig(t, *actual)
			assert.Empty(t, mappedInput["redmine"])
		})
	})
}

func TestTargetComponent_DeepCopyInto(t *testing.T) {
	t.Run("should copy empty target component", func(t *testing.T) {
		// given
		input := TargetComponent{}
		actual := TargetComponent{}

		// when
		input.DeepCopyInto(&actual)

		// then
		require.NotSame(t, &input, &actual)
		require.NotNil(t, actual)
		assert.Empty(t, actual)
	})
	t.Run("should copy simple target component", func(t *testing.T) {
		// given
		inputDeployConfig := DeployConfig{"redmine": CombinedDoguConfig{Config: DoguConfig{
			Absent: []string{"redmineKeyToBeDeleted"},
		}}}
		input := TargetComponent{
			Name:         "k8s/my-comp",
			Version:      "1.2.3",
			TargetState:  bpcore.TargetStateAbsent.String(),
			DeployConfig: inputDeployConfig,
		}
		actual := TargetComponent{}

		// when
		input.DeepCopyInto(&actual)

		// then
		require.NotSame(t, &input, &actual)
		require.NotNil(t, actual)
		expected := TargetComponent{
			Name:        "k8s/my-comp",
			Version:     "1.2.3",
			TargetState: bpcore.TargetStateAbsent.String(),
			DeployConfig: DeployConfig{"redmine": CombinedDoguConfig{Config: DoguConfig{
				Absent: []string{"redmineKeyToBeDeleted"},
			}}},
		}
		// types get sadly unaliased. Check the values here instead
		assert.Equal(t, expected.Name, actual.Name)
		assert.Equal(t, expected.Version, actual.Version)
		assert.Equal(t, expected.TargetState, actual.TargetState)
		expectedMapConfig := unaliasDeployConfig(t, expected.DeployConfig)
		actualMapConfig := unaliasDeployConfig(t, actual.DeployConfig)
		expectedRedmineConfig := expectedMapConfig["redmine"]
		actualRedmineConfig := actualMapConfig["redmine"]
		assert.NotSame(t, &expectedRedmineConfig, &actualRedmineConfig)
	})
}

func unaliasDeployConfig(t *testing.T, config DeployConfig) map[string]interface{} {
	t.Helper()

	return config
}
