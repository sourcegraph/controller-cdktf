package env


type EnvMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/env#name Env#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/env#namespace Env#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

