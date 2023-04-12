package pagerule


type PageRuleActionsCacheKeyFields struct {
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#cookie PageRule#cookie}
	Cookie *PageRuleActionsCacheKeyFieldsCookie `field:"required" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#header PageRule#header}
	Header *PageRuleActionsCacheKeyFieldsHeader `field:"required" json:"header" yaml:"header"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#host PageRule#host}
	Host *PageRuleActionsCacheKeyFieldsHost `field:"required" json:"host" yaml:"host"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#query_string PageRule#query_string}
	QueryString *PageRuleActionsCacheKeyFieldsQueryString `field:"required" json:"queryString" yaml:"queryString"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#user PageRule#user}
	User *PageRuleActionsCacheKeyFieldsUser `field:"required" json:"user" yaml:"user"`
}

