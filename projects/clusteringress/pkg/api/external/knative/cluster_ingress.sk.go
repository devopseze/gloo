// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"log"
	"sort"

	github_com_solo_io_gloo_projects_clusteringress_api_external_knative "github.com/solo-io/gloo/projects/clusteringress/api/external/knative"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewClusterIngress(namespace, name string) *ClusterIngress {
	clusteringress := &ClusterIngress{}
	clusteringress.ClusterIngress.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return clusteringress
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableClusterIngress interface {
	resources.Resource
	Clone() *github_com_solo_io_gloo_projects_clusteringress_api_external_knative.ClusterIngress
}

var _ CloneableClusterIngress = &github_com_solo_io_gloo_projects_clusteringress_api_external_knative.ClusterIngress{}

type ClusterIngress struct {
	github_com_solo_io_gloo_projects_clusteringress_api_external_knative.ClusterIngress
}

func (r *ClusterIngress) Clone() resources.Resource {
	return &ClusterIngress{ClusterIngress: *r.ClusterIngress.Clone()}
}

func (r *ClusterIngress) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	clone := r.ClusterIngress.Clone()
	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})
	err := binary.Write(hasher, binary.LittleEndian, hashutils.HashAll(clone))
	if err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (r *ClusterIngress) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *ClusterIngress) GroupVersionKind() schema.GroupVersionKind {
	return ClusterIngressGVK
}

type ClusterIngressList []*ClusterIngress

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ClusterIngressList) Find(namespace, name string) (*ClusterIngress, error) {
	for _, clusterIngress := range list {
		if clusterIngress.GetMetadata().Name == name {
			if namespace == "" || clusterIngress.GetMetadata().Namespace == namespace {
				return clusterIngress, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find clusterIngress %v.%v", namespace, name)
}

func (list ClusterIngressList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, clusterIngress := range list {
		ress = append(ress, clusterIngress)
	}
	return ress
}

func (list ClusterIngressList) Names() []string {
	var names []string
	for _, clusterIngress := range list {
		names = append(names, clusterIngress.GetMetadata().Name)
	}
	return names
}

func (list ClusterIngressList) NamespacesDotNames() []string {
	var names []string
	for _, clusterIngress := range list {
		names = append(names, clusterIngress.GetMetadata().Namespace+"."+clusterIngress.GetMetadata().Name)
	}
	return names
}

func (list ClusterIngressList) Sort() ClusterIngressList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ClusterIngressList) Clone() ClusterIngressList {
	var clusterIngressList ClusterIngressList
	for _, clusterIngress := range list {
		clusterIngressList = append(clusterIngressList, resources.Clone(clusterIngress).(*ClusterIngress))
	}
	return clusterIngressList
}

func (list ClusterIngressList) Each(f func(element *ClusterIngress)) {
	for _, clusterIngress := range list {
		f(clusterIngress)
	}
}

func (list ClusterIngressList) EachResource(f func(element resources.Resource)) {
	for _, clusterIngress := range list {
		f(clusterIngress)
	}
}

func (list ClusterIngressList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *ClusterIngress) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var (
	ClusterIngressGVK = schema.GroupVersionKind{
		Version: "v1alpha1",
		Group:   "networking.internal.knative.dev",
		Kind:    "ClusterIngress",
	}
)
