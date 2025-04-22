package googlecontainercluster


type GoogleContainerClusterMasterAuthorizedNetworksConfig struct {
	// cidr_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#cidr_blocks GoogleContainerCluster#cidr_blocks}
	CidrBlocks interface{} `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Whether Kubernetes master is accessible via Google Compute Engine Public IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#gcp_public_cidrs_access_enabled GoogleContainerCluster#gcp_public_cidrs_access_enabled}
	GcpPublicCidrsAccessEnabled interface{} `field:"optional" json:"gcpPublicCidrsAccessEnabled" yaml:"gcpPublicCidrsAccessEnabled"`
	// Whether authorized networks is enforced on the private endpoint or not. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#private_endpoint_enforcement_enabled GoogleContainerCluster#private_endpoint_enforcement_enabled}
	PrivateEndpointEnforcementEnabled interface{} `field:"optional" json:"privateEndpointEnforcementEnabled" yaml:"privateEndpointEnforcementEnabled"`
}

