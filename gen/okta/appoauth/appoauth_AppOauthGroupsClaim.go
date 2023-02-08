package appoauth


type AppOauthGroupsClaim struct {
	// Name of the claim that will be used in the token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#name AppOauth#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Groups claim type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#type AppOauth#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the claim.
	//
	// Can be an Okta Expression Language statement that evaluates at the time the token is minted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#value AppOauth#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Groups claim filter. Can only be set if type is FILTER.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#filter_type AppOauth#filter_type}
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

