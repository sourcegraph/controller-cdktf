package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersSourceCodeCloudStorageSource struct {
	// The Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#bucket GoogleCloudRunV2Service#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The Cloud Storage object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#object GoogleCloudRunV2Service#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// The Cloud Storage object generation.
	//
	// The is an int64 value. As with most Google APIs, its JSON representation will be a string instead of an integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#generation GoogleCloudRunV2Service#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

