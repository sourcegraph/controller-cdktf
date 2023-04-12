package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdaptiveProtectionConfig struct {
	// layer_7_ddos_defense_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#layer_7_ddos_defense_config GoogleComputeSecurityPolicy#layer_7_ddos_defense_config}
	Layer7DdosDefenseConfig *GoogleComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig `field:"optional" json:"layer7DdosDefenseConfig" yaml:"layer7DdosDefenseConfig"`
}

