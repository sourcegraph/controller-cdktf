package googlefilestoreinstance


type GoogleFilestoreInstancePerformanceConfigFixedIops struct {
	// The number of IOPS to provision for the instance. max_iops must be in multiple of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_filestore_instance#max_iops GoogleFilestoreInstance#max_iops}
	MaxIops *float64 `field:"optional" json:"maxIops" yaml:"maxIops"`
}

