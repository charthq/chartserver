/*
Copyright 2019 Replicated, Inc..

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/charthq/chartserver/pkg/apis/chartserver/v1beta1"
	scheme "github.com/charthq/chartserver/pkg/client/chartserverclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChartVersionsGetter has a method to return a ChartVersionInterface.
// A group's client should implement this interface.
type ChartVersionsGetter interface {
	ChartVersions(namespace string) ChartVersionInterface
}

// ChartVersionInterface has methods to work with ChartVersion resources.
type ChartVersionInterface interface {
	Create(*v1beta1.ChartVersion) (*v1beta1.ChartVersion, error)
	Update(*v1beta1.ChartVersion) (*v1beta1.ChartVersion, error)
	UpdateStatus(*v1beta1.ChartVersion) (*v1beta1.ChartVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ChartVersion, error)
	List(opts v1.ListOptions) (*v1beta1.ChartVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ChartVersion, err error)
	ChartVersionExpansion
}

// chartVersions implements ChartVersionInterface
type chartVersions struct {
	client rest.Interface
	ns     string
}

// newChartVersions returns a ChartVersions
func newChartVersions(c *ChartserverV1beta1Client, namespace string) *chartVersions {
	return &chartVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chartVersion, and returns the corresponding chartVersion object, and an error if there is any.
func (c *chartVersions) Get(name string, options v1.GetOptions) (result *v1beta1.ChartVersion, err error) {
	result = &v1beta1.ChartVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChartVersions that match those selectors.
func (c *chartVersions) List(opts v1.ListOptions) (result *v1beta1.ChartVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ChartVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chartVersions.
func (c *chartVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chartversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a chartVersion and creates it.  Returns the server's representation of the chartVersion, and an error, if there is any.
func (c *chartVersions) Create(chartVersion *v1beta1.ChartVersion) (result *v1beta1.ChartVersion, err error) {
	result = &v1beta1.ChartVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chartversions").
		Body(chartVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a chartVersion and updates it. Returns the server's representation of the chartVersion, and an error, if there is any.
func (c *chartVersions) Update(chartVersion *v1beta1.ChartVersion) (result *v1beta1.ChartVersion, err error) {
	result = &v1beta1.ChartVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartversions").
		Name(chartVersion.Name).
		Body(chartVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *chartVersions) UpdateStatus(chartVersion *v1beta1.ChartVersion) (result *v1beta1.ChartVersion, err error) {
	result = &v1beta1.ChartVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartversions").
		Name(chartVersion.Name).
		SubResource("status").
		Body(chartVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the chartVersion and deletes it. Returns an error if one occurs.
func (c *chartVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chartVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched chartVersion.
func (c *chartVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ChartVersion, err error) {
	result = &v1beta1.ChartVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chartversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
