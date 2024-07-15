package googlecontainercluster


type GoogleContainerClusterMasterAuthorizedNetworksConfigCidrBlocks struct {
	// External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#cidr_block GoogleContainerCluster#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Field for users to identify CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#display_name GoogleContainerCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

