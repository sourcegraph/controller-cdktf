package labels


type LabelsMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/labels#name Labels#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/labels#namespace Labels#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

