package accessidentityprovider


type AccessIdentityProviderConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#api_token AccessIdentityProvider#api_token}.
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#apps_domain AccessIdentityProvider#apps_domain}.
	AppsDomain *string `field:"optional" json:"appsDomain" yaml:"appsDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#attributes AccessIdentityProvider#attributes}.
	Attributes *[]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#auth_url AccessIdentityProvider#auth_url}.
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#centrify_account AccessIdentityProvider#centrify_account}.
	CentrifyAccount *string `field:"optional" json:"centrifyAccount" yaml:"centrifyAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#centrify_app_id AccessIdentityProvider#centrify_app_id}.
	CentrifyAppId *string `field:"optional" json:"centrifyAppId" yaml:"centrifyAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#certs_url AccessIdentityProvider#certs_url}.
	CertsUrl *string `field:"optional" json:"certsUrl" yaml:"certsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#claims AccessIdentityProvider#claims}.
	Claims *[]*string `field:"optional" json:"claims" yaml:"claims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#client_id AccessIdentityProvider#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#client_secret AccessIdentityProvider#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#directory_id AccessIdentityProvider#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#email_attribute_name AccessIdentityProvider#email_attribute_name}.
	EmailAttributeName *string `field:"optional" json:"emailAttributeName" yaml:"emailAttributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#idp_public_cert AccessIdentityProvider#idp_public_cert}.
	IdpPublicCert *string `field:"optional" json:"idpPublicCert" yaml:"idpPublicCert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#issuer_url AccessIdentityProvider#issuer_url}.
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#okta_account AccessIdentityProvider#okta_account}.
	OktaAccount *string `field:"optional" json:"oktaAccount" yaml:"oktaAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#onelogin_account AccessIdentityProvider#onelogin_account}.
	OneloginAccount *string `field:"optional" json:"oneloginAccount" yaml:"oneloginAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#pkce_enabled AccessIdentityProvider#pkce_enabled}.
	PkceEnabled interface{} `field:"optional" json:"pkceEnabled" yaml:"pkceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#redirect_url AccessIdentityProvider#redirect_url}.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#scopes AccessIdentityProvider#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#sign_request AccessIdentityProvider#sign_request}.
	SignRequest interface{} `field:"optional" json:"signRequest" yaml:"signRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#sso_target_url AccessIdentityProvider#sso_target_url}.
	SsoTargetUrl *string `field:"optional" json:"ssoTargetUrl" yaml:"ssoTargetUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#support_groups AccessIdentityProvider#support_groups}.
	SupportGroups interface{} `field:"optional" json:"supportGroups" yaml:"supportGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#token_url AccessIdentityProvider#token_url}.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

