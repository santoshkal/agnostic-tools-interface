package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var kubeconfig string

// func init() {
// 	// kubeconfig file parsing
// 	flag.StringVar(&kubeconfig, "kubeconfig", "", "/home/a0557.kube/config")
// 	flag.Parse()
// }

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// create the config object from kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error creating config: %v\n", err)
	}

	// create clientset (set of muliple clients) for each Group (e.g. Core),
	// the Version (V1) of Group and Kind (e.g. Pods) so GVK.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creatingclientset: %v", err)
	}

	// ctx := context.Background()
	// executes GET request to K8s API to get pods 'cart' from 'prepayment' namespace
	// pod, err := clientset.CoreV1().Pods("prepayment").Get(ctx, "cart", metav1.GetOptions{})

	// openAPIURL := clientset.CoreV1().RESTClient().Post().AbsPath("/openapi/v2").URL()
	urlscheme, err := clientset.Discovery().OpenAPIV3().Paths()
	// Create a request to fetch OpenAPI specs
	// req, err := http.Get(openAPIURL.Path)
	if err != nil {
		fmt.Printf("Error fetching oprnapiv3 request to fetch OpenAPI specs: %v", err)
	}

	jsonData, err := json.Marshal(urlscheme)
	if err != nil {
		fmt.Printf("Error marshalling json: %v", err)
	}
	fmt.Println(string(jsonData))
	// // Perform the request
	// // transport, err := rest.TransportFor(config)
	// // if err != nil {
	// // 	fmt.Printf("Error creatinga transport to fetch OpenAPI specs: %v", err)
	// // }

	// roundTripper := &http.Transport{
	// 	Proxy:                 http.ProxyFromEnvironment,
	// 	IdleConnTimeout:       config.Timeout.Truncate(5 * time.Second),
	// 	TLSHandshakeTimeout:   5 * time.Second,
	// 	MaxIdleConns:          int(config.Timeout.Abs().Truncate(5 * time.Second)),
	// 	ExpectContinueTimeout: 5 * time.Second,
	// }

	// client := &http.Client{
	// 	Transport: roundTripper,
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Printf("Error gettng response from requerst: %v", err)
	// }
	// defer resp.Body.Close()

	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading response body: %v", err)
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Errorf("failed to retrieve OpenAPI specs: %s", resp.Status)
	// }

	// fmt.Printf("Response Body: %v", string(body))
}
