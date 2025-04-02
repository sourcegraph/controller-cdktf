package directgcm


type DirectGcmHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_gcm#default_duration DirectGcm#default_duration}
	DefaultDuration interface{} `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_gcm#max_duration DirectGcm#max_duration}
	MaxDuration interface{} `field:"required" json:"maxDuration" yaml:"maxDuration"`
}

