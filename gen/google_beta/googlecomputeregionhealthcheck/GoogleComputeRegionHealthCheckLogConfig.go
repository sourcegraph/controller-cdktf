package googlecomputeregionhealthcheck


type GoogleComputeRegionHealthCheckLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// This is false by default,
	// which means no health check logging will be done.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_health_check#enable GoogleComputeRegionHealthCheck#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

