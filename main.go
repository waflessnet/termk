package main

import (
	"flag"
	"termk/k8i"
	"termk/kui"
)

func main() {
	var kconfig string
	flag.StringVar(&kconfig, "kconfig", "kubeconfig.conf", "set path config k8")
	flag.Parse()
	k8i.SetPathKubeConfig(kconfig)
	clusters := k8i.GetClusters()
	kui.ShowListClusters(clusters)
}
