package googlecomputeinstance


type GoogleComputeInstanceServiceAccount struct {
	// A list of service scopes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#scopes GoogleComputeInstance#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#email GoogleComputeInstance#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

