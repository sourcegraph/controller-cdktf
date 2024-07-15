package googleprivatecacapool


type GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage struct {
	// base_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#base_key_usage GooglePrivatecaCaPool#base_key_usage}
	BaseKeyUsage *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage `field:"required" json:"baseKeyUsage" yaml:"baseKeyUsage"`
	// extended_key_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#extended_key_usage GooglePrivatecaCaPool#extended_key_usage}
	ExtendedKeyUsage *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage `field:"required" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// unknown_extended_key_usages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#unknown_extended_key_usages GooglePrivatecaCaPool#unknown_extended_key_usages}
	UnknownExtendedKeyUsages interface{} `field:"optional" json:"unknownExtendedKeyUsages" yaml:"unknownExtendedKeyUsages"`
}

