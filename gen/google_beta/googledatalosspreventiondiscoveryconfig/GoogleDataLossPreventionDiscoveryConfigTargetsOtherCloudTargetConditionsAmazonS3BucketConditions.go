package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsAmazonS3BucketConditions struct {
	// Bucket types that should be profiled. Optional. Defaults to TYPE_ALL_SUPPORTED if unspecified. Possible values: ["TYPE_ALL_SUPPORTED", "TYPE_GENERAL_PURPOSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#bucket_types GoogleDataLossPreventionDiscoveryConfig#bucket_types}
	BucketTypes *[]*string `field:"optional" json:"bucketTypes" yaml:"bucketTypes"`
	// Object classes that should be profiled. Optional. Defaults to ALL_SUPPORTED_CLASSES if unspecified. Possible values: ["ALL_SUPPORTED_CLASSES", "STANDARD", "STANDARD_INFREQUENT_ACCESS", "GLACIER_INSTANT_RETRIEVAL", "INTELLIGENT_TIERING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#object_storage_classes GoogleDataLossPreventionDiscoveryConfig#object_storage_classes}
	ObjectStorageClasses *[]*string `field:"optional" json:"objectStorageClasses" yaml:"objectStorageClasses"`
}

