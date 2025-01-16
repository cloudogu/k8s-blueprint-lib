package v2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cescommons "github.com/cloudogu/ces-commons-lib/dogu"
)

var (
	dogu1     = cescommons.SimpleName("dogu1")
	dogu2     = cescommons.SimpleName("dogu2")
	dogu1Key1 = DoguConfigKey{DoguName: dogu1, Key: "key1"}
)

func TestGlobalConfig_GetGlobalConfigKeys(t *testing.T) {
	var (
		globalKey1 = GlobalConfigKey("key1")
		globalKey2 = GlobalConfigKey("key2")
	)
	config := GlobalConfig{
		Present: map[GlobalConfigKey]GlobalConfigValue{
			globalKey1: "value",
		},
		Absent: []GlobalConfigKey{
			globalKey2,
		},
	}

	keys := config.GetGlobalConfigKeys()

	assert.ElementsMatch(t, keys, []GlobalConfigKey{globalKey1, globalKey2})
}

func TestConfig_GetDoguConfigKeys(t *testing.T) {
	var (
		nginx       = cescommons.SimpleName("nginx")
		postfix     = cescommons.SimpleName("postfix")
		nginxKey1   = DoguConfigKey{DoguName: nginx, Key: "key1"}
		nginxKey2   = DoguConfigKey{DoguName: nginx, Key: "key2"}
		postfixKey1 = DoguConfigKey{DoguName: postfix, Key: "key1"}
		postfixKey2 = DoguConfigKey{DoguName: postfix, Key: "key2"}
	)
	config := Config{
		Dogus: map[cescommons.SimpleName]CombinedDoguConfig{
			nginx: {
				DoguName: nginx,
				Config: DoguConfig{
					Present: map[DoguConfigKey]DoguConfigValue{
						nginxKey1: "value",
					},
					Absent: []DoguConfigKey{
						nginxKey2,
					},
				},
				SensitiveConfig: SensitiveDoguConfig{},
			},
			postfix: {
				DoguName: postfix,
				Config: DoguConfig{
					Present: map[DoguConfigKey]DoguConfigValue{
						postfixKey1: "value",
					},
					Absent: []DoguConfigKey{
						postfixKey2,
					},
				},
				SensitiveConfig: SensitiveDoguConfig{},
			},
		},
	}

	keys := config.GetDoguConfigKeys()

	assert.ElementsMatch(t, keys, []DoguConfigKey{nginxKey1, nginxKey2, postfixKey1, postfixKey2})
}

func TestConfig_GetSensitiveDoguConfigKeys(t *testing.T) {
	var (
		nginx       = cescommons.SimpleName("nginx")
		postfix     = cescommons.SimpleName("postfix")
		nginxKey1   = SensitiveDoguConfigKey{DoguName: nginx, Key: "key1"}
		nginxKey2   = SensitiveDoguConfigKey{DoguName: nginx, Key: "key2"}
		postfixKey1 = SensitiveDoguConfigKey{DoguName: postfix, Key: "key1"}
		postfixKey2 = SensitiveDoguConfigKey{DoguName: postfix, Key: "key2"}
	)
	config := Config{
		Dogus: map[cescommons.SimpleName]CombinedDoguConfig{
			nginx: {
				DoguName: nginx,
				SensitiveConfig: SensitiveDoguConfig{
					Present: map[SensitiveDoguConfigKey]SensitiveDoguConfigValue{
						nginxKey1: "value",
					},
					Absent: []SensitiveDoguConfigKey{
						nginxKey2,
					},
				},
			},
			postfix: {
				DoguName: postfix,
				SensitiveConfig: SensitiveDoguConfig{
					Present: map[SensitiveDoguConfigKey]SensitiveDoguConfigValue{
						postfixKey1: "value",
					},
					Absent: []SensitiveDoguConfigKey{
						postfixKey2,
					},
				},
			},
		},
	}

	keys := config.GetSensitiveDoguConfigKeys()

	assert.ElementsMatch(t, keys, []SensitiveDoguConfigKey{nginxKey1, nginxKey2, postfixKey1, postfixKey2})
}

func TestConfig_GetDogusWithChangedConfig(t *testing.T) {
	presentConfig := map[DoguConfigKey]DoguConfigValue{
		dogu1Key1: "val",
	}
	AbsentConfig := []DoguConfigKey{
		dogu1Key1,
	}
	emptyPresentConfig := map[DoguConfigKey]DoguConfigValue{}
	var emptyAbsentConfig []DoguConfigKey

	type args struct {
		doguConfig      DoguConfig
		withDogu2Change bool
	}

	var emptyResult []cescommons.SimpleName
	var tests = []struct {
		name string
		args args
		want []cescommons.SimpleName
	}{
		{
			name: "should get multiple Dogus",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: AbsentConfig}, withDogu2Change: true},
			want: []cescommons.SimpleName{dogu1, dogu2},
		},
		{
			name: "should get Dogus with changed present and absent config",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: AbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should get Dogus with changed present config",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: emptyAbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should get Dogus with changed absent config",
			args: args{doguConfig: DoguConfig{Present: emptyPresentConfig, Absent: AbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should not get Dogus with no config changes",
			args: args{doguConfig: DoguConfig{Present: emptyPresentConfig, Absent: emptyAbsentConfig}},
			want: emptyResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyDoguConfig := struct {
				Present map[DoguConfigKey]DoguConfigValue
				Absent  []DoguConfigKey
			}{}
			config := Config{
				Dogus: map[cescommons.SimpleName]CombinedDoguConfig{
					dogu1: {
						DoguName:        dogu1,
						Config:          tt.args.doguConfig,
						SensitiveConfig: emptyDoguConfig,
					},
				},
				Global: GlobalConfig{},
			}

			if tt.args.withDogu2Change {
				config.Dogus[dogu2] = CombinedDoguConfig{
					DoguName:        dogu2,
					Config:          tt.args.doguConfig,
					SensitiveConfig: emptyDoguConfig,
				}
			}

			changedDogus := config.GetDogusWithChangedConfig()
			assert.Len(t, changedDogus, len(tt.want))
			for _, doguName := range tt.want {
				assert.Contains(t, changedDogus, doguName)
			}
		})
	}
}

func TestConfig_GetDogusWithChangedSensitiveConfig(t *testing.T) {
	presentConfig := map[DoguConfigKey]DoguConfigValue{
		dogu1Key1: "val",
	}
	AbsentConfig := []DoguConfigKey{
		dogu1Key1,
	}
	emptyPresentConfig := map[DoguConfigKey]DoguConfigValue{}
	var emptyAbsentConfig []DoguConfigKey

	type args struct {
		doguConfig      DoguConfig
		withDogu2Change bool
	}

	var emptyResult []cescommons.SimpleName
	var tests = []struct {
		name string
		args args
		want []cescommons.SimpleName
	}{
		{
			name: "should get multiple Dogus",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: AbsentConfig}, withDogu2Change: true},
			want: []cescommons.SimpleName{dogu1, dogu2},
		},
		{
			name: "should get Dogus with changed present and absent config",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: AbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should get Dogus with changed present config",
			args: args{doguConfig: DoguConfig{Present: presentConfig, Absent: emptyAbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should get Dogus with changed absent config",
			args: args{doguConfig: DoguConfig{Present: emptyPresentConfig, Absent: AbsentConfig}},
			want: []cescommons.SimpleName{dogu1},
		},
		{
			name: "should not get Dogus with no config changes",
			args: args{doguConfig: DoguConfig{Present: emptyPresentConfig, Absent: emptyAbsentConfig}},
			want: emptyResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyDoguConfig := struct {
				Present map[DoguConfigKey]DoguConfigValue
				Absent  []DoguConfigKey
			}{}
			config := Config{
				Dogus: map[cescommons.SimpleName]CombinedDoguConfig{
					dogu1: {
						DoguName:        dogu1,
						Config:          emptyDoguConfig,
						SensitiveConfig: tt.args.doguConfig,
					},
				},
				Global: GlobalConfig{},
			}

			if tt.args.withDogu2Change {
				config.Dogus[dogu2] = CombinedDoguConfig{
					DoguName:        dogu2,
					Config:          emptyDoguConfig,
					SensitiveConfig: tt.args.doguConfig,
				}
			}

			changedDogus := config.GetDogusWithChangedSensitiveConfig()
			assert.Len(t, changedDogus, len(tt.want))
			for _, doguName := range tt.want {
				assert.Contains(t, changedDogus, doguName)
			}
		})
	}
}
