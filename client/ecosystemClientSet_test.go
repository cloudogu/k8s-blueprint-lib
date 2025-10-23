package kubernetes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestNewForConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		config := &rest.Config{}

		// when
		clientSet, err := newForConfig(config)

		// then
		require.NoError(t, err)
		require.NotNil(t, clientSet)
	})

}

func TestEcoSystemV1Alpha1Client_Dogus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		config := &rest.Config{}
		clientSet, err := newForConfig(config)
		require.NoError(t, err)
		require.NotNil(t, clientSet)

		// when
		client := clientSet.Blueprints("ecosystem")

		// then
		require.NotNil(t, client)
	})
}

func TestNewClientSet(t *testing.T) {
	type args struct {
		config    *rest.Config
		clientSet *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		args    args
		want    *ClientSet
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error on nil config",
			args: args{
				config:    nil,
				clientSet: nil,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return error on nil clientSet",
			args: args{
				config:    &rest.Config{},
				clientSet: nil,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientSet(tt.args.config, tt.args.clientSet)
			if !tt.wantErr(t, err, fmt.Sprintf("NewClientSet(%v, %v)", tt.args.config, tt.args.clientSet)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewClientSet(%v, %v)", tt.args.config, tt.args.clientSet)
		})
	}
}
