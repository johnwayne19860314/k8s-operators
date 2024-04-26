package prac

// import (
// 	"log"

// 	"k8s.io/client-go/informers"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	// config
// 	// client
// 	// informer
// 	// add event handler
// 	// informer start

// 	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
// 	if err != nil {
// 		inClusterConfig, err := rest.InClusterConfig()
// 		if err != nil {
// 			log.Fatalln("can not load config")
// 		}
// 		config = inClusterConfig
// 	}

// 	clientSet, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		log.Fatalln(" can not get clientset")
// 	}

// 	factory := informers.NewSharedInformerFactory(clientSet, 0)
// 	serviceInformer := factory.Core().V1().Services()
// 	ingressInformer := factory.Networking().V1().Ingresses()

// 	ctl := newController(clientSet, ingressInformer, serviceInformer)

// 	stopChan := make(chan struct{})
// 	factory.Start(stopChan)
// 	factory.WaitForCacheSync(stopChan)

// 	ctl.run(stopChan)

// }
