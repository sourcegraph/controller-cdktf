package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersReadinessProbeHttpGet struct {
	// Path to access on the HTTP server. If set, it should not be empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#path GoogleCloudRunV2Service#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port number to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// If not specified, defaults to the same value as container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_service#port GoogleCloudRunV2Service#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

