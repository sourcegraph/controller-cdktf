package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc struct {
	// Optional.
	//
	// Port number of the gRPC service. Number must be in the range 1 to 65535. If not specified, defaults to the exposed port of the container, which is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#port GoogleCloudRunV2WorkerPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Optional.
	//
	// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md ). If this is not specified, the default behavior is defined by gRPC
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#service GoogleCloudRunV2WorkerPool#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

