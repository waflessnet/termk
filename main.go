package main

import (
	"termk/k8i"
	"termk/kui"
)

func main() {

	k8i.SetPathKubeConfig("kubeconfig.conf")
	//clusters := k8i.GetClusters()
	//kui.ShowListClusters(clusters)
	clusters := k8i.GetClusters()
	kui.ShowListClusters(clusters)
	//kui.InfoK("cluster", "namespace", "pod")
}
