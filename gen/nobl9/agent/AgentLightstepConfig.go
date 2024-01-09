package agent


type AgentLightstepConfig struct {
	// Organization name registered in Lightstep.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#organization Agent#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Name of the Lightstep project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#project Agent#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

