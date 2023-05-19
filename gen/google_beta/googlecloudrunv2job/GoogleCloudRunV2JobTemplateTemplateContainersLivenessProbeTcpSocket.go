package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersLivenessProbeTcpSocket struct {
	// Port number to access on the container.
	//
	// Must be in the range 1 to 65535. If not specified, defaults to 8080.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloud_run_v2_job#port GoogleCloudRunV2Job#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

