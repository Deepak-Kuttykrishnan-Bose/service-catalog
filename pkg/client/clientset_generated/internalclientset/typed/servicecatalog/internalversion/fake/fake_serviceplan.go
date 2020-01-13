/*
Copyright 2020 The Kubernetes Authors.

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
	servicecatalog "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServicePlans implements ServicePlanInterface
type FakeServicePlans struct {
	Fake *FakeServicecatalog
	ns   string
}

var serviceplansResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "serviceplans"}

var serviceplansKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ServicePlan"}

// Get takes name of the servicePlan, and returns the corresponding servicePlan object, and an error if there is any.
func (c *FakeServicePlans) Get(name string, options v1.GetOptions) (result *servicecatalog.ServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceplansResource, c.ns, name), &servicecatalog.ServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}

// List takes label and field selectors, and returns the list of ServicePlans that match those selectors.
func (c *FakeServicePlans) List(opts v1.ListOptions) (result *servicecatalog.ServicePlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceplansResource, serviceplansKind, c.ns, opts), &servicecatalog.ServicePlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ServicePlanList{ListMeta: obj.(*servicecatalog.ServicePlanList).ListMeta}
	for _, item := range obj.(*servicecatalog.ServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicePlans.
func (c *FakeServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceplansResource, c.ns, opts))

}

// Create takes the representation of a servicePlan and creates it.  Returns the server's representation of the servicePlan, and an error, if there is any.
func (c *FakeServicePlans) Create(servicePlan *servicecatalog.ServicePlan) (result *servicecatalog.ServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceplansResource, c.ns, servicePlan), &servicecatalog.ServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}

// Update takes the representation of a servicePlan and updates it. Returns the server's representation of the servicePlan, and an error, if there is any.
func (c *FakeServicePlans) Update(servicePlan *servicecatalog.ServicePlan) (result *servicecatalog.ServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceplansResource, c.ns, servicePlan), &servicecatalog.ServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicePlans) UpdateStatus(servicePlan *servicecatalog.ServicePlan) (*servicecatalog.ServicePlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceplansResource, "status", c.ns, servicePlan), &servicecatalog.ServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}

// Delete takes name of the servicePlan and deletes it. Returns an error if one occurs.
func (c *FakeServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceplansResource, c.ns, name), &servicecatalog.ServicePlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceplansResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &servicecatalog.ServicePlanList{})
	return err
}

// Patch applies the patch and returns the patched servicePlan.
func (c *FakeServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceplansResource, c.ns, name, pt, data, subresources...), &servicecatalog.ServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
