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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/rancher/helm-controller/pkg/apis/helm.cattle.io/v1"
	clientset "github.com/rancher/helm-controller/pkg/generated/clientset/versioned/typed/helm.cattle.io/v1"
	informers "github.com/rancher/helm-controller/pkg/generated/informers/externalversions/helm.cattle.io/v1"
	listers "github.com/rancher/helm-controller/pkg/generated/listers/helm.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type HelmChartHandler func(string, *v1.HelmChart) (*v1.HelmChart, error)

type HelmChartController interface {
	HelmChartClient

	OnChange(ctx context.Context, name string, sync HelmChartHandler)
	OnRemove(ctx context.Context, name string, sync HelmChartHandler)
	Enqueue(namespace, name string)

	Cache() HelmChartCache

	Informer() cache.SharedIndexInformer
	GroupVersionKind() schema.GroupVersionKind

	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
	AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler)
	Updater() generic.Updater
}

type HelmChartClient interface {
	Create(*v1.HelmChart) (*v1.HelmChart, error)
	Update(*v1.HelmChart) (*v1.HelmChart, error)
	UpdateStatus(*v1.HelmChart) (*v1.HelmChart, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChart, error)
	List(namespace string, opts metav1.ListOptions) (*v1.HelmChartList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelmChart, err error)
}

type HelmChartCache interface {
	Get(namespace, name string) (*v1.HelmChart, error)
	List(namespace string, selector labels.Selector) ([]*v1.HelmChart, error)

	AddIndexer(indexName string, indexer HelmChartIndexer)
	GetByIndex(indexName, key string) ([]*v1.HelmChart, error)
}

type HelmChartIndexer func(obj *v1.HelmChart) ([]string, error)

type helmChartController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.HelmChartsGetter
	informer          informers.HelmChartInformer
	gvk               schema.GroupVersionKind
}

func NewHelmChartController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.HelmChartsGetter, informer informers.HelmChartInformer) HelmChartController {
	return &helmChartController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromHelmChartHandlerToHandler(sync HelmChartHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.HelmChart
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.HelmChart))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *helmChartController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.HelmChart))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateHelmChartOnChange(updater generic.Updater, handler HelmChartHandler) HelmChartHandler {
	return func(key string, obj *v1.HelmChart) (*v1.HelmChart, error) {
		if obj == nil {
			return handler(key, nil)
		}

		copyObj := obj.DeepCopy()
		newObj, err := handler(key, copyObj)
		if newObj != nil {
			copyObj = newObj
		}
		if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
			newObj, err := updater(copyObj)
			if newObj != nil && err == nil {
				copyObj = newObj.(*v1.HelmChart)
			}
		}

		return copyObj, err
	}
}

func (c *helmChartController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *helmChartController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *helmChartController) OnChange(ctx context.Context, name string, sync HelmChartHandler) {
	c.AddGenericHandler(ctx, name, FromHelmChartHandlerToHandler(sync))
}

func (c *helmChartController) OnRemove(ctx context.Context, name string, sync HelmChartHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromHelmChartHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *helmChartController) Enqueue(namespace, name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), namespace, name)
}

func (c *helmChartController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *helmChartController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *helmChartController) Cache() HelmChartCache {
	return &helmChartCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *helmChartController) Create(obj *v1.HelmChart) (*v1.HelmChart, error) {
	return c.clientGetter.HelmCharts(obj.Namespace).Create(obj)
}

func (c *helmChartController) Update(obj *v1.HelmChart) (*v1.HelmChart, error) {
	return c.clientGetter.HelmCharts(obj.Namespace).Update(obj)
}

func (c *helmChartController) UpdateStatus(obj *v1.HelmChart) (*v1.HelmChart, error) {
	return c.clientGetter.HelmCharts(obj.Namespace).UpdateStatus(obj)
}

func (c *helmChartController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.HelmCharts(namespace).Delete(name, options)
}

func (c *helmChartController) Get(namespace, name string, options metav1.GetOptions) (*v1.HelmChart, error) {
	return c.clientGetter.HelmCharts(namespace).Get(name, options)
}

func (c *helmChartController) List(namespace string, opts metav1.ListOptions) (*v1.HelmChartList, error) {
	return c.clientGetter.HelmCharts(namespace).List(opts)
}

func (c *helmChartController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.HelmCharts(namespace).Watch(opts)
}

func (c *helmChartController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelmChart, err error) {
	return c.clientGetter.HelmCharts(namespace).Patch(name, pt, data, subresources...)
}

type helmChartCache struct {
	lister  listers.HelmChartLister
	indexer cache.Indexer
}

func (c *helmChartCache) Get(namespace, name string) (*v1.HelmChart, error) {
	return c.lister.HelmCharts(namespace).Get(name)
}

func (c *helmChartCache) List(namespace string, selector labels.Selector) ([]*v1.HelmChart, error) {
	return c.lister.HelmCharts(namespace).List(selector)
}

func (c *helmChartCache) AddIndexer(indexName string, indexer HelmChartIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.HelmChart))
		},
	}))
}

func (c *helmChartCache) GetByIndex(indexName, key string) (result []*v1.HelmChart, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.HelmChart))
	}
	return result, nil
}
