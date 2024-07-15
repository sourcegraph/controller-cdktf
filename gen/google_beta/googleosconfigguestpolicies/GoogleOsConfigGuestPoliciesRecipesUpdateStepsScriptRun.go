package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun struct {
	// The shell script to be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#script GoogleOsConfigGuestPolicies#script}
	Script *string `field:"required" json:"script" yaml:"script"`
	// Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#allowed_exit_codes GoogleOsConfigGuestPolicies#allowed_exit_codes}
	AllowedExitCodes *[]*float64 `field:"optional" json:"allowedExitCodes" yaml:"allowedExitCodes"`
	// The script interpreter to use to run the script.
	//
	// If no interpreter is specified the script is executed directly,
	// which likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#interpreter GoogleOsConfigGuestPolicies#interpreter}
	Interpreter *string `field:"optional" json:"interpreter" yaml:"interpreter"`
}

