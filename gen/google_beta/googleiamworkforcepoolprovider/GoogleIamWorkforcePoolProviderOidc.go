package googleiamworkforcepoolprovider


type GoogleIamWorkforcePoolProviderOidc struct {
	// The client ID. Must match the audience claim of the JWT issued by the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_iam_workforce_pool_provider#client_id GoogleIamWorkforcePoolProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_iam_workforce_pool_provider#issuer_uri GoogleIamWorkforcePoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_iam_workforce_pool_provider#client_secret GoogleIamWorkforcePoolProvider#client_secret}
	ClientSecret *GoogleIamWorkforcePoolProviderOidcClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// web_sso_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_iam_workforce_pool_provider#web_sso_config GoogleIamWorkforcePoolProvider#web_sso_config}
	WebSsoConfig *GoogleIamWorkforcePoolProviderOidcWebSsoConfig `field:"optional" json:"webSsoConfig" yaml:"webSsoConfig"`
}

