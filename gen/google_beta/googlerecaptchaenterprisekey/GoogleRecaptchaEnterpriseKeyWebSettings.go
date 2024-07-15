package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyWebSettings struct {
	// Required. Describes how this key is integrated with the website. Possible values: SCORE, CHECKBOX, INVISIBLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#integration_type GoogleRecaptchaEnterpriseKey#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// If set to true, it means allowed_domains will not be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allow_all_domains GoogleRecaptchaEnterpriseKey#allow_all_domains}
	AllowAllDomains interface{} `field:"optional" json:"allowAllDomains" yaml:"allowAllDomains"`
	// If set to true, the key can be used on AMP (Accelerated Mobile Pages) websites.
	//
	// This is supported only for the SCORE integration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allow_amp_traffic GoogleRecaptchaEnterpriseKey#allow_amp_traffic}
	AllowAmpTraffic interface{} `field:"optional" json:"allowAmpTraffic" yaml:"allowAmpTraffic"`
	// Domains or subdomains of websites allowed to use the key.
	//
	// All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allowed_domains GoogleRecaptchaEnterpriseKey#allowed_domains}
	AllowedDomains *[]*string `field:"optional" json:"allowedDomains" yaml:"allowedDomains"`
	// Settings for the frequency and difficulty at which this key triggers captcha challenges.
	//
	// This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED, USABILITY, BALANCE, SECURITY
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#challenge_security_preference GoogleRecaptchaEnterpriseKey#challenge_security_preference}
	ChallengeSecurityPreference *string `field:"optional" json:"challengeSecurityPreference" yaml:"challengeSecurityPreference"`
}

