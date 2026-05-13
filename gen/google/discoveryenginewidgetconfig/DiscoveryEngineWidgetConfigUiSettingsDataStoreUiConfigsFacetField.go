package discoveryenginewidgetconfig


type DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetField struct {
	// Registered field name. The format is 'field.abc'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#field DiscoveryEngineWidgetConfig#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// The field name that end users will see.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#display_name DiscoveryEngineWidgetConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

