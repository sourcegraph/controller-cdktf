package googlesccmanagementprojectsecurityhealthanalyticscustommodule


type GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module#name GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module#value_expression GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule#value_expression}
	ValueExpression *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

