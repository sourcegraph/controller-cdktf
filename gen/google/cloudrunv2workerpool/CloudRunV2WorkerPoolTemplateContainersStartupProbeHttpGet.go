package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#http_headers CloudRunV2WorkerPool#http_headers}
	HttpHeaders *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Optional. Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#path CloudRunV2WorkerPool#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Optional.
	//
	// Port number to access on the container. Must be in the range 1 to 65535. If not specified, defaults to the exposed port of the container, which is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#port CloudRunV2WorkerPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

