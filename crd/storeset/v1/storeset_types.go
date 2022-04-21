package v1

import (
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LocalPvSpec struct {
	NodeAffinity      *v1.VolumeNodeAffinity `json:"nodeAffinity,omitempty"`
	Capacity          resource.Quantity      `json:"capacity"`
	LocalVolumeSource *v1.LocalVolumeSource  `json:"source,omitempty"`
}

type StoreStatefulSetSpec struct {
	Replicas *int32 `json:"replicas,omitempty"`
	Image    string `json:"image,omitempty"`
}

type PublisherDeptSpec struct {
	Replicas *int32 `json:"replicas,omitempty"`
	Image    string `json:"image,omitempty"`
}

// StoreSetSpec defines the desired state of StoreSet
type StoreSetSpec struct {
	Volume    LocalPvSpec          `json:"volume,omitempty"`
	Store     StoreStatefulSetSpec `json:"store,omitempty"`
	Publisher PublisherDeptSpec    `json:"publisher,omitempty"`
}

type LocalPvStatus struct {
	Name   string                    `json:"name"`
	Status v1.PersistentVolumeStatus `json:",inline"`
}

type StoreStatefulSetStatus struct {
	WorkloadName string                `json:"workloadName"`
	ServiceName  string                `json:"serviceName"`
	Workload     v12.StatefulSetStatus `json:"workload"`
	Service      v1.ServiceStatus      `json:"service"`
}

type PublisherDeptStatus struct {
	Name   string               `json:"name"`
	Status v12.DeploymentStatus `json:",inline"`
}

// StoreSetStatus defines the observed state of StoreSet
type StoreSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	VolumeStatus LocalPvStatus          `json:"volume"`
	StoreStatus  StoreStatefulSetStatus `json:"Store"`
	Ready        bool                   `json:"ready"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoreSetSpec   `json:"spec,omitempty"`
	Status StoreSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type StoreSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoreSet `json:"items"`
}
