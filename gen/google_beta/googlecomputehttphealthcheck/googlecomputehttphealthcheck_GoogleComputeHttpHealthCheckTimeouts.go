package googlecomputehttphealthcheck


type GoogleComputeHttpHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_http_health_check#create GoogleComputeHttpHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_http_health_check#delete GoogleComputeHttpHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_http_health_check#update GoogleComputeHttpHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

