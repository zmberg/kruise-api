/*
Copyright 2022 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodePodProbes implements NodePodProbeInterface
type FakeNodePodProbes struct {
	Fake *FakeAppsV1alpha1
}

var nodepodprobesResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "nodepodprobes"}

var nodepodprobesKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "NodePodProbe"}

// Get takes name of the nodePodProbe, and returns the corresponding nodePodProbe object, and an error if there is any.
func (c *FakeNodePodProbes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodePodProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodepodprobesResource, name), &v1alpha1.NodePodProbe{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePodProbe), err
}

// List takes label and field selectors, and returns the list of NodePodProbes that match those selectors.
func (c *FakeNodePodProbes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodePodProbeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodepodprobesResource, nodepodprobesKind, opts), &v1alpha1.NodePodProbeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodePodProbeList{ListMeta: obj.(*v1alpha1.NodePodProbeList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodePodProbeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodePodProbes.
func (c *FakeNodePodProbes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodepodprobesResource, opts))
}

// Create takes the representation of a nodePodProbe and creates it.  Returns the server's representation of the nodePodProbe, and an error, if there is any.
func (c *FakeNodePodProbes) Create(ctx context.Context, nodePodProbe *v1alpha1.NodePodProbe, opts v1.CreateOptions) (result *v1alpha1.NodePodProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodepodprobesResource, nodePodProbe), &v1alpha1.NodePodProbe{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePodProbe), err
}

// Update takes the representation of a nodePodProbe and updates it. Returns the server's representation of the nodePodProbe, and an error, if there is any.
func (c *FakeNodePodProbes) Update(ctx context.Context, nodePodProbe *v1alpha1.NodePodProbe, opts v1.UpdateOptions) (result *v1alpha1.NodePodProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodepodprobesResource, nodePodProbe), &v1alpha1.NodePodProbe{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePodProbe), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodePodProbes) UpdateStatus(ctx context.Context, nodePodProbe *v1alpha1.NodePodProbe, opts v1.UpdateOptions) (*v1alpha1.NodePodProbe, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodepodprobesResource, "status", nodePodProbe), &v1alpha1.NodePodProbe{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePodProbe), err
}

// Delete takes name of the nodePodProbe and deletes it. Returns an error if one occurs.
func (c *FakeNodePodProbes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodepodprobesResource, name), &v1alpha1.NodePodProbe{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodePodProbes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodepodprobesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodePodProbeList{})
	return err
}

// Patch applies the patch and returns the patched nodePodProbe.
func (c *FakeNodePodProbes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodePodProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodepodprobesResource, name, pt, data, subresources...), &v1alpha1.NodePodProbe{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePodProbe), err
}
