package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"
	"io"
	"strings"
	"bytes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "Path to the kubeconfig file")
	flag.Parse()

	// Load config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// List all pods in all namespaces
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Listing current Pods:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s/%s (Phase: %s)\n", pod.Namespace, pod.Name, pod.Status.Phase)
	}

	// Watch logs of DB logs
	getLogsForMatchingPods(clientset, "p-35v0yfm35f", "p-")

	// Watch for pod changes
	fmt.Println("\nWatching for Pod events:")
	watchPods(clientset)
}

func watchPods(clientset *kubernetes.Clientset) {
	watcher, err := clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	ch := watcher.ResultChan()
	// fmt.Println(len(ch))
	for event := range ch {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			fmt.Println("Unexpected type")
			continue
		}

		switch event.Type {
		case watch.Added:
			fmt.Printf("[ADDED] Pod: %s/%s\n", pod.Namespace, pod.Name)
		case watch.Modified:
			fmt.Printf("[MODIFIED] Pod: %s/%s\n", pod.Namespace, pod.Name)
		case watch.Deleted:
			fmt.Printf("[DELETED] Pod: %s/%s\n", pod.Namespace, pod.Name)
		}
	}

	time.Sleep(1 * time.Second)
}

func getLogsForMatchingPods(clientset *kubernetes.Clientset, namespace string, prefix string) {
    pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        fmt.Printf("Error listing pods: %v\n", err)
        return
    }

    fmt.Printf("\nFetching logs for pods starting with '%s':\n", prefix)
    for _, pod := range pods.Items {
        if strings.HasPrefix(pod.Name, prefix) {
            for _, container := range pod.Spec.Containers {
                fmt.Printf("  --> Getting logs for %s/%s [container: %s]\n", pod.Namespace, pod.Name, container.Name)

                req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &corev1.PodLogOptions{
                    Container: container.Name,
                })

                podLogs, err := req.Stream(context.TODO())
                if err != nil {
                    fmt.Printf("    [ERROR] Could not get logs: %v\n", err)
                    continue
                }
                defer podLogs.Close()

                logBuf := new(bytes.Buffer)
                _, err = io.Copy(logBuf, podLogs)
                if err != nil {
                    fmt.Printf("    [ERROR] Reading logs: %v\n", err)
                    continue
                }

                logContent := logBuf.String()
                logLines := strings.Split(logContent, "\n")

                errorLines := 0
                for _, line := range logLines {
                    if strings.Contains(strings.ToLower(line), "error") {
                        errorLines++
                    }
                }

                fmt.Printf("    [SUMMARY] %d lines, %d error(s)\n", len(logLines), errorLines)
            }
        }
    }
}

