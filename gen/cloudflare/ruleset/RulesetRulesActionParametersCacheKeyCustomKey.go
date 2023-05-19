package ruleset


type RulesetRulesActionParametersCacheKeyCustomKey struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#cookie Ruleset#cookie}
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#header Ruleset#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#host Ruleset#host}
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#query_string Ruleset#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#user Ruleset#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

