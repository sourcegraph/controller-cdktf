package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction struct {
	// The id of the relevant artifact in the recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#artifact_id GoogleOsConfigGuestPolicies#artifact_id}
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// The type of the archive to extract. Possible values: ["TAR", "TAR_GZIP", "TAR_BZIP", "TAR_LZMA", "TAR_XZ", "ZIP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#type GoogleOsConfigGuestPolicies#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#destination GoogleOsConfigGuestPolicies#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

