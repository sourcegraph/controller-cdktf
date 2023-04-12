package googlecomputeforwardingrule


type GoogleComputeForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_forwarding_rule#namespace GoogleComputeForwardingRule#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Service Directory service to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_forwarding_rule#service GoogleComputeForwardingRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

