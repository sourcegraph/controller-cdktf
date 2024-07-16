package googlecomputeforwardingrule


type GoogleComputeForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_forwarding_rule#namespace GoogleComputeForwardingRule#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Service Directory service to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_forwarding_rule#service GoogleComputeForwardingRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

