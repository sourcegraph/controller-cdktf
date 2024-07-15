package googlesccmanagementorganizationsecurityhealthanalyticscustommodule


type GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#name GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#value_expression GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#value_expression}
	ValueExpression *GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

