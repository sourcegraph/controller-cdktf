package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdaptiveProtectionConfig struct {
	// auto_deploy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#auto_deploy_config GoogleComputeSecurityPolicy#auto_deploy_config}
	AutoDeployConfig *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig `field:"optional" json:"autoDeployConfig" yaml:"autoDeployConfig"`
	// layer_7_ddos_defense_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#layer_7_ddos_defense_config GoogleComputeSecurityPolicy#layer_7_ddos_defense_config}
	Layer7DdosDefenseConfig *GoogleComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig `field:"optional" json:"layer7DdosDefenseConfig" yaml:"layer7DdosDefenseConfig"`
}

