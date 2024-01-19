package v1

import (
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Step struct {
  Name    string `json:"name,omitempty"`
  Command string `json:"command,omitempty"`
}

type KuramaJobSpec struct {
  Steps []Step `json:"steps,omitempty"`
}

type KuramaJobStatus struct {}

type KuramaJob struct {
  metav1.TypeMeta        `json:",inline"`
  metav1.ObjectMeta      `json:"metadata,omitempty"`
  Spec   KuramaJobSpec   `json:"spec,omitempty"`
  Status KuramaJobStatus `json:"status,omitempty"`
}

type KuramaJobList struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ListMeta   `json:"metadata,omitempty"`
  Items []KuramaJob `json:"items"`
}

func init() {
  SchemeBuilder.Register(&KuramaJob{}, &KuramaJobList{})
}
