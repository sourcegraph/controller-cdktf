package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyRulesNetworkMatchUserDefinedFields struct {
	// Name of the user-defined field, as given in the definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_region_security_policy#name GoogleComputeRegionSecurityPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Matching values of the field.
	//
	// Each element can be a 32-bit unsigned decimal or hexadecimal (starting with "0x") number (e.g. "64") or range (e.g. "0x400-0x7ff").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_region_security_policy#values GoogleComputeRegionSecurityPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

