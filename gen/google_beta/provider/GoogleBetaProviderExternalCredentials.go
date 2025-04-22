package provider


type GoogleBetaProviderExternalCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs#audience GoogleBetaProvider#audience}.
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs#identity_token GoogleBetaProvider#identity_token}.
	IdentityToken *string `field:"required" json:"identityToken" yaml:"identityToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs#service_account_email GoogleBetaProvider#service_account_email}.
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

