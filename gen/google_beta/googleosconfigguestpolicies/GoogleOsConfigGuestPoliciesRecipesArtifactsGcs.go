package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesArtifactsGcs struct {
	// Bucket of the Google Cloud Storage object. Given an example URL: https://storage.googleapis.com/my-bucket/foo/bar#1234567 this value would be my-bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#bucket GoogleOsConfigGuestPolicies#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Must be provided if allowInsecure is false.
	//
	// Generation number of the Google Cloud Storage object.
	// https://storage.googleapis.com/my-bucket/foo/bar#1234567 this value would be 1234567.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#generation GoogleOsConfigGuestPolicies#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
	// Name of the Google Cloud Storage object. Given an example URL: https://storage.googleapis.com/my-bucket/foo/bar#1234567 this value would be foo/bar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#object GoogleOsConfigGuestPolicies#object}
	Object *string `field:"optional" json:"object" yaml:"object"`
}

