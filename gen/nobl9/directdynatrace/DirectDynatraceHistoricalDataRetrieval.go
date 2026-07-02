package directdynatrace

type DirectDynatraceHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_dynatrace#default_duration DirectDynatrace#default_duration}
	DefaultDuration any `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_dynatrace#max_duration DirectDynatrace#max_duration}
	MaxDuration any `field:"required" json:"maxDuration" yaml:"maxDuration"`
}
