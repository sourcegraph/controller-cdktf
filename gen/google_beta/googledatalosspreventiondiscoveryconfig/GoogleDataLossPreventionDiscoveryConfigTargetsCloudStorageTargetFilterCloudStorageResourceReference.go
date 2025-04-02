package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference struct {
	// The bucket to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_data_loss_prevention_discovery_config#bucket_name GoogleDataLossPreventionDiscoveryConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// If within a project-level config, then this must match the config's project id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_data_loss_prevention_discovery_config#project_id GoogleDataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

