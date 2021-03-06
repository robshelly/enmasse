/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1"
	scheme "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MessagingInfrastructuresGetter has a method to return a MessagingInfrastructureInterface.
// A group's client should implement this interface.
type MessagingInfrastructuresGetter interface {
	MessagingInfrastructures(namespace string) MessagingInfrastructureInterface
}

// MessagingInfrastructureInterface has methods to work with MessagingInfrastructure resources.
type MessagingInfrastructureInterface interface {
	Create(*v1.MessagingInfrastructure) (*v1.MessagingInfrastructure, error)
	Update(*v1.MessagingInfrastructure) (*v1.MessagingInfrastructure, error)
	UpdateStatus(*v1.MessagingInfrastructure) (*v1.MessagingInfrastructure, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MessagingInfrastructure, error)
	List(opts metav1.ListOptions) (*v1.MessagingInfrastructureList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MessagingInfrastructure, err error)
	MessagingInfrastructureExpansion
}

// messagingInfrastructures implements MessagingInfrastructureInterface
type messagingInfrastructures struct {
	client rest.Interface
	ns     string
}

// newMessagingInfrastructures returns a MessagingInfrastructures
func newMessagingInfrastructures(c *EnmasseV1Client, namespace string) *messagingInfrastructures {
	return &messagingInfrastructures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the messagingInfrastructure, and returns the corresponding messagingInfrastructure object, and an error if there is any.
func (c *messagingInfrastructures) Get(name string, options metav1.GetOptions) (result *v1.MessagingInfrastructure, err error) {
	result = &v1.MessagingInfrastructure{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MessagingInfrastructures that match those selectors.
func (c *messagingInfrastructures) List(opts metav1.ListOptions) (result *v1.MessagingInfrastructureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MessagingInfrastructureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messagingInfrastructures.
func (c *messagingInfrastructures) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a messagingInfrastructure and creates it.  Returns the server's representation of the messagingInfrastructure, and an error, if there is any.
func (c *messagingInfrastructures) Create(messagingInfrastructure *v1.MessagingInfrastructure) (result *v1.MessagingInfrastructure, err error) {
	result = &v1.MessagingInfrastructure{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		Body(messagingInfrastructure).
		Do().
		Into(result)
	return
}

// Update takes the representation of a messagingInfrastructure and updates it. Returns the server's representation of the messagingInfrastructure, and an error, if there is any.
func (c *messagingInfrastructures) Update(messagingInfrastructure *v1.MessagingInfrastructure) (result *v1.MessagingInfrastructure, err error) {
	result = &v1.MessagingInfrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		Name(messagingInfrastructure.Name).
		Body(messagingInfrastructure).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *messagingInfrastructures) UpdateStatus(messagingInfrastructure *v1.MessagingInfrastructure) (result *v1.MessagingInfrastructure, err error) {
	result = &v1.MessagingInfrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		Name(messagingInfrastructure.Name).
		SubResource("status").
		Body(messagingInfrastructure).
		Do().
		Into(result)
	return
}

// Delete takes name of the messagingInfrastructure and deletes it. Returns an error if one occurs.
func (c *messagingInfrastructures) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messagingInfrastructures) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched messagingInfrastructure.
func (c *messagingInfrastructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MessagingInfrastructure, err error) {
	result = &v1.MessagingInfrastructure{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messaginginfrastructures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
