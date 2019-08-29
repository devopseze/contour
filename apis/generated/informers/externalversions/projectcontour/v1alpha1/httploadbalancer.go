/*
Copyright 2019 VMware

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/heptio/contour/apis/generated/clientset/versioned"
	internalinterfaces "github.com/heptio/contour/apis/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/heptio/contour/apis/generated/listers/projectcontour/v1alpha1"
	projectcontourv1alpha1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HTTPLoadBalancerInformer provides access to a shared informer and lister for
// HTTPLoadBalancers.
type HTTPLoadBalancerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HTTPLoadBalancerLister
}

type hTTPLoadBalancerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPLoadBalancerInformer constructs a new informer for HTTPLoadBalancer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPLoadBalancerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPLoadBalancerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPLoadBalancerInformer constructs a new informer for HTTPLoadBalancer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPLoadBalancerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcontourV1alpha1().HTTPLoadBalancers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcontourV1alpha1().HTTPLoadBalancers(namespace).Watch(options)
			},
		},
		&projectcontourv1alpha1.HTTPLoadBalancer{},
		resyncPeriod,
		indexers,
	)
}

func (f *hTTPLoadBalancerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPLoadBalancerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hTTPLoadBalancerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcontourv1alpha1.HTTPLoadBalancer{}, f.defaultInformer)
}

func (f *hTTPLoadBalancerInformer) Lister() v1alpha1.HTTPLoadBalancerLister {
	return v1alpha1.NewHTTPLoadBalancerLister(f.Informer().GetIndexer())
}
