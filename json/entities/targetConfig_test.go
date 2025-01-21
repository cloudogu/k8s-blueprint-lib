package entities

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDoguConfigMap_DeepCopy(t *testing.T) {
	t.Run("should copy empty config map", func(t *testing.T) {
		// given
		input := &DoguConfigMap{}

		// when
		actual := input.DeepCopy()

		// then
		require.NotSame(t, input, actual)
		require.NotNil(t, actual)
		assert.Empty(t, actual)
	})
	t.Run("should copy single value map", func(t *testing.T) {
		// given
		input := &DoguConfigMap{"redmine": CombinedDoguConfig{Config: DoguConfig{
			Absent: []string{"redmineKeyToBeDeleted"},
		}}}

		// when
		actual := input.DeepCopy()

		// then
		require.NotSame(t, input, actual)
		require.NotNil(t, actual)
		mappedActual := unaliasDoguConfig(t, *actual)
		mappedInput := unaliasDoguConfig(t, *input)
		mappedActualSlice := mappedActual["redmine"].Config.Absent
		mappedInputSlice := mappedInput["redmine"].Config.Absent
		assert.Equal(t, []string{"redmineKeyToBeDeleted"}, mappedActual["redmine"].Config.Absent)
		assert.NotSame(t, &mappedActualSlice, &mappedInputSlice)
	})
	t.Run("mutation of one DoguConfigMap should not mutate a copy", func(t *testing.T) {
		// given
		input := &DoguConfigMap{"redmine": CombinedDoguConfig{Config: DoguConfig{
			Absent: []string{"redmineKeyToBeDeleted"},
		}}}

		// when
		actual := input.DeepCopy()
		mappedInput := unaliasDoguConfig(t, *input)
		mappedInput["redmine"] = CombinedDoguConfig{Config: DoguConfig{}} // overwrite ALL the things \o/

		// then
		mappedActual := unaliasDoguConfig(t, *actual)
		assert.Equal(t, []string{"redmineKeyToBeDeleted"}, mappedActual["redmine"].Config.Absent)
		assert.Empty(t, mappedInput["redmine"])
	})
}

func TestTargetConfig_DeepCopyInto(t *testing.T) {
	t.Run("should copy empty objects", func(t *testing.T) {
		// given
		input := &TargetConfig{}

		// when
		actual := &TargetConfig{}
		input.DeepCopyInto(actual)

		// then
		require.NotNil(t, actual)
		assert.Empty(t, actual)
	})
	t.Run("should copy simple config objects", func(t *testing.T) {
		// given
		input := &TargetConfig{
			Dogus: DoguConfigMap{"redmine": CombinedDoguConfig{Config: DoguConfig{
				Present: map[string]string{"hello": "world"}}}},
			Global: GlobalConfig{
				Present: map[string]string{"hello": "golang"}},
		}

		// when
		actual := &TargetConfig{}
		input.DeepCopyInto(actual)

		// then
		require.NotNil(t, actual)
		assert.NotEmpty(t, actual)
		mappedActualDogus := unaliasDoguConfig(t, actual.Dogus)

		assert.Equal(t, DoguConfig{Present: map[string]string{"hello": "world"}}, mappedActualDogus["redmine"].Config)
		assert.Equal(t, GlobalConfig{Present: map[string]string{"hello": "golang"}}, actual.Global)
	})
}

func unaliasDoguConfig(t *testing.T, config DoguConfigMap) map[string]CombinedDoguConfig {
	t.Helper()

	return config
}
