package agent


type AgentSplunkObservabilityConfig struct {
	// SplunkObservability Realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#realm Agent#realm}
	Realm *string `field:"required" json:"realm" yaml:"realm"`
}

