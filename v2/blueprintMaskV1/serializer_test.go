package blueprintMaskV1

import (
	"fmt"
	"testing"

	"github.com/cloudogu/cesapp-lib/core"
	"github.com/stretchr/testify/assert"

	"github.com/cloudogu/blueprint-lib/bpcore"
)

var (
	version1_2_0_1, _ = core.ParseVersion("1.2.0-1")
	version3_0_2_2, _ = core.ParseVersion("3.0.2-2")
)

func TestSerializeBlueprintMask_ok(t *testing.T) {
	type args struct {
		spec BlueprintMaskV1
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty blueprint mask",
			args{spec: BlueprintMaskV1{
				GeneralBlueprintMask: bpcore.GeneralBlueprintMask{API: bpcore.MaskV1},
			}},
			`{"blueprintMaskApi":"v1","blueprintMaskId":"","dogus":null}`,
			assert.NoError,
		},
		{
			"dogus in blueprint mask",
			args{spec: BlueprintMaskV1{
				GeneralBlueprintMask: bpcore.GeneralBlueprintMask{API: bpcore.MaskV1},
				Dogus: []MaskTargetDogu{
					{Name: "official/nginx", Version: version1_2_0_1.String(), TargetState: bpcore.TargetStatePresent},
					{Name: "premium/jira", Version: version3_0_2_2.String(), TargetState: bpcore.TargetStateAbsent},
				},
			}},
			`{"blueprintMaskApi":"v1","blueprintMaskId":"","dogus":[{"name":"official/nginx","version":"1.2.0-1","targetState":"present"},{"name":"premium/jira","version":"3.0.2-2","targetState":"absent"}]}`,
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Serialize(tt.args.spec)
			if !tt.wantErr(t, err, fmt.Sprintf("SerializeBlueprintMask(%v)", tt.args.spec)) {
				return
			}
			assert.Equalf(t, tt.want, got, "SerializeBlueprintMask(%v)", tt.args.spec)
		})
	}
}

func TestDeserializeBlueprintMask_ok(t *testing.T) {
	type args struct {
		spec string
	}
	tests := []struct {
		name    string
		args    args
		want    BlueprintMaskV1
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty blueprint mask",
			args{spec: `{"blueprintMaskApi":"v1"}`},
			BlueprintMaskV1{GeneralBlueprintMask: bpcore.GeneralBlueprintMask{API: bpcore.MaskV1}},
			assert.NoError,
		},
		{
			"dogus in blueprint mask",
			args{spec: `{"blueprintMaskApi":"v1","dogus":[{"name":"official/nginx","version":"1.2.0-1","targetState":"present"},{"name":"premium/jira","version":"3.0.2-2","targetState":"absent"}]}`},
			BlueprintMaskV1{
				GeneralBlueprintMask: bpcore.GeneralBlueprintMask{API: bpcore.MaskV1},
				Dogus: []MaskTargetDogu{
					{Name: "official/nginx", Version: version1_2_0_1.String(), TargetState: bpcore.TargetStatePresent},
					{Name: "premium/jira", Version: version3_0_2_2.String(), TargetState: bpcore.TargetStateAbsent},
				}},
			assert.NoError,
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

func TestDeserializeBlueprintMask_errors(t *testing.T) {
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
			want:    "cannot deserialize blueprint mask: invalid character 'a' looking for beginning of object key string",
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
