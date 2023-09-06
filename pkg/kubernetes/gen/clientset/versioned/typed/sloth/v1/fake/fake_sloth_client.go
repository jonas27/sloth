// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/jonas27/sloth/pkg/kubernetes/gen/clientset/versioned/typed/sloth/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSlothV1 struct {
	*testing.Fake
}

func (c *FakeSlothV1) PrometheusServiceLevels(namespace string) v1.PrometheusServiceLevelInterface {
	return &FakePrometheusServiceLevels{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSlothV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
