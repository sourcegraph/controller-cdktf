package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumesSecretItems struct {
	// The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#key GoogleCloudRunService#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The relative path of the file to map the key to.
	//
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#path GoogleCloudRunService#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Mode bits to use on this file, must be a value between 0000 and 0777.
	//
	// If
	// not specified, the volume defaultMode will be used. This might be in
	// conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_service#mode GoogleCloudRunService#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
}

