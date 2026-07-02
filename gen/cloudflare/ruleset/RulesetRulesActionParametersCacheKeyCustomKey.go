package ruleset

type RulesetRulesActionParametersCacheKeyCustomKey struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#cookie Ruleset#cookie}
	Cookie any `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#header Ruleset#header}
	Header any `field:"optional" json:"header" yaml:"header"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#host Ruleset#host}
	Host any `field:"optional" json:"host" yaml:"host"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#query_string Ruleset#query_string}
	QueryString any `field:"optional" json:"queryString" yaml:"queryString"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ruleset#user Ruleset#user}
	User any `field:"optional" json:"user" yaml:"user"`
}
