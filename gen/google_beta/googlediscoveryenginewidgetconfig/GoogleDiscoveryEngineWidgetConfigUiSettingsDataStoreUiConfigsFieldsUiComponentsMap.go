package googlediscoveryenginewidgetconfig


type GoogleDiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMap struct {
	// Registered field name. The format is 'field.abc'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#field GoogleDiscoveryEngineWidgetConfig#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#ui_component GoogleDiscoveryEngineWidgetConfig#ui_component}.
	UiComponent *string `field:"required" json:"uiComponent" yaml:"uiComponent"`
	// Possible values: ["MOBILE", "DESKTOP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#device_visibility GoogleDiscoveryEngineWidgetConfig#device_visibility}
	DeviceVisibility *[]*string `field:"optional" json:"deviceVisibility" yaml:"deviceVisibility"`
	// The template to customize how the field is displayed.
	//
	// An example value would be a string that looks like: "Price: {value}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#display_template GoogleDiscoveryEngineWidgetConfig#display_template}
	DisplayTemplate *string `field:"optional" json:"displayTemplate" yaml:"displayTemplate"`
}

