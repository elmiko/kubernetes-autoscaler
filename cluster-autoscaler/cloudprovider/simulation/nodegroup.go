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
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/klog/v2"
)

type simulationNodeGroup struct {
	cloudprovider.NodeGroup
}

func newSimulationNodeGroup(nodegroup cloudprovider.NodeGroup) cloudprovider.NodeGroup {
	return &simulationNodeGroup{
		NodeGroup: nodegroup,
	}
}

// IncreaseSize increases the size of the node group. To delete a node you need
// to explicitly name it and use DeleteNode. This function should wait until
// node group size is updated. Implementation required.
func (ng simulationNodeGroup) IncreaseSize(delta int) error {
	klog.V(0).Infof("Simulation: skipping IncreaseSize for node group %q", ng.Id())
	return nil
}

// AtomicIncreaseSize tries to increase the size of the node group atomically.
// It returns error if requesting the entire delta fails. The method doesn't wait until the new instances appear.
// Implementation is optional. Implementation of this method generally requires external cloud provider support
// for atomically requesting multiple instances. If implemented, CA will take advantage of the method while scaling up
// BestEffortAtomicScaleUp ProvisioningClass, guaranteeing that all instances required for such a
// ProvisioningRequest are provisioned atomically.
func (ng simulationNodeGroup) AtomicIncreaseSize(delta int) error {
	klog.V(0).Infof("Simulation: skipping AtomicIncreaseSize for node group %q", ng.Id())
	return nil
}

// DeleteNodes deletes nodes from this node group. Error is returned either on
// failure or if the given node doesn't belong to this node group. This function
// should wait until node group size is updated. Implementation required.
func (ng simulationNodeGroup) DeleteNodes([]*apiv1.Node) error {
	klog.V(0).Infof("Simulation: skipping DeleteNodes for node group %q", ng.Id())
	return nil
}

// ForceDeleteNodes deletes nodes from this node group, without checking for
// constraints like minimal size validation etc. Error is returned either on
// failure or if the given node doesn't belong to this node group. This function
// should wait until node group size is updated.
func (ng simulationNodeGroup) ForceDeleteNodes([]*apiv1.Node) error {
	klog.V(0).Infof("Simulation: skipping ForceDeleteNodes for node group %q", ng.Id())
	return nil
}

// DecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the
// request for new nodes that have not been yet fulfilled. Delta should be negative.
// It is assumed that cloud provider will not delete the existing nodes when there
// is an option to just decrease the target. Implementation required.
func (ng simulationNodeGroup) DecreaseTargetSize(delta int) error {
	klog.V(0).Infof("Simulation: skipping DecreaseTargetSize for node group %q", ng.Id())
	return nil
}

// Create creates the node group on the cloud provider side. Implementation optional.
func (ng simulationNodeGroup) Create() (cloudprovider.NodeGroup, error) {
	klog.V(0).Infof("Simulation: skipping Create for node group %q", ng.Id())
	return nil, cloudprovider.ErrNotImplemented
}

// Delete deletes the node group on the cloud provider side.
// This will be executed only for autoprovisioned node groups, once their size drops to 0.
// Implementation optional.
func (ng simulationNodeGroup) Delete() error {
	klog.V(0).Infof("Simulation: skipping Delete for node group %q", ng.Id())
	return cloudprovider.ErrNotImplemented
}
