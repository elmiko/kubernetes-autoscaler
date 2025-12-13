/*
Copyright 2025 The Kubernetes Authors.

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

package simulation

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/klog/v2"
)

type simulationProvider struct {
	cloudprovider.CloudProvider
}

func NewSimulationProvider(provider cloudprovider.CloudProvider) cloudprovider.CloudProvider {
	return &simulationProvider{
		CloudProvider: provider,
	}
}

// NodeGroups returns all node groups configured for this cloud provider.
func (p simulationProvider) NodeGroups() []cloudprovider.NodeGroup {
	nodegroups := []cloudprovider.NodeGroup{}
	for _, ng := range p.CloudProvider.NodeGroups() {
		nodegroups = append(nodegroups, newSimulationNodeGroup(ng))
	}

	return nodegroups
}

// NodeGroupForNode returns the node group for the given node, nil if the node
// should not be processed by cluster autoscaler, or non-nil error if such
// occurred. Must be implemented.
func (p simulationProvider) NodeGroupForNode(node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	ng, err := p.CloudProvider.NodeGroupForNode(node)
	if err != nil {
		return nil, err
	}
	if ng == nil {
		return nil, nil
	}

	return newSimulationNodeGroup(ng), nil
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
// Implementation optional.
func (p simulationProvider) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string, taints []apiv1.Taint, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	klog.V(0).Infof("Simulation: not creating new node group for machine type %q", machineType)
	return nil, cloudprovider.ErrNotImplemented
}
