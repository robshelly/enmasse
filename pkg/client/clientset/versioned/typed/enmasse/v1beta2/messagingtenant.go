/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	"time"

	v1beta2 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1beta2"
	scheme "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MessagingTenantsGetter has a method to return a MessagingTenantInterface.
// A group's client should implement this interface.
type MessagingTenantsGetter interface {
	MessagingTenants(namespace string) MessagingTenantInterface
}

// MessagingTenantInterface has methods to work with MessagingTenant resources.
type MessagingTenantInterface interface {
	Create(*v1beta2.MessagingTenant) (*v1beta2.MessagingTenant, error)
	Update(*v1beta2.MessagingTenant) (*v1beta2.MessagingTenant, error)
	UpdateStatus(*v1beta2.MessagingTenant) (*v1beta2.MessagingTenant, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta2.MessagingTenant, error)
	List(opts v1.ListOptions) (*v1beta2.MessagingTenantList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta2.MessagingTenant, err error)
	MessagingTenantExpansion
}

// messagingTenants implements MessagingTenantInterface
type messagingTenants struct {
	client rest.Interface
	ns     string
}

// newMessagingTenants returns a MessagingTenants
func newMessagingTenants(c *EnmasseV1beta2Client, namespace string) *messagingTenants {
	return &messagingTenants{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the messagingTenant, and returns the corresponding messagingTenant object, and an error if there is any.
func (c *messagingTenants) Get(name string, options v1.GetOptions) (result *v1beta2.MessagingTenant, err error) {
	result = &v1beta2.MessagingTenant{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingtenants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MessagingTenants that match those selectors.
func (c *messagingTenants) List(opts v1.ListOptions) (result *v1beta2.MessagingTenantList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.MessagingTenantList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingtenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messagingTenants.
func (c *messagingTenants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messagingtenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a messagingTenant and creates it.  Returns the server's representation of the messagingTenant, and an error, if there is any.
func (c *messagingTenants) Create(messagingTenant *v1beta2.MessagingTenant) (result *v1beta2.MessagingTenant, err error) {
	result = &v1beta2.MessagingTenant{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messagingtenants").
		Body(messagingTenant).
		Do().
		Into(result)
	return
}

// Update takes the representation of a messagingTenant and updates it. Returns the server's representation of the messagingTenant, and an error, if there is any.
func (c *messagingTenants) Update(messagingTenant *v1beta2.MessagingTenant) (result *v1beta2.MessagingTenant, err error) {
	result = &v1beta2.MessagingTenant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagingtenants").
		Name(messagingTenant.Name).
		Body(messagingTenant).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *messagingTenants) UpdateStatus(messagingTenant *v1beta2.MessagingTenant) (result *v1beta2.MessagingTenant, err error) {
	result = &v1beta2.MessagingTenant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagingtenants").
		Name(messagingTenant.Name).
		SubResource("status").
		Body(messagingTenant).
		Do().
		Into(result)
	return
}

// Delete takes name of the messagingTenant and deletes it. Returns an error if one occurs.
func (c *messagingTenants) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingtenants").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messagingTenants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingtenants").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched messagingTenant.
func (c *messagingTenants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta2.MessagingTenant, err error) {
	result = &v1beta2.MessagingTenant{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messagingtenants").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
