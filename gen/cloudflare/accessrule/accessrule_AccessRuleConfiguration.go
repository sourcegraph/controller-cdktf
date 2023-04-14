package accessrule


type AccessRuleConfiguration struct {
	// The request property to target.
	//
	// Available values: `ip`, `ip6`, `ip_range`, `asn`, `country`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_rule#target AccessRule#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The value to target. Depends on target's type. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_rule#value AccessRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

