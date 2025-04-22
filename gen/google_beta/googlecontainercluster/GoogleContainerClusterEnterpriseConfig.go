package googlecontainercluster


type GoogleContainerClusterEnterpriseConfig struct {
	// Indicates the desired cluster tier. Available options include STANDARD and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#desired_tier GoogleContainerCluster#desired_tier}
	DesiredTier *string `field:"optional" json:"desiredTier" yaml:"desiredTier"`
}

