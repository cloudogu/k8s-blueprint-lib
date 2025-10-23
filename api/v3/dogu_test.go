package v3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogu_DeepCopy(t *testing.T) {
	falsePtr := false
	version1234 := "1.2.3-4"
	tests := []struct {
		name  string
		given *Dogu
		want  *Dogu
	}{
		{
			name: "empty",
			given: &Dogu{
				Name:    "test",
				Version: &version1234,
				Absent:  &falsePtr,
			},
			want: &Dogu{
				Name:    "test",
				Version: &version1234,
				Absent:  &falsePtr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.given.DeepCopy(), "DeepCopy()")
		})
	}
}
