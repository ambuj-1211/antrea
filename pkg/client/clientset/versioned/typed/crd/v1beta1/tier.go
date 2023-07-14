// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "antrea.io/antrea/pkg/apis/crd/v1beta1"
	scheme "antrea.io/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TiersGetter has a method to return a TierInterface.
// A group's client should implement this interface.
type TiersGetter interface {
	Tiers() TierInterface
}

// TierInterface has methods to work with Tier resources.
type TierInterface interface {
	Create(ctx context.Context, tier *v1beta1.Tier, opts v1.CreateOptions) (*v1beta1.Tier, error)
	Update(ctx context.Context, tier *v1beta1.Tier, opts v1.UpdateOptions) (*v1beta1.Tier, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Tier, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TierList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Tier, err error)
	TierExpansion
}

// tiers implements TierInterface
type tiers struct {
	client rest.Interface
}

// newTiers returns a Tiers
func newTiers(c *CrdV1beta1Client) *tiers {
	return &tiers{
		client: c.RESTClient(),
	}
}

// Get takes name of the tier, and returns the corresponding tier object, and an error if there is any.
func (c *tiers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Tier, err error) {
	result = &v1beta1.Tier{}
	err = c.client.Get().
		Resource("tiers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tiers that match those selectors.
func (c *tiers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TierList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.TierList{}
	err = c.client.Get().
		Resource("tiers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tiers.
func (c *tiers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tiers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tier and creates it.  Returns the server's representation of the tier, and an error, if there is any.
func (c *tiers) Create(ctx context.Context, tier *v1beta1.Tier, opts v1.CreateOptions) (result *v1beta1.Tier, err error) {
	result = &v1beta1.Tier{}
	err = c.client.Post().
		Resource("tiers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tier).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tier and updates it. Returns the server's representation of the tier, and an error, if there is any.
func (c *tiers) Update(ctx context.Context, tier *v1beta1.Tier, opts v1.UpdateOptions) (result *v1beta1.Tier, err error) {
	result = &v1beta1.Tier{}
	err = c.client.Put().
		Resource("tiers").
		Name(tier.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tier).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tier and deletes it. Returns an error if one occurs.
func (c *tiers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tiers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tiers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("tiers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tier.
func (c *tiers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Tier, err error) {
	result = &v1beta1.Tier{}
	err = c.client.Patch(pt).
		Resource("tiers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
