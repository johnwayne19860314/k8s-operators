package v1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JohnFooSpec defines the desired state of Foo
type JohnFooSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
}

// JohnFooStatus defines the observed state of JohnFoo.
// It should always be reconstructable from the state of the cluster and/or outside world.
type JohnFooStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JohnFoo is the Schema for the johnfooes API
// +k8s:openapi-gen=true
type JohnFoo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JohnFooSpec   `json:"spec,omitempty"`
	Status JohnFooStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JohnFooList contains a list of JohnFoo
type JohnFooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JohnFoo `json:"items"`
}

func init(){
	Schema.AddKnownTypes(GroupVersion, &JohnFoo{}, &JohnFooList{})
}	