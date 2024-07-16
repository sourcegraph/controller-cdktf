package googlecomputerouternat


type GoogleComputeRouterNatRules struct {
	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	//
	// If it evaluates to true, the corresponding action is enforced.
	//
	// The following examples are valid match expressions for public NAT:
	//
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	//
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	//
	// The following example is a valid match expression for private NAT:
	//
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_nat#match GoogleComputeRouterNat#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// An integer uniquely identifying a rule in the list.
	//
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_nat#rule_number GoogleComputeRouterNat#rule_number}
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_nat#action GoogleComputeRouterNat#action}
	Action *GoogleComputeRouterNatRulesAction `field:"optional" json:"action" yaml:"action"`
	// An optional description of this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_nat#description GoogleComputeRouterNat#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

