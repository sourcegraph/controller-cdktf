package googlecomputebackendservice


type GoogleComputeBackendServiceLocalityLbPolicies struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#custom_policy GoogleComputeBackendService#custom_policy}
	CustomPolicy *GoogleComputeBackendServiceLocalityLbPoliciesCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_backend_service#policy GoogleComputeBackendService#policy}
	Policy *GoogleComputeBackendServiceLocalityLbPoliciesPolicy `field:"optional" json:"policy" yaml:"policy"`
}

