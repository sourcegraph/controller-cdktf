package pagerule


type PageRuleActionsCacheKeyFieldsHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#check_presence PageRule#check_presence}.
	CheckPresence *[]*string `field:"optional" json:"checkPresence" yaml:"checkPresence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#exclude PageRule#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/page_rule#include PageRule#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

