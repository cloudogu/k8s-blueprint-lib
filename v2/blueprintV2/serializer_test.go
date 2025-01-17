package blueprintV2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudogu/blueprint-lib/bpcore"
	"github.com/cloudogu/blueprint-lib/v2/entities"
)

const (
	v1201 = "1.2.0-1"
	v2341 = "2.3.4-1"
)

const v0211 = "0.2.1-1"

func TestSerializeBlueprint_ok(t *testing.T) {

	type args struct {
		spec BlueprintV2
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
			}},
			`{"blueprintApi":"v2","config":{"global":{}}}`,
			assert.NoError,
		},
		{
			"dogus in blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Dogus: []entities.TargetDogu{
					{
						Name: "official/nginx", Version: v1201, TargetState: bpcore.TargetStatePresent, PlatformConfig: entities.PlatformConfig{ResourceConfig: entities.ResourceConfig{MinVolumeSize: "2Gi"}, ReverseProxyConfig: entities.ReverseProxyConfig{MaxBodySize: "2Gi", AdditionalConfig: "additional", RewriteTarget: "/"}},
					},
					{
						Name: "premium/jira", Version: v2341, TargetState: bpcore.TargetStateAbsent,
					},
				},
			}},
			`{"blueprintApi":"v2","dogus":[{"name":"official/nginx","version":"1.2.0-1","targetState":"present","platformConfig":{"resource":{"minVolumeSize":"2Gi"},"reverseProxy":{"maxBodySize":"2Gi","rewriteTarget":"/","additionalConfig":"additional"}}},{"name":"premium/jira","version":"2.3.4-1","targetState":"absent","platformConfig":{"resource":{},"reverseProxy":{}}}],"config":{"global":{}}}`,
			assert.NoError,
		},
		{
			"components in blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Components: []entities.TargetComponent{
					{Name: "k8s/blueprint-operator", Version: v0211, TargetState: bpcore.TargetStatePresent},
					{Name: "k8s/dogu-operator", Version: v2341, TargetState: bpcore.TargetStateAbsent, DeployConfig: map[string]interface{}{"deployNamespace": "ecosystem", "overwriteConfig": map[string]string{"key": "value"}}},
				},
			}},
			`{"blueprintApi":"v2","components":[{"name":"k8s/blueprint-operator","version":"0.2.1-1","targetState":"present"},{"name":"k8s/dogu-operator","version":"2.3.4-1","targetState":"absent","deployConfig":{"deployNamespace":"ecosystem","overwriteConfig":{"key":"value"}}}],"config":{"global":{}}}`,
			assert.NoError,
		},
		{
			"regular dogu config in blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Dogus: map[string]entities.CombinedDoguConfig{
						"ldap": {
							Config: entities.DoguConfig{
								Present: map[string]string{"container_config/memory_limit": "500m",
									"container_config/swap_limit":          "500m",
									"password_change/notification_enabled": "true",
								},
								Absent: []string{
									"password_change/mail_subject",
									"password_change/mail_text",
									"user_search_size_limit",
								},
							},
						},
					},
				},
			}},
			`{"blueprintApi":"v2","config":{"dogus":{"ldap":{"config":{"present":{"container_config/memory_limit":"500m","container_config/swap_limit":"500m","password_change/notification_enabled":"true"},"absent":["password_change/mail_subject","password_change/mail_text","user_search_size_limit"]},"sensitiveConfig":{}}},"global":{}}}`,
			assert.NoError,
		},
		{
			"sensitive dogu config in blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Dogus: map[string]entities.CombinedDoguConfig{
						"redmine": {
							SensitiveConfig: entities.SensitiveDoguConfig{
								Present: map[string]string{
									"my-secret-password":   "password-value",
									"my-secret-password-2": "password-value-2",
								},
								Absent: []string{
									"my-secret-password-3",
								},
							},
						},
					},
				},
			}},
			`{"blueprintApi":"v2","config":{"dogus":{"redmine":{"config":{},"sensitiveConfig":{"present":{"my-secret-password":"password-value","my-secret-password-2":"password-value-2"},"absent":["my-secret-password-3"]}}},"global":{}}}`,
			assert.NoError,
		},
		{
			"global config in blueprint",
			args{spec: BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Global: entities.GlobalConfig{
						Present: map[string]string{
							"key_provider": "pkcs1v15",
							"fqdn":         "ces.example.com",
							"admin_group":  "ces-admin",
						},
						Absent: []string{
							"default_dogu",
							"some_other_key",
						},
					},
				},
			}},
			`{"blueprintApi":"v2","config":{"global":{"present":{"admin_group":"ces-admin","fqdn":"ces.example.com","key_provider":"pkcs1v15"},"absent":["default_dogu","some_other_key"]}}}`,
			assert.NoError,
		},
		{
			name: "component config",
			args: args{
				spec: BlueprintV2{
					GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
					Components: []entities.TargetComponent{
						{Name: "k8s/name", Version: v2341, DeployConfig: map[string]interface{}{"key": "value"}},
					},
				},
			},
			want:    `{"blueprintApi":"v2","components":[{"name":"k8s/name","version":"2.3.4-1","targetState":"present","deployConfig":{"key":"value"}}],"config":{"global":{}}}`,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Serialize(tt.args.spec)
			if !tt.wantErr(t, err, fmt.Sprintf("SerializeBlueprint(%v)", tt.args.spec)) {
				return
			}
			assert.Equalf(t, tt.want, got, "SerializeBlueprint(%v)", tt.args.spec)
		})
	}
}

func TestDeserializeBlueprint_ok(t *testing.T) {
	type args struct {
		spec string
	}
	tests := []struct {
		name    string
		args    args
		want    BlueprintV2
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty blueprint",
			args{spec: `{"blueprintApi":"v2"}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
			},
			assert.NoError,
		},
		{
			"dogus in blueprint",
			args{spec: `{"blueprintApi":"v2","dogus":[{"name":"official/nginx","version":"1.2.0-1","targetState":"present"},{"name":"premium/jira","version":"2.3.4-1","targetState":"absent"}]}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Dogus: []entities.TargetDogu{
					{Name: "official/nginx", Version: v1201, TargetState: bpcore.TargetStatePresent},
					{Name: "premium/jira", Version: v2341, TargetState: bpcore.TargetStateAbsent},
				}},
			assert.NoError,
		},
		{
			"components in blueprint",
			args{spec: `{"blueprintApi":"v2","components":[{"name":"k8s/blueprint-operator","version":"0.2.1-1","targetState":"present"},{"name":"k8s/dogu-operator","version":"2.3.4-1","targetState":"absent"}]}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Components: []entities.TargetComponent{
					{Name: "k8s/blueprint-operator", Version: v0211, TargetState: bpcore.TargetStatePresent},
					{Name: "k8s/dogu-operator", Version: v2341, TargetState: bpcore.TargetStateAbsent},
				},
			},
			assert.NoError,
		},
		{
			"regular dogu config in blueprint",
			args{spec: `{"blueprintApi":"v2","config":{"dogus":{"ldap":{"config":{"present":{"container_config/memory_limit":"500m","container_config/swap_limit":"500m","password_change/notification_enabled":"true"},"absent":["password_change/mail_subject","password_change/mail_text","user_search_size_limit"]}}}}}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Dogus: map[string]entities.CombinedDoguConfig{
						"ldap": {
							Config: entities.DoguConfig{
								Present: map[string]string{
									"container_config/memory_limit":        "500m",
									"container_config/swap_limit":          "500m",
									"password_change/notification_enabled": "true",
								},
								Absent: []string{
									"password_change/mail_subject",
									"password_change/mail_text",
									"user_search_size_limit",
								},
							},
						},
					},
				},
			},
			assert.NoError,
		},
		{
			"sensitive dogu config in blueprint",
			args{spec: `{"blueprintApi":"v2","config":{"dogus":{"redmine":{"sensitiveConfig":{"present":{"my-secret-password":"password-value","my-secret-password-2":"password-value-2"},"absent":["my-secret-password-3"]}}}}}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Dogus: map[string]entities.CombinedDoguConfig{
						"redmine": {
							SensitiveConfig: entities.SensitiveDoguConfig{
								Present: map[string]string{
									"my-secret-password":   "password-value",
									"my-secret-password-2": "password-value-2",
								},
								Absent: []string{
									"my-secret-password-3",
								},
							},
						},
					},
				},
			},
			assert.NoError,
		},
		{
			"global config in blueprint",
			args{spec: `{"blueprintApi":"v2","config":{"global":{"present":{"admin_group":"ces-admin","fqdn":"ces.example.com","key_provider":"pkcs1v15"},"absent":["default_dogu","some_other_key"]}}}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Config: entities.TargetConfig{
					Global: entities.GlobalConfig{
						Present: map[string]string{
							"key_provider": "pkcs1v15",
							"fqdn":         "ces.example.com",
							"admin_group":  "ces-admin",
						},
						Absent: []string{
							"default_dogu",
							"some_other_key",
						},
					},
				},
			},
			assert.NoError,
		},
		{
			"component package config",
			args{spec: `{"blueprintApi":"v2","components":[{"name":"k8s/name","version":"2.3.4-1","targetState":"present","deployConfig":{"key":"value"}}],"config":{"global":{}}}`},
			BlueprintV2{
				GeneralBlueprint: bpcore.GeneralBlueprint{API: bpcore.V2},
				Components: []entities.TargetComponent{
					{Name: "k8s/name", Version: v2341, DeployConfig: map[string]interface{}{"key": "value"}},
				},
			},
			assert.NoError,
		},
		{
			"component package config error",
			args{spec: `{"blueprintApi":"v2","components":[{"name":"k8s/name","version":"3.2.1-1","targetState":"present","deployConfig":{"key"":""::"value:{"}}],"config":{"global":{}}}`},
			BlueprintV2{},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Deserialize(tt.args.spec)
			if !tt.wantErr(t, err, fmt.Sprintf("SerializeBlueprint(%v)", tt.args.spec)) {
				return
			}
			assert.Equalf(t, tt.want, got, "SerializeBlueprint(%v)", tt.args.spec)
		})
	}
}

func TestDeserializeBlueprint_errors(t *testing.T) {

	type args struct {
		rawBlueprint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "json syntax error",
			args:    args{`{a}`},
			want:    "cannot deserialize blueprint: invalid character 'a' looking for beginning of object key string",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Deserialize(tt.args.rawBlueprint)
			if !tt.wantErr(t, err, fmt.Sprintf("DeserializeBlueprint(%v)", tt.args.rawBlueprint)) {
				return
			}
			assert.ErrorContains(t, err, tt.want, "DeserializeBlueprint(%v)", tt.args.rawBlueprint)
		})
	}
}

func TestDeserializeBlueprint_testErrorType(t *testing.T) {
	_, err := Deserialize(`}`)
	require.Error(t, err)
	assert.ErrorContains(t, err, "cannot deserialize blueprint")
}
