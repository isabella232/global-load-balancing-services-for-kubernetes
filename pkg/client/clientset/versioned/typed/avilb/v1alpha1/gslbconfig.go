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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "gitlab.eng.vmware.com/orion/mcc/pkg/apis/avilb/v1alpha1"
	scheme "gitlab.eng.vmware.com/orion/mcc/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GSLBConfigsGetter has a method to return a GSLBConfigInterface.
// A group's client should implement this interface.
type GSLBConfigsGetter interface {
	GSLBConfigs(namespace string) GSLBConfigInterface
}

// GSLBConfigInterface has methods to work with GSLBConfig resources.
type GSLBConfigInterface interface {
	Create(*v1alpha1.GSLBConfig) (*v1alpha1.GSLBConfig, error)
	Update(*v1alpha1.GSLBConfig) (*v1alpha1.GSLBConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GSLBConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.GSLBConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBConfig, err error)
	GSLBConfigExpansion
}

// gSLBConfigs implements GSLBConfigInterface
type gSLBConfigs struct {
	client rest.Interface
	ns     string
}

// newGSLBConfigs returns a GSLBConfigs
func newGSLBConfigs(c *AvilbV1alpha1Client, namespace string) *gSLBConfigs {
	return &gSLBConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gSLBConfig, and returns the corresponding gSLBConfig object, and an error if there is any.
func (c *gSLBConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.GSLBConfig, err error) {
	result = &v1alpha1.GSLBConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gslbconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GSLBConfigs that match those selectors.
func (c *gSLBConfigs) List(opts v1.ListOptions) (result *v1alpha1.GSLBConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GSLBConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gslbconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gSLBConfigs.
func (c *gSLBConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gslbconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gSLBConfig and creates it.  Returns the server's representation of the gSLBConfig, and an error, if there is any.
func (c *gSLBConfigs) Create(gSLBConfig *v1alpha1.GSLBConfig) (result *v1alpha1.GSLBConfig, err error) {
	result = &v1alpha1.GSLBConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gslbconfigs").
		Body(gSLBConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gSLBConfig and updates it. Returns the server's representation of the gSLBConfig, and an error, if there is any.
func (c *gSLBConfigs) Update(gSLBConfig *v1alpha1.GSLBConfig) (result *v1alpha1.GSLBConfig, err error) {
	result = &v1alpha1.GSLBConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gslbconfigs").
		Name(gSLBConfig.Name).
		Body(gSLBConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the gSLBConfig and deletes it. Returns an error if one occurs.
func (c *gSLBConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gslbconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gSLBConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gslbconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gSLBConfig.
func (c *gSLBConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBConfig, err error) {
	result = &v1alpha1.GSLBConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gslbconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
