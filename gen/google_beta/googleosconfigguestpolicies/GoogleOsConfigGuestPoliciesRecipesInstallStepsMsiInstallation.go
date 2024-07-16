package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation struct {
	// The id of the relevant artifact in the recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#artifact_id GoogleOsConfigGuestPolicies#artifact_id}
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#allowed_exit_codes GoogleOsConfigGuestPolicies#allowed_exit_codes}
	AllowedExitCodes *[]*float64 `field:"optional" json:"allowedExitCodes" yaml:"allowedExitCodes"`
	// The flags to use when installing the MSI. Defaults to the install flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#flags GoogleOsConfigGuestPolicies#flags}
	Flags *[]*string `field:"optional" json:"flags" yaml:"flags"`
}

