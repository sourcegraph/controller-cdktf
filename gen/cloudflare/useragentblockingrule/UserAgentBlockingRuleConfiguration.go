package useragentblockingrule


type UserAgentBlockingRuleConfiguration struct {
	// The configuration target for this rule. You must set the target to ua for User Agent Blocking rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#target UserAgentBlockingRule#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The exact user agent string to match. This value will be compared to the received User-Agent HTTP header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#value UserAgentBlockingRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

