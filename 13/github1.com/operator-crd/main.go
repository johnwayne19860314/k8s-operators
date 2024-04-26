package main

import (
	"context"
	"log"

	clientset "github1.com/operator-crd/pkg/generated/clientset/versioned"
	"github1.com/operator-crd/pkg/generated/informers/externalversions"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln(err)
	}
	clientSet, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	list, err := clientSet.CrdV1().Foos("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, foo := range list.Items {
		println(foo.Name)
	}

	factory := externalversions.NewSharedInformerFactory(clientSet,0)
	factory.Crd().V1().Foos().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {},
	})

}


// 	factory := externalversions.NewSharedInformerFactory(clientset, 0)
// 	factory.Crd().V1().Foos().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc: func(obj interface{}) {
// 			//todo
// 		},
// 	})
// 	//TODO
// }
