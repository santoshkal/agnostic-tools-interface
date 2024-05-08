package validate

import (
	"flag"
	"fmt"

	"github.com/imroc/req/v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string
var reqClient = req.C()

func init() {
	// kubeconfig file parsing
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func CreateClient(kubeconfig string) (*kubernetes.Clientset, error) {
	// create the config object from kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	// replaceTransport(config, reqClient.GetTransport())

	// create clientset (set of muliple clients) for each Group (e.g. Core),
	// the Version (V1) of Group and Kind (e.g. Pods) so GVK.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating clientset: %v", err)
	}
	return clientset, nil
}

// replace client-go's Transport with *req.Transport
// func replaceTransport(config *rest.Config, t *req.Transport) {
// 	// Extract tls.Config from rest.Config
// 	tlsConfig, err := rest.TLSConfigFor(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// Set TLSClientConfig to req's Transport.
// 	t.TLSClientConfig = tlsConfig
// 	// Override with req's Transport.
// 	config.Transport = t
// 	// rest.Config.TLSClientConfig should be empty if
// 	// custom Transport been set.
// 	config.TLSClientConfig = rest.TLSClientConfig{
// 		Insecure: true,
// 	}
// }
