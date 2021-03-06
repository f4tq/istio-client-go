package v1alpha3
/*
Portions Copyright 2017 The Kubernetes Authors.
Portions Copyright 2018 Adobe.
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

import (
	"github.com/golang/protobuf/proto"
	istiov1alpha3 "istio.io/api/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualService is a Istio VirtualService resource
type ServiceEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceEntrySpec `json:"spec"`
}

func (se *ServiceEntry) GetSpecMessage() proto.Message {
	return &se.Spec.ServiceEntry
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualServiceList is a list of VirtualService resources
type ServiceEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ServiceEntry `json:"items"`
}

// VirtualServiceSpec is a wrapper around Istio VirtualService
type ServiceEntrySpec struct {
	istiov1alpha3.ServiceEntry
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *ServiceEntrySpec) DeepCopyInto(out *ServiceEntrySpec) {
	*out = *in
}
