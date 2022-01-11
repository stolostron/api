// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/stolostron/api/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedClusterSetBindings implements ManagedClusterSetBindingInterface
type FakeManagedClusterSetBindings struct {
	Fake *FakeClusterV1alpha1
	ns   string
}

var managedclustersetbindingsResource = schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1alpha1", Resource: "managedclustersetbindings"}

var managedclustersetbindingsKind = schema.GroupVersionKind{Group: "cluster.open-cluster-management.io", Version: "v1alpha1", Kind: "ManagedClusterSetBinding"}

// Get takes name of the managedClusterSetBinding, and returns the corresponding managedClusterSetBinding object, and an error if there is any.
func (c *FakeManagedClusterSetBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedClusterSetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedclustersetbindingsResource, c.ns, name), &v1alpha1.ManagedClusterSetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedClusterSetBinding), err
}

// List takes label and field selectors, and returns the list of ManagedClusterSetBindings that match those selectors.
func (c *FakeManagedClusterSetBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedClusterSetBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedclustersetbindingsResource, managedclustersetbindingsKind, c.ns, opts), &v1alpha1.ManagedClusterSetBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedClusterSetBindingList{ListMeta: obj.(*v1alpha1.ManagedClusterSetBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedClusterSetBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedClusterSetBindings.
func (c *FakeManagedClusterSetBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedclustersetbindingsResource, c.ns, opts))

}

// Create takes the representation of a managedClusterSetBinding and creates it.  Returns the server's representation of the managedClusterSetBinding, and an error, if there is any.
func (c *FakeManagedClusterSetBindings) Create(ctx context.Context, managedClusterSetBinding *v1alpha1.ManagedClusterSetBinding, opts v1.CreateOptions) (result *v1alpha1.ManagedClusterSetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedclustersetbindingsResource, c.ns, managedClusterSetBinding), &v1alpha1.ManagedClusterSetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedClusterSetBinding), err
}

// Update takes the representation of a managedClusterSetBinding and updates it. Returns the server's representation of the managedClusterSetBinding, and an error, if there is any.
func (c *FakeManagedClusterSetBindings) Update(ctx context.Context, managedClusterSetBinding *v1alpha1.ManagedClusterSetBinding, opts v1.UpdateOptions) (result *v1alpha1.ManagedClusterSetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedclustersetbindingsResource, c.ns, managedClusterSetBinding), &v1alpha1.ManagedClusterSetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedClusterSetBinding), err
}

// Delete takes name of the managedClusterSetBinding and deletes it. Returns an error if one occurs.
func (c *FakeManagedClusterSetBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedclustersetbindingsResource, c.ns, name), &v1alpha1.ManagedClusterSetBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedClusterSetBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedclustersetbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedClusterSetBindingList{})
	return err
}

// Patch applies the patch and returns the patched managedClusterSetBinding.
func (c *FakeManagedClusterSetBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedClusterSetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedclustersetbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedClusterSetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedClusterSetBinding), err
}
