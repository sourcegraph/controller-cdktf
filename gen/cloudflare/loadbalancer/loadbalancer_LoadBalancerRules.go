package loadbalancer


type LoadBalancerRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#name LoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#condition LoadBalancer#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#disabled LoadBalancer#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#fixed_response LoadBalancer#fixed_response}
	FixedResponse *LoadBalancerRulesFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#overrides LoadBalancer#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#priority LoadBalancer#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#terminates LoadBalancer#terminates}.
	Terminates interface{} `field:"optional" json:"terminates" yaml:"terminates"`
}

