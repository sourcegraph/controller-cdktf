package googleapigeeaddonsconfig


type GoogleApigeeAddonsConfigAddonsConfig struct {
	// advanced_api_ops_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_addons_config#advanced_api_ops_config GoogleApigeeAddonsConfig#advanced_api_ops_config}
	AdvancedApiOpsConfig *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig `field:"optional" json:"advancedApiOpsConfig" yaml:"advancedApiOpsConfig"`
	// api_security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_addons_config#api_security_config GoogleApigeeAddonsConfig#api_security_config}
	ApiSecurityConfig *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig `field:"optional" json:"apiSecurityConfig" yaml:"apiSecurityConfig"`
	// connectors_platform_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_addons_config#connectors_platform_config GoogleApigeeAddonsConfig#connectors_platform_config}
	ConnectorsPlatformConfig *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig `field:"optional" json:"connectorsPlatformConfig" yaml:"connectorsPlatformConfig"`
	// integration_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_addons_config#integration_config GoogleApigeeAddonsConfig#integration_config}
	IntegrationConfig *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig `field:"optional" json:"integrationConfig" yaml:"integrationConfig"`
	// monetization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_addons_config#monetization_config GoogleApigeeAddonsConfig#monetization_config}
	MonetizationConfig *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig `field:"optional" json:"monetizationConfig" yaml:"monetizationConfig"`
}

