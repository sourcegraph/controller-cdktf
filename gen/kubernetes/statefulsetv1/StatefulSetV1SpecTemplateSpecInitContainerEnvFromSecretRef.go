package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerEnvFromSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/stateful_set_v1#name StatefulSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/stateful_set_v1#optional StatefulSetV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

