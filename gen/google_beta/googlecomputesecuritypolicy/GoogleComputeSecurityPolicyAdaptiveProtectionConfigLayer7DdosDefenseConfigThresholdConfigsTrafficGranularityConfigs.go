package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfigsTrafficGranularityConfigs struct {
	// Type of this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_security_policy#type GoogleComputeSecurityPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If enabled, traffic matching each unique value for the specified type constitutes a separate traffic unit.
	//
	// It can only be set to true if value is empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_security_policy#enable_each_unique_value GoogleComputeSecurityPolicy#enable_each_unique_value}
	EnableEachUniqueValue interface{} `field:"optional" json:"enableEachUniqueValue" yaml:"enableEachUniqueValue"`
	// Requests that match this value constitute a granular traffic unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_security_policy#value GoogleComputeSecurityPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

