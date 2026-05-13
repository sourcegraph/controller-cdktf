package discoveryenginewidgetconfig


type DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMap struct {
	// Registered field name. The format is 'field.abc'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#field DiscoveryEngineWidgetConfig#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#ui_component DiscoveryEngineWidgetConfig#ui_component}.
	UiComponent *string `field:"required" json:"uiComponent" yaml:"uiComponent"`
	// Possible values: ["MOBILE", "DESKTOP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#device_visibility DiscoveryEngineWidgetConfig#device_visibility}
	DeviceVisibility *[]*string `field:"optional" json:"deviceVisibility" yaml:"deviceVisibility"`
	// The template to customize how the field is displayed.
	//
	// An example value would be a string that looks like: "Price: {value}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#display_template DiscoveryEngineWidgetConfig#display_template}
	DisplayTemplate *string `field:"optional" json:"displayTemplate" yaml:"displayTemplate"`
}

