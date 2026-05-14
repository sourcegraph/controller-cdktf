package computeorganizationsecuritypolicy


type ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig struct {
	// A list of content types to be parsed as JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_organization_security_policy#content_types ComputeOrganizationSecurityPolicy#content_types}
	ContentTypes *[]*string `field:"required" json:"contentTypes" yaml:"contentTypes"`
}

