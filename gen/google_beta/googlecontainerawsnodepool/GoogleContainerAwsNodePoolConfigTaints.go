package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigTaints struct {
	// The taint effect. Possible values: EFFECT_UNSPECIFIED, NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#effect GoogleContainerAwsNodePool#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Key for the taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#key GoogleContainerAwsNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for the taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#value GoogleContainerAwsNodePool#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

