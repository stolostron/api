// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/stolostron/api/client/addon/clientset/versioned/typed/addon/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAddonV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAddonV1alpha1) ClusterManagementAddOns() v1alpha1.ClusterManagementAddOnInterface {
	return &FakeClusterManagementAddOns{c}
}

func (c *FakeAddonV1alpha1) ManagedClusterAddOns(namespace string) v1alpha1.ManagedClusterAddOnInterface {
	return &FakeManagedClusterAddOns{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAddonV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
