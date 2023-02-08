package datacloudflarewafrules


type DataCloudflareWafRulesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_rules#description DataCloudflareWafRules#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_rules#group_id DataCloudflareWafRules#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/d/waf_rules#mode DataCloudflareWafRules#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

