package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#fixed GoogleOsConfigOsPolicyAssignment#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#percent GoogleOsConfigOsPolicyAssignment#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

