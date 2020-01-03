/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package internalversion

import (
	time "time"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	clientsetinternalversion "github.com/gardener/machine-controller-manager/pkg/client/clientset/internalversion"
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/machine-controller-manager/pkg/client/listers/machine/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OpenStackMachineClassInformer provides access to a shared informer and lister for
// OpenStackMachineClasses.
type OpenStackMachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.OpenStackMachineClassLister
}

type openStackMachineClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOpenStackMachineClassInformer constructs a new informer for OpenStackMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenStackMachineClassInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenStackMachineClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOpenStackMachineClassInformer constructs a new informer for OpenStackMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenStackMachineClassInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Machine().OpenStackMachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Machine().OpenStackMachineClasses(namespace).Watch(options)
			},
		},
		&machine.OpenStackMachineClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *openStackMachineClassInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenStackMachineClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openStackMachineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine.OpenStackMachineClass{}, f.defaultInformer)
}

func (f *openStackMachineClassInformer) Lister() internalversion.OpenStackMachineClassLister {
	return internalversion.NewOpenStackMachineClassLister(f.Informer().GetIndexer())
}
