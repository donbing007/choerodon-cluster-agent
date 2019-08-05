// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/choerodon/choerodon-cluster-agent/pkg/client/clientset/versioned"
	choerodonv1alpha1 "github.com/choerodon/choerodon-cluster-agent/pkg/client/clientset/versioned/typed/choerodon/v1alpha1"
	fakechoerodonv1alpha1 "github.com/choerodon/choerodon-cluster-agent/pkg/client/clientset/versioned/typed/choerodon/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ChoerodonV1alpha1 retrieves the ChoerodonV1alpha1Client
func (c *Clientset) ChoerodonV1alpha1() choerodonv1alpha1.ChoerodonV1alpha1Interface {
	return &fakechoerodonv1alpha1.FakeChoerodonV1alpha1{Fake: &c.Fake}
}

// Choerodon retrieves the ChoerodonV1alpha1Client
func (c *Clientset) Choerodon() choerodonv1alpha1.ChoerodonV1alpha1Interface {
	return &fakechoerodonv1alpha1.FakeChoerodonV1alpha1{Fake: &c.Fake}
}