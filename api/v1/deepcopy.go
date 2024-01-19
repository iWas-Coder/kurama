package v1

import (
  runtime "k8s.io/apimachinery/pkg/runtime"
)

func (self *KuramaJobSpec) DeepCopyInto(out *KuramaJobSpec) {
  *out = *self
}

func (self *KuramaJobSpec) DeepCopy() *KuramaJobSpec {
  if self == nil {
    return nil
  }
  out := new(KuramaJobSpec)
  self.DeepCopyInto(out)
  return out
}

func (self *KuramaJobStatus) DeepCopyInto(out *KuramaJobStatus) {
  *out = *self
}

func (self *KuramaJobStatus) DeepCopy() *KuramaJobStatus {
  if self == nil {
    return nil
  }
  out := new(KuramaJobStatus)
  self.DeepCopyInto(out)
  return out
}

func (self *KuramaJob) DeepCopyInto(out *KuramaJob) {
  *out = *self
  out.TypeMeta = self.TypeMeta
  self.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
  out.Spec = self.Spec
  out.Status = self.Status
}

func (self *KuramaJob) DeepCopy() *KuramaJob {
  if self == nil {
    return nil
  }
  out := new(KuramaJob)
  self.DeepCopyInto(out)
  return out
}

func (self *KuramaJob) DeepCopyObject() runtime.Object {
  c := self.DeepCopy()
  if c != nil {
    return c
  }
  return nil
}

func (self *KuramaJobList) DeepCopyInto(out *KuramaJobList) {
  *out = *self
  out.TypeMeta = self.TypeMeta
  self.ListMeta.DeepCopyInto(&out.ListMeta)
  if self.Items != nil {
    i, j := &self.Items, &out.Items
    *j = make([]KuramaJob, len(*i))
    for k := range *i {
      (*i)[k].DeepCopyInto(&(*j)[k])
    }
  }
}

func (self *KuramaJobList) DeepCopy() *KuramaJobList {
  if self == nil {
    return nil
  }
  out := new(KuramaJobList)
  self.DeepCopyInto(out)
  return out
}

func (self *KuramaJobList) DeepCopyObject() runtime.Object {
  c := self.DeepCopy()
  if c != nil {
    return c
  }
  return nil
}
