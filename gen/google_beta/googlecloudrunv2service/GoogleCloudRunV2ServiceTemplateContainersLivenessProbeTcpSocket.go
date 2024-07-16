package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersLivenessProbeTcpSocket struct {
	// Port number to access on the container.
	//
	// Must be in the range 1 to 65535.
	// If not specified, defaults to the exposed port of the container, which
	// is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service#port GoogleCloudRunV2Service#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

