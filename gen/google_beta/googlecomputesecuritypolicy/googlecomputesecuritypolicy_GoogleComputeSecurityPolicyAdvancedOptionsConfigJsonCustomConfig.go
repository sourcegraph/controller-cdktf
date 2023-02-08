package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// A list of custom Content-Type header values to apply the JSON parsing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#content_types GoogleComputeSecurityPolicy#content_types}
	ContentTypes *[]*string `field:"required" json:"contentTypes" yaml:"contentTypes"`
}

