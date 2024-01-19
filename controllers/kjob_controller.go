package controllers

import (
  "context"
  "k8s.io/apimachinery/pkg/runtime"
  "k8s.io/apimachinery/pkg/api/errors"
  ctrl "sigs.k8s.io/controller-runtime"
  "sigs.k8s.io/controller-runtime/pkg/log"
  "sigs.k8s.io/controller-runtime/pkg/client"
  kurama_api_v1 "github.com/iWas-Coder/kurama/api/v1"
)

type KuramaJobReconciler struct {
  client.Client
  Scheme *runtime.Scheme
}

func (self *KuramaJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
  logger := log.FromContext(ctx)

  // Reconcile job to perform
  var kjob kurama_api_v1.KuramaJob
  err := self.Get(ctx, req.NamespacedName, &kjob)
  if err != nil {
    if errors.IsNotFound(err) {
      logger.Info("KuramaJob instance deleted")
      return ctrl.Result{}, nil
    }
    return ctrl.Result{}, err
  }

  logger.Info(
    "KuramaJob instance applied",
    "Steps",
    kjob.Spec.Steps,
  )
  return ctrl.Result{}, nil
}

func (self *KuramaJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
  return ctrl.NewControllerManagedBy(mgr).
    For(&kurama_api_v1.KuramaJob{}).
    Complete(self)
}
