package pagerule


type PageRuleActionsCacheKeyFieldsUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#device_type PageRule#device_type}.
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#geo PageRule#geo}.
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/page_rule#lang PageRule#lang}.
	Lang interface{} `field:"optional" json:"lang" yaml:"lang"`
}

