package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialJwt struct {
	// Identifies which algorithm is used to generate the signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#jwt_header GoogleIntegrationsAuthConfig#jwt_header}
	JwtHeader *string `field:"optional" json:"jwtHeader" yaml:"jwtHeader"`
	// Contains a set of claims.
	//
	// The JWT specification defines seven Registered Claim Names which are the standard fields commonly included in tokens. Custom claims are usually also included, depending on the purpose of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#jwt_payload GoogleIntegrationsAuthConfig#jwt_payload}
	JwtPayload *string `field:"optional" json:"jwtPayload" yaml:"jwtPayload"`
	// User's pre-shared secret to sign the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#secret GoogleIntegrationsAuthConfig#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

