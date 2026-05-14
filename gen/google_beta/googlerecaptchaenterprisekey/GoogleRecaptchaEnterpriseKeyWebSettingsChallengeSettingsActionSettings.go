package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettings struct {
	// The action name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_recaptcha_enterprise_key#action GoogleRecaptchaEnterpriseKey#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// A challenge is triggered if the end-user score is below that threshold.
	//
	// Value must be between 0 and 1 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_recaptcha_enterprise_key#score_threshold GoogleRecaptchaEnterpriseKey#score_threshold}
	ScoreThreshold *float64 `field:"required" json:"scoreThreshold" yaml:"scoreThreshold"`
}

