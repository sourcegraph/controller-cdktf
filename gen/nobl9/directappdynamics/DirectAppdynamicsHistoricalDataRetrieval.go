package directappdynamics


type DirectAppdynamicsHistoricalDataRetrieval struct {
	// default_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_appdynamics#default_duration DirectAppdynamics#default_duration}
	DefaultDuration interface{} `field:"required" json:"defaultDuration" yaml:"defaultDuration"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/direct_appdynamics#max_duration DirectAppdynamics#max_duration}
	MaxDuration interface{} `field:"required" json:"maxDuration" yaml:"maxDuration"`
}

