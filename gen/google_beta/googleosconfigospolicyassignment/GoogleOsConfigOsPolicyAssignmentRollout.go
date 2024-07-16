package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#disruption_budget GoogleOsConfigOsPolicyAssignment#disruption_budget}
	DisruptionBudget *GoogleOsConfigOsPolicyAssignmentRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// This determines the minimum duration of time to wait after the configuration changes are applied through the current rollout.
	//
	// A VM continues to count towards the 'disruption_budget' at least until this duration of time has passed after configuration changes are applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#min_wait_duration GoogleOsConfigOsPolicyAssignment#min_wait_duration}
	MinWaitDuration *string `field:"required" json:"minWaitDuration" yaml:"minWaitDuration"`
}

