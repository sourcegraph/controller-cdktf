package recaptchaenterprisekey


type RecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettings struct {
	// The action name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/recaptcha_enterprise_key#action RecaptchaEnterpriseKey#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// A challenge is triggered if the end-user score is below that threshold.
	//
	// Value must be between 0 and 1 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/recaptcha_enterprise_key#score_threshold RecaptchaEnterpriseKey#score_threshold}
	ScoreThreshold *float64 `field:"required" json:"scoreThreshold" yaml:"scoreThreshold"`
}

