package pagerule


type PageRuleActionsCacheKeyFields struct {
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#host PageRule#host}
	Host *PageRuleActionsCacheKeyFieldsHost `field:"required" json:"host" yaml:"host"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#query_string PageRule#query_string}
	QueryString *PageRuleActionsCacheKeyFieldsQueryString `field:"required" json:"queryString" yaml:"queryString"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#user PageRule#user}
	User *PageRuleActionsCacheKeyFieldsUser `field:"required" json:"user" yaml:"user"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#cookie PageRule#cookie}
	Cookie *PageRuleActionsCacheKeyFieldsCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#header PageRule#header}
	Header *PageRuleActionsCacheKeyFieldsHeader `field:"optional" json:"header" yaml:"header"`
}

