package labels


type LabelsMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/labels#name Labels#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/labels#namespace Labels#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

