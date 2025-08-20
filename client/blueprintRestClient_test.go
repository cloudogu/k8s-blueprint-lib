package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"

	bpv2 "github.com/cloudogu/k8s-blueprint-lib/v2/api/v2"
)

var testCtx = context.Background()

var apiBasePathTestNS = fmt.Sprintf("/apis/%s/%s/namespaces/test/%s", bpv2.GroupVersion.Group, bpv2.GroupVersion.Version, resourceName)

func Test_blueprintClient_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, apiBasePathTestNS+"/testblueprint", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			blueprint := &bpv2.Blueprint{ObjectMeta: &v1.ObjectMeta{Name: "testblueprint", Namespace: "test"}}
			blueprintBytes, err := json.Marshal(blueprint)
			require.NoError(t, err)
			_, err = writer.Write(blueprintBytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		_, err = dClient.Get(testCtx, "testblueprint", v1.GetOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_Watch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, apiBasePathTestNS, request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)
			assert.Equal(t, "labelSelector=test&watch=true", request.URL.RawQuery)

			writer.Header().Add("content-type", "application/json")
			_, err := writer.Write([]byte("egal"))
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		_, err = dClient.Watch(testCtx, v1.ListOptions{LabelSelector: "test"})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		blueprint := &bpv2.Blueprint{ObjectMeta: &v1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPost, request.Method)
			assert.Equal(t, apiBasePathTestNS, request.URL.Path, resourceName)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdBlueprint := &bpv2.Blueprint{}
			require.NoError(t, json.Unmarshal(bytes, createdBlueprint))
			assert.Equal(t, "tocreate", createdBlueprint.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		_, err = dClient.Create(testCtx, blueprint, v1.CreateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		blueprint := &bpv2.Blueprint{ObjectMeta: &v1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, apiBasePathTestNS+"/tocreate", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdBlueprint := &bpv2.Blueprint{}
			require.NoError(t, json.Unmarshal(bytes, createdBlueprint))
			assert.Equal(t, "tocreate", createdBlueprint.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		_, err = dClient.Update(testCtx, blueprint, v1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_UpdateStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		blueprint := &bpv2.Blueprint{ObjectMeta: &v1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, apiBasePathTestNS+"/tocreate/status", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdBlueprint := &bpv2.Blueprint{}
			require.NoError(t, json.Unmarshal(bytes, createdBlueprint))
			assert.Equal(t, "tocreate", createdBlueprint.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		_, err = dClient.UpdateStatus(testCtx, blueprint, v1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, apiBasePathTestNS+"/testblueprint", request.URL.Path)

			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		err = dClient.Delete(testCtx, "testblueprint", v1.DeleteOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_DeleteCollection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, apiBasePathTestNS, request.URL.Path)
			assert.Equal(t, "labelSelector=test", request.URL.RawQuery)
			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		// when
		err = dClient.DeleteCollection(testCtx, v1.DeleteOptions{}, v1.ListOptions{LabelSelector: "test"})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintClient_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPatch, request.Method)
			assert.Equal(t, apiBasePathTestNS+"/testblueprint", request.URL.Path)
			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)
			assert.Equal(t, []byte("test"), bytes)
			result, err := json.Marshal(bpv2.Blueprint{})
			require.NoError(t, err)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(result)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.Blueprints("test")

		patchData := []byte("test")

		// when
		_, err = dClient.Patch(testCtx, "testblueprint", types.JSONPatchType, patchData, v1.PatchOptions{})

		// then
		require.NoError(t, err)
	})
}
