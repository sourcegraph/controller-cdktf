package containercluster


type ContainerClusterMasterAuthorizedNetworksConfig struct {
	// cidr_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/container_cluster#cidr_blocks ContainerCluster#cidr_blocks}
	CidrBlocks interface{} `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Whether Kubernetes master is accessible via Google Compute Engine Public IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/container_cluster#gcp_public_cidrs_access_enabled ContainerCluster#gcp_public_cidrs_access_enabled}
	GcpPublicCidrsAccessEnabled interface{} `field:"optional" json:"gcpPublicCidrsAccessEnabled" yaml:"gcpPublicCidrsAccessEnabled"`
}

