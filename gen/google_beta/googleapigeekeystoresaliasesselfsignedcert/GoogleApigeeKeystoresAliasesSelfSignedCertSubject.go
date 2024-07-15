package googleapigeekeystoresaliasesselfsignedcert


type GoogleApigeeKeystoresAliasesSelfSignedCertSubject struct {
	// Common name of the organization. Maximum length is 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#common_name GoogleApigeeKeystoresAliasesSelfSignedCert#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-letter country code. Example, IN for India, US for United States of America.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#country_code GoogleApigeeKeystoresAliasesSelfSignedCert#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Email address. Max 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#email GoogleApigeeKeystoresAliasesSelfSignedCert#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// City or town name. Maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#locality GoogleApigeeKeystoresAliasesSelfSignedCert#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Organization name. Maximum length is 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#org GoogleApigeeKeystoresAliasesSelfSignedCert#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Organization team name. Maximum length is 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#org_unit GoogleApigeeKeystoresAliasesSelfSignedCert#org_unit}
	OrgUnit *string `field:"optional" json:"orgUnit" yaml:"orgUnit"`
	// State or district name. Maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert#state GoogleApigeeKeystoresAliasesSelfSignedCert#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

