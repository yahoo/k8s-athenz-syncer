/*
Copyright 2019, Oath Inc.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/yahoo/k8s-athenz-syncer/pkg/apis/athenz/v1"
	scheme "github.com/yahoo/k8s-athenz-syncer/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AthenzDomainsGetter has a method to return a AthenzDomainInterface.
// A group's client should implement this interface.
type AthenzDomainsGetter interface {
	AthenzDomains() AthenzDomainInterface
}

// AthenzDomainInterface has methods to work with AthenzDomain resources.
type AthenzDomainInterface interface {
	Create(ctx context.Context, athenzDomain *v1.AthenzDomain, opts metav1.CreateOptions) (*v1.AthenzDomain, error)
	Update(ctx context.Context, athenzDomain *v1.AthenzDomain, opts metav1.UpdateOptions) (*v1.AthenzDomain, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AthenzDomain, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AthenzDomainList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AthenzDomain, err error)
	AthenzDomainExpansion
}

// athenzDomains implements AthenzDomainInterface
type athenzDomains struct {
	client rest.Interface
}

// newAthenzDomains returns a AthenzDomains
func newAthenzDomains(c *AthenzV1Client) *athenzDomains {
	return &athenzDomains{
		client: c.RESTClient(),
	}
}

// Get takes name of the athenzDomain, and returns the corresponding athenzDomain object, and an error if there is any.
func (c *athenzDomains) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Get().
		Resource("athenzdomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AthenzDomains that match those selectors.
func (c *athenzDomains) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AthenzDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AthenzDomainList{}
	err = c.client.Get().
		Resource("athenzdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested athenzDomains.
func (c *athenzDomains) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("athenzdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a athenzDomain and creates it.  Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *athenzDomains) Create(ctx context.Context, athenzDomain *v1.AthenzDomain, opts metav1.CreateOptions) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Post().
		Resource("athenzdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(athenzDomain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a athenzDomain and updates it. Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *athenzDomains) Update(ctx context.Context, athenzDomain *v1.AthenzDomain, opts metav1.UpdateOptions) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Put().
		Resource("athenzdomains").
		Name(athenzDomain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(athenzDomain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the athenzDomain and deletes it. Returns an error if one occurs.
func (c *athenzDomains) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("athenzdomains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *athenzDomains) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("athenzdomains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched athenzDomain.
func (c *athenzDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Patch(pt).
		Resource("athenzdomains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
