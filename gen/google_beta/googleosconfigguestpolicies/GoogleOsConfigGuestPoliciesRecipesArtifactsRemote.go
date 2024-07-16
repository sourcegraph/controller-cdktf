package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesArtifactsRemote struct {
	// Must be provided if allowInsecure is false.
	//
	// SHA256 checksum in hex format, to compare to the checksum of the artifact.
	// If the checksum is not empty and it doesn't match the artifact then the recipe installation fails before running any
	// of the steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#check_sum GoogleOsConfigGuestPolicies#check_sum}
	CheckSum *string `field:"optional" json:"checkSum" yaml:"checkSum"`
	// URI from which to fetch the object. It should contain both the protocol and path following the format {protocol}://{location}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#uri GoogleOsConfigGuestPolicies#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

