package bpcore

import (
	"reflect"
	"testing"
)

func TestTargetState_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		state   TargetState
		want    []byte
		wantErr bool
	}{
		{"present", TargetStatePresent, []byte(`"present"`), false},
		{"absent", TargetStateAbsent, []byte(`"absent"`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.state.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestTargetState_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		state   TargetState
		args    args
		wantErr bool
	}{
		{"present", TargetStatePresent, args{[]byte(`"present"`)}, false},
		{"absent", TargetStateAbsent, args{[]byte(`"absent"`)}, false},
		{"parsing error", TargetStateAbsent, args{[]byte(`banane`)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.state.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
