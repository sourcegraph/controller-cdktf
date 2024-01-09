package directdatadog


type DirectDatadogHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#default_duration DirectDatadog#default_duration}
	DefaultDuration interface{} `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_datadog#max_duration DirectDatadog#max_duration}
	MaxDuration interface{} `field:"required" json:"maxDuration" yaml:"maxDuration"`
}

