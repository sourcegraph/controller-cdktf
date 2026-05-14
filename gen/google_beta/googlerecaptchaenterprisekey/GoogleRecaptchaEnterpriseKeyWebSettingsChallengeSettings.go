package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings struct {
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_recaptcha_enterprise_key#default_settings GoogleRecaptchaEnterpriseKey#default_settings}
	DefaultSettings *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings `field:"required" json:"defaultSettings" yaml:"defaultSettings"`
	// action_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_recaptcha_enterprise_key#action_settings GoogleRecaptchaEnterpriseKey#action_settings}
	ActionSettings interface{} `field:"optional" json:"actionSettings" yaml:"actionSettings"`
}

