package oauthapp


type OauthAppGroupsClaim struct {
	// Name of the claim that will be used in the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/oauth_app#name OauthApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Groups claim type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/oauth_app#type OauthApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the claim.
	//
	// Can be an Okta Expression Language statement that evaluates at the time the token is minted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/oauth_app#value OauthApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Groups claim filter. Can only be set if type is FILTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/oauth_app#filter_type OauthApp#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

