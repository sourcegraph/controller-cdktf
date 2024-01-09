package slo


type SloIndicator struct {
	// Name of the metric source (agent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#name Slo#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Kind of the metric source. One of {Agent, Direct}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#kind Slo#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Name of the metric source project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#project Slo#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

