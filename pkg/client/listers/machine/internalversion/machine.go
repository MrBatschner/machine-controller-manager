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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineLister helps list Machines.
type MachineLister interface {
	// List lists all Machines in the indexer.
	List(selector labels.Selector) (ret []*machine.Machine, err error)
	// Machines returns an object that can list and get Machines.
	Machines(namespace string) MachineNamespaceLister
	MachineListerExpansion
}

// machineLister implements the MachineLister interface.
type machineLister struct {
	indexer cache.Indexer
}

// NewMachineLister returns a new MachineLister.
func NewMachineLister(indexer cache.Indexer) MachineLister {
	return &machineLister{indexer: indexer}
}

// List lists all Machines in the indexer.
func (s *machineLister) List(selector labels.Selector) (ret []*machine.Machine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.Machine))
	})
	return ret, err
}

// Machines returns an object that can list and get Machines.
func (s *machineLister) Machines(namespace string) MachineNamespaceLister {
	return machineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineNamespaceLister helps list and get Machines.
type MachineNamespaceLister interface {
	// List lists all Machines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*machine.Machine, err error)
	// Get retrieves the Machine from the indexer for a given namespace and name.
	Get(name string) (*machine.Machine, error)
	MachineNamespaceListerExpansion
}

// machineNamespaceLister implements the MachineNamespaceLister
// interface.
type machineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Machines in the indexer for a given namespace.
func (s machineNamespaceLister) List(selector labels.Selector) (ret []*machine.Machine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.Machine))
	})
	return ret, err
}

// Get retrieves the Machine from the indexer for a given namespace and name.
func (s machineNamespaceLister) Get(name string) (*machine.Machine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("machine"), name)
	}
	return obj.(*machine.Machine), nil
}
