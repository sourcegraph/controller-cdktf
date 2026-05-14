package googlecomputeorganizationsecuritypolicy


type GoogleComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// A list of content types to be parsed as JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#content_types GoogleComputeOrganizationSecurityPolicy#content_types}
	ContentTypes *[]*string `field:"required" json:"contentTypes" yaml:"contentTypes"`
}

