/*
Copyright 2018 the Heptio Ark contributors.

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

package v1

import (
	time "time"

	arkv1 "github.com/heptio/ark/pkg/apis/ark/v1"
	versioned "github.com/heptio/ark/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/heptio/ark/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/heptio/ark/pkg/generated/listers/ark/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScheduleInformer provides access to a shared informer and lister for
// Schedules.
type ScheduleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ScheduleLister
}

type scheduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduleInformer constructs a new informer for Schedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduleInformer constructs a new informer for Schedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArkV1().Schedules(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArkV1().Schedules(namespace).Watch(options)
			},
		},
		&arkv1.Schedule{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&arkv1.Schedule{}, f.defaultInformer)
}

func (f *scheduleInformer) Lister() v1.ScheduleLister {
	return v1.NewScheduleLister(f.Informer().GetIndexer())
}
