package googlecomputeinstance


type GoogleComputeInstanceServiceAccount struct {
	// A list of service scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#scopes GoogleComputeInstance#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#email GoogleComputeInstance#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

