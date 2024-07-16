package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRecaptchaOptionsConfig struct {
	// A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy.
	//
	// The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#redirect_site_key GoogleComputeSecurityPolicy#redirect_site_key}
	RedirectSiteKey *string `field:"required" json:"redirectSiteKey" yaml:"redirectSiteKey"`
}

