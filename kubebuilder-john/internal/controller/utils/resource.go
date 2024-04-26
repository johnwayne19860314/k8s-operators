package utils

import (
	"bytes"
	"html/template"

	"github.com/john/api/v1beta1"
	"k8s.io/apimachinery/pkg/util/yaml"
	appv1 "k8s.io/api/apps/v1"
	netv1 "k8s.io/api/networking/v1"
	corev1 "k8s.io/api/core/v1"
)

func parseTemplate(templateName string, app *v1beta1.App) []byte {
	temp, err := template.ParseFiles("../internal/controller/template/" + templateName + ".yml")
	if err != nil {
		panic(err)
	}

	b := new(bytes.Buffer)
	err = temp.Execute(b, app)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func NewDeployment(app *v1beta1.App) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", app), d)
	if err != nil {
		panic(err)
	}
	return d
}

func NewService(app *v1beta1.App) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", app), s)
	if err != nil {
		panic(err)
	}
	return s
}


func NewIngress(app *v1beta1.App) *netv1.Ingress {
	i := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", app), i)
	if err != nil {
		panic(err)
	}
	return i
}
