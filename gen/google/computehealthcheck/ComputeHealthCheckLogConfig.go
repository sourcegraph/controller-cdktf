package computehealthcheck

type ComputeHealthCheckLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// This is false by default,
	// which means no health check logging will be done.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_health_check#enable ComputeHealthCheck#enable}
	Enable any `field:"optional" json:"enable" yaml:"enable"`
}
