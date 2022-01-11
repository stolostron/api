// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/stolostron/api/work/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManifestWorkLister helps list ManifestWorks.
// All objects returned here must be treated as read-only.
type ManifestWorkLister interface {
	// List lists all ManifestWorks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ManifestWork, err error)
	// ManifestWorks returns an object that can list and get ManifestWorks.
	ManifestWorks(namespace string) ManifestWorkNamespaceLister
	ManifestWorkListerExpansion
}

// manifestWorkLister implements the ManifestWorkLister interface.
type manifestWorkLister struct {
	indexer cache.Indexer
}

// NewManifestWorkLister returns a new ManifestWorkLister.
func NewManifestWorkLister(indexer cache.Indexer) ManifestWorkLister {
	return &manifestWorkLister{indexer: indexer}
}

// List lists all ManifestWorks in the indexer.
func (s *manifestWorkLister) List(selector labels.Selector) (ret []*v1.ManifestWork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ManifestWork))
	})
	return ret, err
}

// ManifestWorks returns an object that can list and get ManifestWorks.
func (s *manifestWorkLister) ManifestWorks(namespace string) ManifestWorkNamespaceLister {
	return manifestWorkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManifestWorkNamespaceLister helps list and get ManifestWorks.
// All objects returned here must be treated as read-only.
type ManifestWorkNamespaceLister interface {
	// List lists all ManifestWorks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ManifestWork, err error)
	// Get retrieves the ManifestWork from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ManifestWork, error)
	ManifestWorkNamespaceListerExpansion
}

// manifestWorkNamespaceLister implements the ManifestWorkNamespaceLister
// interface.
type manifestWorkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManifestWorks in the indexer for a given namespace.
func (s manifestWorkNamespaceLister) List(selector labels.Selector) (ret []*v1.ManifestWork, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ManifestWork))
	})
	return ret, err
}

// Get retrieves the ManifestWork from the indexer for a given namespace and name.
func (s manifestWorkNamespaceLister) Get(name string) (*v1.ManifestWork, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("manifestwork"), name)
	}
	return obj.(*v1.ManifestWork), nil
}
