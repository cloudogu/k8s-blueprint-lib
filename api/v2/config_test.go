package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigEntry_Validate(t *testing.T) {
	testValue := "testValue"
	truePtr := true
	type fields struct {
		Key       string
		Absent    *bool
		Value     *string
		SecretRef *Reference
		ConfigRef *Reference
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Absent without secretRef and value is valid",
			fields: fields{
				Key:    "testKey",
				Absent: &truePtr,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Absent with secretRef is not valid",
			fields: fields{
				Key:       "testKey",
				Absent:    &truePtr,
				SecretRef: &Reference{Name: "testSecret"},
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorContains(t, err, "absent entries cannot have value, configRef or secretRef")
			},
		},
		{
			name: "Absent with value is not valid",
			fields: fields{
				Key:    "testKey",
				Absent: &truePtr,
				Value:  &testValue,
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorContains(t, err, "absent entries cannot have value, configRef or secretRef")
			},
		},
		{
			name: "Only value is valid",
			fields: fields{
				Key:   "testKey",
				Value: &testValue,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Only SecretRef is valid",
			fields: fields{
				Key:       "testKey",
				SecretRef: &Reference{Name: "testSecret"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Only ConfigRef is valid",
			fields: fields{
				Key:       "testKey",
				ConfigRef: &Reference{Name: "testConfigMap"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Value and SecretRef is not valid",
			fields: fields{
				Key:       "testKey",
				Value:     &testValue,
				SecretRef: &Reference{Name: "testSecret"},
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorContains(t, err, "config entries can have either a value, configRef or a secretRef")
			},
		},
		{
			name: "Value, ConfigRef, SecretRef is not valid",
			fields: fields{
				Key:       "testKey",
				Value:     &testValue,
				SecretRef: &Reference{Name: "testSecret"},
				ConfigRef: &Reference{Name: "testConfigMap"},
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorContains(t, err, "config entries can have either a value, configRef or a secretRef")
			},
		},
		{
			name: "Without Value and SecretRef is not valid",
			fields: fields{
				Key: "testKey",
			},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorContains(t, err, "config entries can have either a value, configRef or a secretRef")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConfigEntry{
				Key:       tt.fields.Key,
				Absent:    tt.fields.Absent,
				Value:     tt.fields.Value,
				SecretRef: tt.fields.SecretRef,
				ConfigRef: tt.fields.ConfigRef,
			}
			tt.wantErr(t, c.Validate(), fmt.Sprintf("Validate()"))
		})
	}
}
