// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/mattermost/mattermost-operator/pkg/client/clientset/versioned/typed/mattermost/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMattermostV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMattermostV1alpha1) ClusterInstallations(namespace string) v1alpha1.ClusterInstallationInterface {
	return &FakeClusterInstallations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMattermostV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
