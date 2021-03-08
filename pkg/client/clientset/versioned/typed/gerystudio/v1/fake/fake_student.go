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

	gerystudiov1 "github.com/ankyhe/k8s-customize-controller/pkg/apis/gerystudio/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStudents implements StudentInterface
type FakeStudents struct {
	Fake *FakeGerystudioV1
	ns   string
}

var studentsResource = schema.GroupVersionResource{Group: "gerystudio.k8s.io", Version: "v1", Resource: "students"}

var studentsKind = schema.GroupVersionKind{Group: "gerystudio.k8s.io", Version: "v1", Kind: "Student"}

// Get takes name of the student, and returns the corresponding student object, and an error if there is any.
func (c *FakeStudents) Get(ctx context.Context, name string, options v1.GetOptions) (result *gerystudiov1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(studentsResource, c.ns, name), &gerystudiov1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gerystudiov1.Student), err
}

// List takes label and field selectors, and returns the list of Students that match those selectors.
func (c *FakeStudents) List(ctx context.Context, opts v1.ListOptions) (result *gerystudiov1.StudentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(studentsResource, studentsKind, c.ns, opts), &gerystudiov1.StudentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gerystudiov1.StudentList{ListMeta: obj.(*gerystudiov1.StudentList).ListMeta}
	for _, item := range obj.(*gerystudiov1.StudentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested students.
func (c *FakeStudents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(studentsResource, c.ns, opts))

}

// Create takes the representation of a student and creates it.  Returns the server's representation of the student, and an error, if there is any.
func (c *FakeStudents) Create(ctx context.Context, student *gerystudiov1.Student, opts v1.CreateOptions) (result *gerystudiov1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(studentsResource, c.ns, student), &gerystudiov1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gerystudiov1.Student), err
}

// Update takes the representation of a student and updates it. Returns the server's representation of the student, and an error, if there is any.
func (c *FakeStudents) Update(ctx context.Context, student *gerystudiov1.Student, opts v1.UpdateOptions) (result *gerystudiov1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(studentsResource, c.ns, student), &gerystudiov1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gerystudiov1.Student), err
}

// Delete takes name of the student and deletes it. Returns an error if one occurs.
func (c *FakeStudents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(studentsResource, c.ns, name), &gerystudiov1.Student{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStudents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(studentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &gerystudiov1.StudentList{})
	return err
}

// Patch applies the patch and returns the patched student.
func (c *FakeStudents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gerystudiov1.Student, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(studentsResource, c.ns, name, pt, data, subresources...), &gerystudiov1.Student{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gerystudiov1.Student), err
}
