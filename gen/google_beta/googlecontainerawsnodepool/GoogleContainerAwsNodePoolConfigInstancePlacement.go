package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigInstancePlacement struct {
	// The tenancy for the instance. Possible values: TENANCY_UNSPECIFIED, DEFAULT, DEDICATED, HOST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#tenancy GoogleContainerAwsNodePool#tenancy}
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

