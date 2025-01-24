package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeRegexesPatternsCloudStorageRegex struct {
	// Regex to test the bucket name against.
	//
	// If empty, all buckets match. Example: "marketing2021" or "(marketing)\d{4}" will both match the bucket gs://marketing2021
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_data_loss_prevention_discovery_config#bucket_name_regex GoogleDataLossPreventionDiscoveryConfig#bucket_name_regex}
	BucketNameRegex *string `field:"optional" json:"bucketNameRegex" yaml:"bucketNameRegex"`
	// For organizations, if unset, will match all projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_data_loss_prevention_discovery_config#project_id_regex GoogleDataLossPreventionDiscoveryConfig#project_id_regex}
	ProjectIdRegex *string `field:"optional" json:"projectIdRegex" yaml:"projectIdRegex"`
}

