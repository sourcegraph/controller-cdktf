package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateVolumesGcs struct {
	// GCS Bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#bucket GoogleCloudRunV2Service#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A list of flags to pass to the gcsfuse command for configuring this volume.
	//
	// Flags should be passed without leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#mount_options GoogleCloudRunV2Service#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// If true, mount the GCS bucket as read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#read_only GoogleCloudRunV2Service#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

