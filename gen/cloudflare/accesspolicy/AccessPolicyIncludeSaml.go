package accesspolicy


type AccessPolicyIncludeSaml struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#attribute_name AccessPolicy#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#attribute_value AccessPolicy#attribute_value}.
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}

