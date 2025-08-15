package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDogu_DeepCopy(t *testing.T) {
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
				Absent:  false,
			},
			want: &Dogu{
				Name:    "test",
				Version: "1.2.3-4",
				Absent:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.given.DeepCopy(), "DeepCopy()")
		})
	}
}
