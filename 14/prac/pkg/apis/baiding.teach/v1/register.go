// +groupName=baiding.teach
package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	Schema = runtime.NewScheme()
	GroupVersion = schema.GroupVersion{Group: "baiding.teach", Version: "v1"}
	Codecs = serializer.NewCodecFactory(Schema) 
)