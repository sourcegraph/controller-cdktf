package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesArtifacts struct {
	// Id of the artifact, which the installation and update steps of this recipe can reference.
	//
	// Artifacts in a recipe cannot have the same id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#id GoogleOsConfigGuestPolicies#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Defaults to false.
	//
	// When false, recipes are subject to validations based on the artifact type:
	// Remote: A checksum must be specified, and only protocols with transport-layer security are permitted.
	// GCS: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#allow_insecure GoogleOsConfigGuestPolicies#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#gcs GoogleOsConfigGuestPolicies#gcs}
	Gcs *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#remote GoogleOsConfigGuestPolicies#remote}
	Remote *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote `field:"optional" json:"remote" yaml:"remote"`
}

