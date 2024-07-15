package googlesccorganizationcustommodule


type GoogleSccOrganizationCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_organization_custom_module#name GoogleSccOrganizationCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_organization_custom_module#value_expression GoogleSccOrganizationCustomModule#value_expression}
	ValueExpression *GoogleSccOrganizationCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

