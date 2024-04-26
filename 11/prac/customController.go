package prac

// import (
// 	"context"
// 	"log"
// 	"reflect"
// 	"time"

// 	"k8s.io/apimachinery/pkg/api/errors"

// 	"k8s.io/apimachinery/pkg/api/meta"
// 	"k8s.io/apimachinery/pkg/util/runtime"
// 	"k8s.io/apimachinery/pkg/util/wait"

// 	v14 "k8s.io/api/core/v1"
// 	v12 "k8s.io/api/networking/v1"
// 	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	informer "k8s.io/client-go/informers/core/v1"
// 	netInformer "k8s.io/client-go/informers/networking/v1"
// 	"k8s.io/client-go/kubernetes"
// 	coreLister "k8s.io/client-go/listers/core/v1"
// 	v1 "k8s.io/client-go/listers/networking/v1"
// 	"k8s.io/client-go/tools/cache"
// 	"k8s.io/client-go/util/workqueue"
// )

// const (
// 	ServiceType = "Service"
// 	maxRetry = 3
// 	numWorkers = 4
	
// )
// type controller struct {
// 	client        kubernetes.Interface
// 	ingressLister v1.IngressLister
// 	serviceLister coreLister.ServiceLister
// 	queue         workqueue.RateLimitingInterface
// }

// func newController(client kubernetes.Interface, ingressInformer netInformer.IngressInformer, serviceInformer informer.ServiceInformer) *controller {

// 	ctl := &controller{
// 		client:        client,
// 		ingressLister: ingressInformer.Lister(),
// 		serviceLister: serviceInformer.Lister(),
// 		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "john_test"),
// 	}
// 	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc:    ctl.addService,
// 		UpdateFunc: ctl.updateService,
// 	})
// 	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		DeleteFunc: ctl.delIngress,
// 	})
// 	return ctl
// }

// func (ctl *controller) enqueue(obj interface{}) {
// 	key, err := cache.MetaNamespaceKeyFunc(obj)
// 	if err != nil {
// 		runtime.HandleError(err)
// 	}
// 	ctl.queue.Add(key)
// }

// func (ctl *controller) addService(obj interface{}) {
// 	ctl.enqueue(obj)
// }

// func (ctl *controller) updateService(oldObj interface{}, newObj interface{}) {
// 	// todo handler error; what if the oldObj is a string
// 	metaOld, _ := meta.Accessor(oldObj)
// 	metaNew, _ := meta.Accessor(newObj)
// 	if reflect.DeepEqual(metaOld.GetAnnotations(),metaNew.GetAnnotations()) {
// 		return
// 	}
// 	ctl.enqueue(newObj)
// }

// func (ctl *controller) delIngress(obj interface{}) {
// 	ing := obj.(*v12.Ingress)
// 	ownerRef := v13.GetControllerOf(ing)

// 	if ownerRef == nil {
// 		return
// 	}
// 	if ownerRef.Kind != ServiceType{
// 		return
// 	}
// 	ctl.queue.Add(ing.Namespace+"/"+ing.Name)

// }

// func (ctl *controller) run(stopChan chan struct{}){
// 	for i:=0; i< numWorkers; i++ {
// 		go wait.Until(ctl.worker,time.Minute,stopChan)
// 	}
// 	<- stopChan
// }

// func (ctl *controller) worker() {
// 	for ctl.processNextItem() {

// 	}
// }

// func (ctl *controller) processNextItem() bool {
// 	item,down := ctl.queue.Get()
// 	if down  {
// 		return false
// 	}
// 	defer ctl.queue.Done(item)
// 	key := item.(string)
// 	if err := ctl.syncService(key); err != nil {
// 		ctl.handlErr(key,err)
// 	}
// 	return true
// }

// func (ctl *controller) handlErr(key string, err error){
// 	if ctl.queue.NumRequeues(key) < maxRetry {
// 		ctl.queue.AddRateLimited(key)
// 	}
// 	runtime.HandleError(err)
// 	ctl.queue.Forget(key)
// }

// func (ctl *controller) syncService(key string) error {
// 	nsKey, name, err := cache.SplitMetaNamespaceKey(key)
// 	if err != nil {
// 		log.Fatalln("can not split meta namespace key")
// 	}

// 	svc , err := ctl.serviceLister.Services(nsKey).Get(name)
// 	if err != nil {
// 		if errors.IsNotFound(err) {
// 			return nil
// 		}
// 		return err
// 	}
// 	_, ok := svc.GetAnnotations()["ingress/http"]
// 	ing, err := ctl.ingressLister.Ingresses(nsKey).Get(name)
// 	if err != nil && !errors.IsNotFound(err) {
// 		return err
// 	}
	
// 	if ok && errors.IsNotFound(err) {
// 		// add
// 		newIng := ctl.constructIngress(svc)
// 		_, err := ctl.client.NetworkingV1().Ingresses(nsKey).Create(context.TODO(),newIng,v13.CreateOptions{})
// 		if err != nil {
// 			return err
// 		}
// 	}else if !ok && ing != nil {
// 		// del
// 		err := ctl.client.NetworkingV1().Ingresses(nsKey).Delete(context.TODO(),name , v13.DeleteOptions{})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (ctl *controller) constructIngress(svc *v14.Service) *v12.Ingress {
// 	ing := v12.Ingress{}
// 	ing.ObjectMeta.OwnerReferences = []v13.OwnerReference{
// 		*v13.NewControllerRef(svc, v14.SchemeGroupVersion.WithKind(ServiceType)),
// 	}
// 	ing.Namespace = svc.Namespace
// 	ing.Name = svc.Name
// 	pathType := v12.PathTypePrefix
// 	ingcn := "nginx"
// 	ing.Spec = v12.IngressSpec{
// 		IngressClassName: &ingcn,
// 		Rules: []v12.IngressRule{
// 			{
// 				Host:"example.com",
// 				IngressRuleValue: v12.IngressRuleValue{
// 					HTTP: &v12.HTTPIngressRuleValue{
// 						Paths: []v12.HTTPIngressPath{
// 							{
// 								Path : "/",
// 								PathType:&pathType,
// 								Backend : v12.IngressBackend{
// 									Service: &v12.IngressServiceBackend{
// 										Name: svc.Name,
// 										Port: v12.ServiceBackendPort{
// 											Number: 80,
// 										},
// 									},
// 								},
// 							},
							
// 						},
// 					},
// 				},
// 			},

// 		},
// 	}
// 	return &ing
// }