// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/choerodon/choerodon-cluster-agent/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// C7NHelmReleases returns a C7NHelmReleaseInformer.
	C7NHelmReleases() C7NHelmReleaseInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// C7NHelmReleases returns a C7NHelmReleaseInformer.
func (v *version) C7NHelmReleases() C7NHelmReleaseInformer {
	return &c7NHelmReleaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
