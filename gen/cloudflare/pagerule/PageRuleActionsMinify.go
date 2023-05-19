package pagerule


type PageRuleActionsMinify struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#css PageRule#css}.
	Css *string `field:"required" json:"css" yaml:"css"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#html PageRule#html}.
	Html *string `field:"required" json:"html" yaml:"html"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#js PageRule#js}.
	Js *string `field:"required" json:"js" yaml:"js"`
}

