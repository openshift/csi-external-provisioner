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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	apisv1alpha2 "sigs.k8s.io/gateway-api/apis/applyconfiguration/apis/v1alpha2"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// FakeTLSRoutes implements TLSRouteInterface
type FakeTLSRoutes struct {
	Fake *FakeGatewayV1alpha2
	ns   string
}

var tlsroutesResource = v1alpha2.SchemeGroupVersion.WithResource("tlsroutes")

var tlsroutesKind = v1alpha2.SchemeGroupVersion.WithKind("TLSRoute")

// Get takes name of the tLSRoute, and returns the corresponding tLSRoute object, and an error if there is any.
func (c *FakeTLSRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TLSRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tlsroutesResource, c.ns, name), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// List takes label and field selectors, and returns the list of TLSRoutes that match those selectors.
func (c *FakeTLSRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TLSRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tlsroutesResource, tlsroutesKind, c.ns, opts), &v1alpha2.TLSRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.TLSRouteList{ListMeta: obj.(*v1alpha2.TLSRouteList).ListMeta}
	for _, item := range obj.(*v1alpha2.TLSRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tLSRoutes.
func (c *FakeTLSRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tlsroutesResource, c.ns, opts))

}

// Create takes the representation of a tLSRoute and creates it.  Returns the server's representation of the tLSRoute, and an error, if there is any.
func (c *FakeTLSRoutes) Create(ctx context.Context, tLSRoute *v1alpha2.TLSRoute, opts v1.CreateOptions) (result *v1alpha2.TLSRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tlsroutesResource, c.ns, tLSRoute), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// Update takes the representation of a tLSRoute and updates it. Returns the server's representation of the tLSRoute, and an error, if there is any.
func (c *FakeTLSRoutes) Update(ctx context.Context, tLSRoute *v1alpha2.TLSRoute, opts v1.UpdateOptions) (result *v1alpha2.TLSRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tlsroutesResource, c.ns, tLSRoute), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTLSRoutes) UpdateStatus(ctx context.Context, tLSRoute *v1alpha2.TLSRoute, opts v1.UpdateOptions) (*v1alpha2.TLSRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tlsroutesResource, "status", c.ns, tLSRoute), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// Delete takes name of the tLSRoute and deletes it. Returns an error if one occurs.
func (c *FakeTLSRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tlsroutesResource, c.ns, name, opts), &v1alpha2.TLSRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTLSRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tlsroutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.TLSRouteList{})
	return err
}

// Patch applies the patch and returns the patched tLSRoute.
func (c *FakeTLSRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TLSRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tlsroutesResource, c.ns, name, pt, data, subresources...), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied tLSRoute.
func (c *FakeTLSRoutes) Apply(ctx context.Context, tLSRoute *apisv1alpha2.TLSRouteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.TLSRoute, err error) {
	if tLSRoute == nil {
		return nil, fmt.Errorf("tLSRoute provided to Apply must not be nil")
	}
	data, err := json.Marshal(tLSRoute)
	if err != nil {
		return nil, err
	}
	name := tLSRoute.Name
	if name == nil {
		return nil, fmt.Errorf("tLSRoute.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tlsroutesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeTLSRoutes) ApplyStatus(ctx context.Context, tLSRoute *apisv1alpha2.TLSRouteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.TLSRoute, err error) {
	if tLSRoute == nil {
		return nil, fmt.Errorf("tLSRoute provided to Apply must not be nil")
	}
	data, err := json.Marshal(tLSRoute)
	if err != nil {
		return nil, err
	}
	name := tLSRoute.Name
	if name == nil {
		return nil, fmt.Errorf("tLSRoute.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tlsroutesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha2.TLSRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TLSRoute), err
}
