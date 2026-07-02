package recaptchaenterprisekey

type RecaptchaEnterpriseKeyWebSettingsChallengeSettings struct {
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/recaptcha_enterprise_key#default_settings RecaptchaEnterpriseKey#default_settings}
	DefaultSettings *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings `field:"required" json:"defaultSettings" yaml:"defaultSettings"`
	// action_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/recaptcha_enterprise_key#action_settings RecaptchaEnterpriseKey#action_settings}
	ActionSettings any `field:"optional" json:"actionSettings" yaml:"actionSettings"`
}
