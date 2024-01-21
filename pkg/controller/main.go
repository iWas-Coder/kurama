package main

import (
  "os"
  "sigs.k8s.io/controller-runtime/pkg/healthz"
  "sigs.k8s.io/controller-runtime/pkg/log/zap"
  k8s_runtime "k8s.io/apimachinery/pkg/runtime"
  k8s_ctrl_runtime "sigs.k8s.io/controller-runtime"
  k8s_util_runtime "k8s.io/apimachinery/pkg/util/runtime"
  k8s_clientgo_scheme "k8s.io/client-go/kubernetes/scheme"
  kurama_v1 "github.com/iWas-Coder/kurama/pkg/controller/v1"
  k8s_metrics_server "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
  scheme = k8s_runtime.NewScheme()
  setupLog = k8s_ctrl_runtime.Log.WithName("setup")
)

func init() {
  k8s_util_runtime.Must(k8s_clientgo_scheme.AddToScheme(scheme))
  k8s_util_runtime.Must(kurama_v1.AddToScheme(scheme))
}

func main() {
  enableLeaderElection := false

  k8s_ctrl_runtime.SetLogger(zap.New(zap.UseDevMode(true)))

  mgr, err := k8s_ctrl_runtime.NewManager(k8s_ctrl_runtime.GetConfigOrDie(), k8s_ctrl_runtime.Options{
    Scheme: scheme,
    Metrics: k8s_metrics_server.Options{ BindAddress: ":8080" },
    HealthProbeBindAddress: ":8081",
    LeaderElection: enableLeaderElection,
    LeaderElectionID: "b9715d24.kurama.io",
  })
  if err != nil {
    setupLog.Error(err, "unable to create manager")
    os.Exit(1)
  }

  err = (&KuramaJobReconciler{
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
  err = mgr.Start(k8s_ctrl_runtime.SetupSignalHandler())
  if err != nil {
    setupLog.Error(err, "unable to run manager")
    os.Exit(1)
  }
}
