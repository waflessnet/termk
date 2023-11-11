package k8i

import (
	"bytes"
	"context"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sort"
	"strings"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

var kubeConfigPath = ""

func SetClient(context, pathKube string) *kubernetes.Clientset {

	//kubeconfig := ""
	// use the current context in kubeconfig
	config, err := buildConfigFromFlags(context, pathKube)
	if err != nil {
		panic(err.Error())
	}

	// create the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return client
}
func GetPathKubeConfig() string {
	return kubeConfigPath
}
func SetPathKubeConfig(kp string) {
	kubeConfigPath = kp
}
func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func GetNameSpaces(client *kubernetes.Clientset) []string {

	namespaces, _ := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	var n []string
	for _, c := range namespaces.Items {
		n = append(n, c.Name)
	}
	sort.Strings(n)
	return n
}

func GetClusters() []string {
	file := GetPathKubeConfig()
	la, _ := clientcmd.LoadFromFile(file)
	var clusters []string

	for _, c := range la.Contexts {
		if strings.Trim(c.Cluster, "") != "" {
			clusters = append(clusters, c.Cluster)
		}
	}

	sort.Strings(clusters)
	return clusters
}

func GetPods(client *kubernetes.Clientset, namespace string) []string {
	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var p []string
	for _, c := range pods.Items {
		p = append(p, c.Name+"  ->  "+string(c.Status.Phase)+" ->  "+c.Status.PodIP+" -> "+c.Status.StartTime.String())
		//p = append(p, c.Name)
		// println(c.Name + "  ->  " + string(c.Status.Phase) + " ->  " + c.Status.PodIP)
	}
	sort.Strings(p)
	return p
}

func GetLogPod(client *kubernetes.Clientset, namespace string, podName string) string {

	podLogOpts := v1.PodLogOptions{Timestamps: true}
	req := client.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "error in opening stream"
	}
	defer func(podLogs io.ReadCloser) {
		err := podLogs.Close()
		if err != nil {

		}
	}(podLogs)

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf"
	}
	str := buf.String()

	return str
}

func DeletePod(client *kubernetes.Clientset, namespace string, podName string) error {
	ctx := context.Background()
	deletePolicy := metav1.DeletePropagationForeground
	return client.CoreV1().Pods(namespace).Delete(ctx, podName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
func DescribePod(client *kubernetes.Clientset, namespace string, podName string) {
	//client.CoreV1().Pods(namespace).
	// future.

}
