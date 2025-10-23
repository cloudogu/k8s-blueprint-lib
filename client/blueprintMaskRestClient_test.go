package kubernetes

import (
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

var maskApiBasePathTestNS = fmt.Sprintf("/apis/%s/%s/namespaces/test/%s", bpv2.GroupVersion.Group, bpv2.GroupVersion.Version, resourceMaskName)

func Test_blueprintMaskClient_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, maskApiBasePathTestNS+"/testblueprintmask", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			blueprintMask := &bpv2.BlueprintMask{ObjectMeta: v1.ObjectMeta{Name: "testblueprintmask", Namespace: "test"}}
			blueprintMaskBytes, err := json.Marshal(blueprintMask)
			require.NoError(t, err)
			_, err = writer.Write(blueprintMaskBytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		_, err = dClient.Get(testCtx, "testblueprintmask", v1.GetOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodGet, request.Method)
			assert.Equal(t, maskApiBasePathTestNS, request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			blueprintMaskList := bpv2.BlueprintMaskList{}
			blueprintMask := &bpv2.BlueprintMask{ObjectMeta: v1.ObjectMeta{Name: "testblueprintmask", Namespace: "test"}}
			blueprintMaskList.Items = append(blueprintMaskList.Items, *blueprintMask)
			blueprintMaskBytes, err := json.Marshal(blueprintMaskList)
			require.NoError(t, err)
			_, err = writer.Write(blueprintMaskBytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		_, err = dClient.List(testCtx, v1.ListOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_Watch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, maskApiBasePathTestNS, request.URL.Path)
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
		dClient := client.BlueprintMasks("test")

		// when
		_, err = dClient.Watch(testCtx, v1.ListOptions{LabelSelector: "test"})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		blueprintMask := &bpv2.BlueprintMask{ObjectMeta: v1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPost, request.Method)
			assert.Equal(t, maskApiBasePathTestNS, request.URL.Path, resourceName)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdBlueprintMask := &bpv2.BlueprintMask{}
			require.NoError(t, json.Unmarshal(bytes, createdBlueprintMask))
			assert.Equal(t, "tocreate", createdBlueprintMask.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		_, err = dClient.Create(testCtx, blueprintMask, v1.CreateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		blueprintMask := &bpv2.BlueprintMask{ObjectMeta: v1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, maskApiBasePathTestNS+"/tocreate", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdBlueprintMask := &bpv2.BlueprintMask{}
			require.NoError(t, json.Unmarshal(bytes, createdBlueprintMask))
			assert.Equal(t, "tocreate", createdBlueprintMask.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		_, err = dClient.Update(testCtx, blueprintMask, v1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, maskApiBasePathTestNS+"/testblueprintmask", request.URL.Path)

			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		err = dClient.Delete(testCtx, "testblueprintmask", v1.DeleteOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_DeleteCollection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, maskApiBasePathTestNS, request.URL.Path)
			assert.Equal(t, "labelSelector=test", request.URL.RawQuery)
			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := newForConfig(&config)
		require.NoError(t, err)
		dClient := client.BlueprintMasks("test")

		// when
		err = dClient.DeleteCollection(testCtx, v1.DeleteOptions{}, v1.ListOptions{LabelSelector: "test"})

		// then
		require.NoError(t, err)
	})
}

func Test_blueprintMaskClient_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPatch, request.Method)
			assert.Equal(t, maskApiBasePathTestNS+"/testblueprintmask", request.URL.Path)
			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)
			assert.Equal(t, []byte("test"), bytes)
			result, err := json.Marshal(bpv2.BlueprintMask{})
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
		dClient := client.BlueprintMasks("test")

		patchData := []byte("test")

		// when
		_, err = dClient.Patch(testCtx, "testblueprintmask", types.JSONPatchType, patchData, v1.PatchOptions{})

		// then
		require.NoError(t, err)
	})
}
