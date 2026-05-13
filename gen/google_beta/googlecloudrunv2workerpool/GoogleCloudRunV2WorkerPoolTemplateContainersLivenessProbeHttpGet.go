package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#http_headers GoogleCloudRunV2WorkerPool#http_headers}
	HttpHeaders *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Optional. Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#path GoogleCloudRunV2WorkerPool#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Optional.
	//
	// Port number to access on the container. Must be in the range 1 to 65535. If not specified, defaults to the exposed port of the container, which is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#port GoogleCloudRunV2WorkerPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

