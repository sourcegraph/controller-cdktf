package accessrule


type AccessRuleConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_rule#target AccessRule#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_rule#value AccessRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

