package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersStartupProbeTcpSocket struct {
	// Optional.
	//
	// Port number to access on the container. Must be in the range 1 to 65535. If not specified, defaults to the exposed port of the container, which is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool#port CloudRunV2WorkerPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

