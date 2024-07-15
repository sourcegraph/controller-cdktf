package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumesCsi struct {
	// Unique name representing the type of file system to be created.
	//
	// Cloud Run supports the following values:
	//   * gcsfuse.run.googleapis.com: Mount a Google Cloud Storage bucket using GCSFuse. This driver requires the
	//     run.googleapis.com/execution-environment annotation to be set to "gen2" and
	//     run.googleapis.com/launch-stage set to "BETA" or "ALPHA".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#driver GoogleCloudRunService#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// If true, all mounts created from this volume will be read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#read_only GoogleCloudRunService#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Driver-specific attributes.
	//
	// The following options are supported for available drivers:
	//   * gcsfuse.run.googleapis.com
	//     * bucketName: The name of the Cloud Storage Bucket that backs this volume. The Cloud Run Service identity must have access to this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#volume_attributes GoogleCloudRunService#volume_attributes}
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

