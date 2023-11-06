package main

import (
	"termk/k8i"
	"termk/ui"
)

func main() {

	//client := k8i.SetClient(pathKube)
	//k8i.GetNameSpaces(client)
	k8i.SetPathKubeConfig("kubeconfig.conf")
	clusters := k8i.GetClusters()
	ui.ShowListClusters(clusters)
}
