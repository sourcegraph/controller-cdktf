package pagerule

type PageRuleActionsCacheKeyFieldsHost struct {
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#resolved PageRule#resolved}
	Resolved any `field:"optional" json:"resolved" yaml:"resolved"`
}
