package computeresourcepolicy


type ComputeResourcePolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#create ComputeResourcePolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#delete ComputeResourcePolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

