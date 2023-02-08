package pagerule


type PageRuleActionsForwardingUrl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#status_code PageRule#status_code}.
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#url PageRule#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

