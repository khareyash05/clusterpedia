// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/clusterpedia-io/api/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterImportPolicies implements ClusterImportPolicyInterface
type FakeClusterImportPolicies struct {
	Fake *FakePolicyV1alpha1
}

var clusterimportpoliciesResource = schema.GroupVersionResource{Group: "policy.clusterpedia.io", Version: "v1alpha1", Resource: "clusterimportpolicies"}

var clusterimportpoliciesKind = schema.GroupVersionKind{Group: "policy.clusterpedia.io", Version: "v1alpha1", Kind: "ClusterImportPolicy"}

// Get takes name of the clusterImportPolicy, and returns the corresponding clusterImportPolicy object, and an error if there is any.
func (c *FakeClusterImportPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterImportPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterimportpoliciesResource, name), &v1alpha1.ClusterImportPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterImportPolicy), err
}

// List takes label and field selectors, and returns the list of ClusterImportPolicies that match those selectors.
func (c *FakeClusterImportPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterImportPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterimportpoliciesResource, clusterimportpoliciesKind, opts), &v1alpha1.ClusterImportPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterImportPolicyList{ListMeta: obj.(*v1alpha1.ClusterImportPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterImportPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterImportPolicies.
func (c *FakeClusterImportPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterimportpoliciesResource, opts))
}

// Create takes the representation of a clusterImportPolicy and creates it.  Returns the server's representation of the clusterImportPolicy, and an error, if there is any.
func (c *FakeClusterImportPolicies) Create(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.CreateOptions) (result *v1alpha1.ClusterImportPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterimportpoliciesResource, clusterImportPolicy), &v1alpha1.ClusterImportPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterImportPolicy), err
}

// Update takes the representation of a clusterImportPolicy and updates it. Returns the server's representation of the clusterImportPolicy, and an error, if there is any.
func (c *FakeClusterImportPolicies) Update(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterImportPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterimportpoliciesResource, clusterImportPolicy), &v1alpha1.ClusterImportPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterImportPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterImportPolicies) UpdateStatus(ctx context.Context, clusterImportPolicy *v1alpha1.ClusterImportPolicy, opts v1.UpdateOptions) (*v1alpha1.ClusterImportPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterimportpoliciesResource, "status", clusterImportPolicy), &v1alpha1.ClusterImportPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterImportPolicy), err
}

// Delete takes name of the clusterImportPolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterImportPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterimportpoliciesResource, name, opts), &v1alpha1.ClusterImportPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterImportPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterimportpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterImportPolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterImportPolicy.
func (c *FakeClusterImportPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterImportPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterimportpoliciesResource, name, pt, data, subresources...), &v1alpha1.ClusterImportPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterImportPolicy), err
}
