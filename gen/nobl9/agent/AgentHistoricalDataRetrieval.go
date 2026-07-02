package agent

type AgentHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/agent#default_duration Agent#default_duration}
	DefaultDuration any `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/agent#max_duration Agent#max_duration}
	MaxDuration any `field:"required" json:"maxDuration" yaml:"maxDuration"`
}
