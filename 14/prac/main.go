package main

import (
	"context"
	"fmt"
	"log"

	v1 "baiding.tech/prac/pkg/apis/baiding.teach/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatal(err)
	}
	config.APIPath = "/apis/"
	config.NegotiatedSerializer = v1.Codecs.WithoutConversion()
	config.GroupVersion = &v1.GroupVersion

	client, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalln(err)
	}

	foo := v1.JohnFoo{}
	err = client.Get().Namespace("default").Resource("johnfoos").Name("john-test").Do(context.TODO()).Into(&foo)
	if err != nil {
		log.Fatalln(err)
	}

	newObj := foo.DeepCopy()
	newObj.Spec.Name = "test2"

	fmt.Println(foo.Spec.Name)
	fmt.Println(foo.Spec.Replicas)

	fmt.Println(newObj.Spec.Name)
}
