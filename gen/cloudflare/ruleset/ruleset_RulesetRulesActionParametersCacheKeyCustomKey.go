package ruleset


type RulesetRulesActionParametersCacheKeyCustomKey struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cookie Ruleset#cookie}
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#header Ruleset#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#host Ruleset#host}
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#query_string Ruleset#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#user Ruleset#user}
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

