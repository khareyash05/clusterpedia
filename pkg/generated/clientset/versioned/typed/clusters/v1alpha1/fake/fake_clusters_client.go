// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned/typed/clusters/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeClustersV1alpha1 struct {
	*testing.Fake
}

func (c *FakeClustersV1alpha1) PediaClusters() v1alpha1.PediaClusterInterface {
	return &FakePediaClusters{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeClustersV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}