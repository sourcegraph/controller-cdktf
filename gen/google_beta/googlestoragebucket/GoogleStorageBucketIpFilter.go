package googlestoragebucket


type GoogleStorageBucketIpFilter struct {
	// The mode of the IP filter. Valid values are 'Enabled' and 'Disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#mode GoogleStorageBucket#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Whether to allow all service agents to access the bucket regardless of the IP filter configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#allow_all_service_agent_access GoogleStorageBucket#allow_all_service_agent_access}
	AllowAllServiceAgentAccess interface{} `field:"optional" json:"allowAllServiceAgentAccess" yaml:"allowAllServiceAgentAccess"`
	// Whether to allow cross-org VPCs in the bucket's IP filter configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#allow_cross_org_vpcs GoogleStorageBucket#allow_cross_org_vpcs}
	AllowCrossOrgVpcs interface{} `field:"optional" json:"allowCrossOrgVpcs" yaml:"allowCrossOrgVpcs"`
	// public_network_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#public_network_source GoogleStorageBucket#public_network_source}
	PublicNetworkSource *GoogleStorageBucketIpFilterPublicNetworkSource `field:"optional" json:"publicNetworkSource" yaml:"publicNetworkSource"`
	// vpc_network_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#vpc_network_sources GoogleStorageBucket#vpc_network_sources}
	VpcNetworkSources interface{} `field:"optional" json:"vpcNetworkSources" yaml:"vpcNetworkSources"`
}

