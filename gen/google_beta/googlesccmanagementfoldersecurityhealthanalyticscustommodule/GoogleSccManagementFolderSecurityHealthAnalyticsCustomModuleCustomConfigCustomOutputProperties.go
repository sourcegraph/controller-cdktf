package googlesccmanagementfoldersecurityhealthanalyticscustommodule


type GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_folder_security_health_analytics_custom_module#name GoogleSccManagementFolderSecurityHealthAnalyticsCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_folder_security_health_analytics_custom_module#value_expression GoogleSccManagementFolderSecurityHealthAnalyticsCustomModule#value_expression}
	ValueExpression *GoogleSccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

