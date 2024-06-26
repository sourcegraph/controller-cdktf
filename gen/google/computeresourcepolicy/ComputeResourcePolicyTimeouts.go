package computeresourcepolicy


type ComputeResourcePolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/compute_resource_policy#create ComputeResourcePolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/compute_resource_policy#delete ComputeResourcePolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

