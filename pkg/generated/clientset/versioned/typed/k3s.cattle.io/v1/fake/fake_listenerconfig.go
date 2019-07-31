/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	k3scattleiov1 "github.com/rancher/k3s/pkg/apis/k3s.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeListenerConfigs implements ListenerConfigInterface
type FakeListenerConfigs struct {
	Fake *FakeK3sV1
	ns   string
}

var listenerconfigsResource = schema.GroupVersionResource{Group: "k3s.cattle.io", Version: "v1", Resource: "listenerconfigs"}

var listenerconfigsKind = schema.GroupVersionKind{Group: "k3s.cattle.io", Version: "v1", Kind: "ListenerConfig"}

// Get takes name of the listenerConfig, and returns the corresponding listenerConfig object, and an error if there is any.
func (c *FakeListenerConfigs) Get(name string, options v1.GetOptions) (result *k3scattleiov1.ListenerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(listenerconfigsResource, c.ns, name), &k3scattleiov1.ListenerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k3scattleiov1.ListenerConfig), err
}

// List takes label and field selectors, and returns the list of ListenerConfigs that match those selectors.
func (c *FakeListenerConfigs) List(opts v1.ListOptions) (result *k3scattleiov1.ListenerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(listenerconfigsResource, listenerconfigsKind, c.ns, opts), &k3scattleiov1.ListenerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k3scattleiov1.ListenerConfigList{ListMeta: obj.(*k3scattleiov1.ListenerConfigList).ListMeta}
	for _, item := range obj.(*k3scattleiov1.ListenerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested listenerConfigs.
func (c *FakeListenerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(listenerconfigsResource, c.ns, opts))

}

// Create takes the representation of a listenerConfig and creates it.  Returns the server's representation of the listenerConfig, and an error, if there is any.
func (c *FakeListenerConfigs) Create(listenerConfig *k3scattleiov1.ListenerConfig) (result *k3scattleiov1.ListenerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(listenerconfigsResource, c.ns, listenerConfig), &k3scattleiov1.ListenerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k3scattleiov1.ListenerConfig), err
}

// Update takes the representation of a listenerConfig and updates it. Returns the server's representation of the listenerConfig, and an error, if there is any.
func (c *FakeListenerConfigs) Update(listenerConfig *k3scattleiov1.ListenerConfig) (result *k3scattleiov1.ListenerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(listenerconfigsResource, c.ns, listenerConfig), &k3scattleiov1.ListenerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k3scattleiov1.ListenerConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeListenerConfigs) UpdateStatus(listenerConfig *k3scattleiov1.ListenerConfig) (*k3scattleiov1.ListenerConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(listenerconfigsResource, "status", c.ns, listenerConfig), &k3scattleiov1.ListenerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k3scattleiov1.ListenerConfig), err
}

// Delete takes name of the listenerConfig and deletes it. Returns an error if one occurs.
func (c *FakeListenerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(listenerconfigsResource, c.ns, name), &k3scattleiov1.ListenerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeListenerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(listenerconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &k3scattleiov1.ListenerConfigList{})
	return err
}

// Patch applies the patch and returns the patched listenerConfig.
func (c *FakeListenerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *k3scattleiov1.ListenerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(listenerconfigsResource, c.ns, name, pt, data, subresources...), &k3scattleiov1.ListenerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k3scattleiov1.ListenerConfig), err
}
