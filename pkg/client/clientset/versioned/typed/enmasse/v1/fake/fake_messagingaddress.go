/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	enmassev1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMessagingAddresses implements MessagingAddressInterface
type FakeMessagingAddresses struct {
	Fake *FakeEnmasseV1
	ns   string
}

var messagingaddressesResource = schema.GroupVersionResource{Group: "enmasse.io", Version: "v1", Resource: "messagingaddresses"}

var messagingaddressesKind = schema.GroupVersionKind{Group: "enmasse.io", Version: "v1", Kind: "MessagingAddress"}

// Get takes name of the messagingAddress, and returns the corresponding messagingAddress object, and an error if there is any.
func (c *FakeMessagingAddresses) Get(name string, options v1.GetOptions) (result *enmassev1.MessagingAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(messagingaddressesResource, c.ns, name), &enmassev1.MessagingAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enmassev1.MessagingAddress), err
}

// List takes label and field selectors, and returns the list of MessagingAddresses that match those selectors.
func (c *FakeMessagingAddresses) List(opts v1.ListOptions) (result *enmassev1.MessagingAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(messagingaddressesResource, messagingaddressesKind, c.ns, opts), &enmassev1.MessagingAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &enmassev1.MessagingAddressList{ListMeta: obj.(*enmassev1.MessagingAddressList).ListMeta}
	for _, item := range obj.(*enmassev1.MessagingAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested messagingAddresses.
func (c *FakeMessagingAddresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(messagingaddressesResource, c.ns, opts))

}

// Create takes the representation of a messagingAddress and creates it.  Returns the server's representation of the messagingAddress, and an error, if there is any.
func (c *FakeMessagingAddresses) Create(messagingAddress *enmassev1.MessagingAddress) (result *enmassev1.MessagingAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(messagingaddressesResource, c.ns, messagingAddress), &enmassev1.MessagingAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enmassev1.MessagingAddress), err
}

// Update takes the representation of a messagingAddress and updates it. Returns the server's representation of the messagingAddress, and an error, if there is any.
func (c *FakeMessagingAddresses) Update(messagingAddress *enmassev1.MessagingAddress) (result *enmassev1.MessagingAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(messagingaddressesResource, c.ns, messagingAddress), &enmassev1.MessagingAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enmassev1.MessagingAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMessagingAddresses) UpdateStatus(messagingAddress *enmassev1.MessagingAddress) (*enmassev1.MessagingAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(messagingaddressesResource, "status", c.ns, messagingAddress), &enmassev1.MessagingAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enmassev1.MessagingAddress), err
}

// Delete takes name of the messagingAddress and deletes it. Returns an error if one occurs.
func (c *FakeMessagingAddresses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(messagingaddressesResource, c.ns, name), &enmassev1.MessagingAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMessagingAddresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(messagingaddressesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &enmassev1.MessagingAddressList{})
	return err
}

// Patch applies the patch and returns the patched messagingAddress.
func (c *FakeMessagingAddresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *enmassev1.MessagingAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(messagingaddressesResource, c.ns, name, pt, data, subresources...), &enmassev1.MessagingAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*enmassev1.MessagingAddress), err
}
