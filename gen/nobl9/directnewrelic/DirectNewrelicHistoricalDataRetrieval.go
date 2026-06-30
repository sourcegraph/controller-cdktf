package directnewrelic

type DirectNewrelicHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_newrelic#default_duration DirectNewrelic#default_duration}
	DefaultDuration any `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_newrelic#max_duration DirectNewrelic#max_duration}
	MaxDuration any `field:"required" json:"maxDuration" yaml:"maxDuration"`
}
