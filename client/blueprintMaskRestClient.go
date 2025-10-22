//nolint:dupl // generifying the rest clients would lead to a lot of unnecessary complexity
package kubernetes

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	"github.com/cloudogu/k8s-blueprint-lib/v2/api/v2"
)

const resourceMaskName = "blueprintmasks"

type BlueprintMaskInterface interface {
	// Create takes the representation of a v2.BlueprintMask and creates it. Returns the server's representation of the v2.BlueprintMask, and an error, if there is any.
	Create(ctx context.Context, blueprintMask *v2.BlueprintMask, opts metav1.CreateOptions) (*v2.BlueprintMask, error)

	// Update takes the representation of a v2.BlueprintMask and updates it. Returns the server's representation of the v2.BlueprintMask, and an error, if there is any.
	Update(ctx context.Context, blueprintMask *v2.BlueprintMask, opts metav1.UpdateOptions) (*v2.BlueprintMask, error)

	// UpdateStatus was generated because the type contains a Status member.
	UpdateStatus(ctx context.Context, blueprintMask *v2.BlueprintMask, opts metav1.UpdateOptions) (*v2.BlueprintMask, error)

	// Delete takes name of the v2.BlueprintMask and deletes it. Returns an error if one occurs.
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error

	// DeleteCollection deletes a collection of objects.
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error

	// Get takes name of the v2.BlueprintMask, and returns the corresponding v2.BlueprintMask object, and an error if there is any.
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v2.BlueprintMask, error)

	// List takes label and field selectors, and returns the list of v2.BlueprintMask that match those selectors.
	List(ctx context.Context, opts metav1.ListOptions) (*v2.BlueprintMaskList, error)

	// Watch returns a watch.Interface that watches the requested blueprintMasks.
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)

	// Patch applies the patch and returns the patched v2.BlueprintMask.
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v2.BlueprintMask, err error)
}

type blueprintMaskClient struct {
	client rest.Interface
	ns     string
}

func (d *blueprintMaskClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v2.BlueprintMask, err error) {
	result = &v2.BlueprintMask{}
	err = d.client.Get().
		Namespace(d.ns).
		Resource(resourceMaskName).
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (d *blueprintMaskClient) List(ctx context.Context, opts metav1.ListOptions) (result *v2.BlueprintMaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.BlueprintMaskList{}
	err = d.client.Get().
		Namespace(d.ns).
		Resource(resourceMaskName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (d *blueprintMaskClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return d.client.Get().
		Namespace(d.ns).
		Resource(resourceMaskName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (d *blueprintMaskClient) Create(ctx context.Context, blueprint *v2.BlueprintMask, opts metav1.CreateOptions) (result *v2.BlueprintMask, err error) {
	result = &v2.BlueprintMask{}
	err = d.client.Post().
		Namespace(d.ns).
		Resource(resourceMaskName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueprint).
		Do(ctx).
		Into(result)
	return
}

func (d *blueprintMaskClient) Update(ctx context.Context, blueprint *v2.BlueprintMask, opts metav1.UpdateOptions) (result *v2.BlueprintMask, err error) {
	result = &v2.BlueprintMask{}
	err = d.client.Put().
		Namespace(d.ns).
		Resource(resourceMaskName).
		Name(blueprint.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueprint).
		Do(ctx).
		Into(result)
	return
}

func (d *blueprintMaskClient) UpdateStatus(ctx context.Context, blueprint *v2.BlueprintMask, opts metav1.UpdateOptions) (result *v2.BlueprintMask, err error) {
	result = &v2.BlueprintMask{}
	err = d.client.Put().
		Namespace(d.ns).
		Resource(resourceMaskName).
		Name(blueprint.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueprint).
		Do(ctx).
		Into(result)
	return
}

func (d *blueprintMaskClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return d.client.Delete().
		Namespace(d.ns).
		Resource(resourceMaskName).
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (d *blueprintMaskClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return d.client.Delete().
		Namespace(d.ns).
		Resource(resourceMaskName).
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (d *blueprintMaskClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v2.BlueprintMask, err error) {
	result = &v2.BlueprintMask{}
	err = d.client.Patch(pt).
		Namespace(d.ns).
		Resource(resourceMaskName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
