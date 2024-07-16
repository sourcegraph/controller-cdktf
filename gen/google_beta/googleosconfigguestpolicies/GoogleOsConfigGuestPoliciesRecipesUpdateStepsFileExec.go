package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec struct {
	// A list of possible return values that the program can return to indicate a success. Defaults to [0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#allowed_exit_codes GoogleOsConfigGuestPolicies#allowed_exit_codes}
	AllowedExitCodes *[]*float64 `field:"optional" json:"allowedExitCodes" yaml:"allowedExitCodes"`
	// Arguments to be passed to the provided executable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#args GoogleOsConfigGuestPolicies#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The id of the relevant artifact in the recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#artifact_id GoogleOsConfigGuestPolicies#artifact_id}
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// The absolute path of the file on the local filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#local_path GoogleOsConfigGuestPolicies#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

