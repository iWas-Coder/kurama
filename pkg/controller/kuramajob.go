package main

import (
  "context"
  k8s_runtime "k8s.io/apimachinery/pkg/runtime"
  k8s_errors "k8s.io/apimachinery/pkg/api/errors"
  k8s_ctrl_runtime "sigs.k8s.io/controller-runtime"
  k8s_log "sigs.k8s.io/controller-runtime/pkg/log"
  k8s_client "sigs.k8s.io/controller-runtime/pkg/client"
  kurama_v1 "github.com/iWas-Coder/kurama/pkg/controller/v1"
)

type KuramaJobReconciler struct {
  k8s_client.Client
  Scheme *k8s_runtime.Scheme
}

func (self *KuramaJobReconciler) Reconcile(ctx context.Context, req k8s_ctrl_runtime.Request) (k8s_ctrl_runtime.Result, error) {
  logger := k8s_log.FromContext(ctx)

  // Reconcile action to perform
  var kjob kurama_v1.KuramaJob
  err := self.Get(ctx, req.NamespacedName, &kjob)
  if err != nil {
    if k8s_errors.IsNotFound(err) {
      logger.Info("KuramaJob instance deleted")
      return k8s_ctrl_runtime.Result{}, nil
    }
    return k8s_ctrl_runtime.Result{}, err
  }

  logger.Info(
    "KuramaJob instance applied",
    "Steps",
    kjob.Spec.Steps,
  )
  return k8s_ctrl_runtime.Result{}, nil
}

func (self *KuramaJobReconciler) SetupWithManager(mgr k8s_ctrl_runtime.Manager) error {
  return k8s_ctrl_runtime.NewControllerManagedBy(mgr).
    For(&kurama_v1.KuramaJob{}).
    Complete(self)
}
