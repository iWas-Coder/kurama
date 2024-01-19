package main

import (
  "os"
  "k8s.io/apimachinery/pkg/runtime"
  ctrl "sigs.k8s.io/controller-runtime"
  "github.com/iWas-Coder/kurama/controllers"
  "sigs.k8s.io/controller-runtime/pkg/healthz"
  "sigs.k8s.io/controller-runtime/pkg/log/zap"
  kurama "github.com/iWas-Coder/kurama/api/v1"
  utilruntime "k8s.io/apimachinery/pkg/util/runtime"
  clientgoscheme "k8s.io/client-go/kubernetes/scheme"
  metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
  scheme = runtime.NewScheme()
  setupLog = ctrl.Log.WithName("setup")
)

func init() {
  utilruntime.Must(clientgoscheme.AddToScheme(scheme))
  utilruntime.Must(kurama.AddToScheme(scheme))
}

func main() {
  enableLeaderElection := false

  ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

  mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
    Scheme: scheme,
    Metrics: metricsserver.Options{ BindAddress: ":8080" },
    HealthProbeBindAddress: ":8081",
    LeaderElection: enableLeaderElection,
    LeaderElectionID: "b9715d24.kurama.io",
  })
  if err != nil {
    setupLog.Error(err, "unable to create manager")
    os.Exit(1)
  }

  err = (&controllers.KuramaJobReconciler{
    Client: mgr.GetClient(),
    Scheme: mgr.GetScheme(),
  }).SetupWithManager(mgr)
  if err != nil {
    setupLog.Error(err, "unable to create controller", "controller", "KuramaJob")
    os.Exit(1)
  }

  err = mgr.AddHealthzCheck("healthz", healthz.Ping)
  if err != nil {
    setupLog.Error(err, "unable to setup health check")
    os.Exit(1)
  }
  err = mgr.AddReadyzCheck("readyz", healthz.Ping)
  if err != nil {
    setupLog.Error(err, "unable to setup ready check")
    os.Exit(1)
  }

  setupLog.Info("starting manager")
  err = mgr.Start(ctrl.SetupSignalHandler())
  if err != nil {
    setupLog.Error(err, "unable to run manager")
    os.Exit(1)
  }
}
