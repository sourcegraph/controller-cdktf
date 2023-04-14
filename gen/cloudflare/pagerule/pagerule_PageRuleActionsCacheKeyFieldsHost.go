package pagerule


type PageRuleActionsCacheKeyFieldsHost struct {
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#resolved PageRule#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}

