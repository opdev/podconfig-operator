/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Vlan struct {
	ParentInterfaceName string `json:"parentInterfaceName,omitempty"`
	VlanID              int16  `json:"vlanID,omitempty"`
	BridgeName          string `json:"bridgeName,omitempty"`
}

// TrunkPortSpec defines the desired state of TrunkPort
type TrunkPortSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Vlans []Vlan `json:"foo,omitempty"`
}

// TrunkPortStatus defines the observed state of TrunkPort
type TrunkPortStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TrunkPort is the Schema for the trunkports API
type TrunkPort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrunkPortSpec   `json:"spec,omitempty"`
	Status TrunkPortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrunkPortList contains a list of TrunkPort
type TrunkPortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrunkPort `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrunkPort{}, &TrunkPortList{})
}