package loadbalancer


type LoadBalancerRules struct {
	// Human readable name for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#name LoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The statement to evaluate to determine if this rule's effects should be applied.
	//
	// An empty condition is always true. See [load balancing rules](https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#condition LoadBalancer#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// A disabled rule will not be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#disabled LoadBalancer#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#fixed_response LoadBalancer#fixed_response}
	FixedResponse *LoadBalancerRulesFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#overrides LoadBalancer#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Priority used when determining the order of rule execution.
	//
	// Lower values are executed first. If not provided, the list order will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#priority LoadBalancer#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Terminates indicates that if this rule is true no further rules should be executed.
	//
	// Note: setting a [`fixed_response`](#fixed_response) forces this field to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#terminates LoadBalancer#terminates}
	Terminates interface{} `field:"optional" json:"terminates" yaml:"terminates"`
}

