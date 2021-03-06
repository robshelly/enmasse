/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MessagingProjectLister helps list MessagingProjects.
type MessagingProjectLister interface {
	// List lists all MessagingProjects in the indexer.
	List(selector labels.Selector) (ret []*v1.MessagingProject, err error)
	// MessagingProjects returns an object that can list and get MessagingProjects.
	MessagingProjects(namespace string) MessagingProjectNamespaceLister
	MessagingProjectListerExpansion
}

// messagingProjectLister implements the MessagingProjectLister interface.
type messagingProjectLister struct {
	indexer cache.Indexer
}

// NewMessagingProjectLister returns a new MessagingProjectLister.
func NewMessagingProjectLister(indexer cache.Indexer) MessagingProjectLister {
	return &messagingProjectLister{indexer: indexer}
}

// List lists all MessagingProjects in the indexer.
func (s *messagingProjectLister) List(selector labels.Selector) (ret []*v1.MessagingProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessagingProject))
	})
	return ret, err
}

// MessagingProjects returns an object that can list and get MessagingProjects.
func (s *messagingProjectLister) MessagingProjects(namespace string) MessagingProjectNamespaceLister {
	return messagingProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MessagingProjectNamespaceLister helps list and get MessagingProjects.
type MessagingProjectNamespaceLister interface {
	// List lists all MessagingProjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MessagingProject, err error)
	// Get retrieves the MessagingProject from the indexer for a given namespace and name.
	Get(name string) (*v1.MessagingProject, error)
	MessagingProjectNamespaceListerExpansion
}

// messagingProjectNamespaceLister implements the MessagingProjectNamespaceLister
// interface.
type messagingProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MessagingProjects in the indexer for a given namespace.
func (s messagingProjectNamespaceLister) List(selector labels.Selector) (ret []*v1.MessagingProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessagingProject))
	})
	return ret, err
}

// Get retrieves the MessagingProject from the indexer for a given namespace and name.
func (s messagingProjectNamespaceLister) Get(name string) (*v1.MessagingProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("messagingproject"), name)
	}
	return obj.(*v1.MessagingProject), nil
}
