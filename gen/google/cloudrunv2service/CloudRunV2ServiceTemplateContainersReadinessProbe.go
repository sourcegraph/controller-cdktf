package cloudrunv2service


type CloudRunV2ServiceTemplateContainersReadinessProbe struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#failure_threshold CloudRunV2Service#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#grpc CloudRunV2Service#grpc}
	Grpc *CloudRunV2ServiceTemplateContainersReadinessProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#http_get CloudRunV2Service#http_get}
	HttpGet *CloudRunV2ServiceTemplateContainersReadinessProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// How often (in seconds) to perform the probe. Default to 10 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#period_seconds CloudRunV2Service#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#success_threshold CloudRunV2Service#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out. Defaults to 1 second. Must be smaller than period_seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_service#timeout_seconds CloudRunV2Service#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

