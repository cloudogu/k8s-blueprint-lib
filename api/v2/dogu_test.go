package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDogu_DeepCopy(t *testing.T) {
	falsePtr := false
	tests := []struct {
		name  string
		given *Dogu
		want  *Dogu
	}{
		{
			name: "empty",
			given: &Dogu{
				Name:    "test",
				Version: "1.2.3-4",
				Absent:  &falsePtr,
			},
			want: &Dogu{
				Name:    "test",
				Version: "1.2.3-4",
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
