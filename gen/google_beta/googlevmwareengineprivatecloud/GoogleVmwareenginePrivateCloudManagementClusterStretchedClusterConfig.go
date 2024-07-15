package googlevmwareengineprivatecloud


type GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig struct {
	// Zone that will remain operational when connection between the two zones is lost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#preferred_location GoogleVmwareenginePrivateCloud#preferred_location}
	PreferredLocation *string `field:"optional" json:"preferredLocation" yaml:"preferredLocation"`
	// Additional zone for a higher level of availability and load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#secondary_location GoogleVmwareenginePrivateCloud#secondary_location}
	SecondaryLocation *string `field:"optional" json:"secondaryLocation" yaml:"secondaryLocation"`
}

