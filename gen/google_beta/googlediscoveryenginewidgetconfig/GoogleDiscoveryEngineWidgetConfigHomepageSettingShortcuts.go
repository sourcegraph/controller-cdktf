package googlediscoveryenginewidgetconfig


type GoogleDiscoveryEngineWidgetConfigHomepageSettingShortcuts struct {
	// Destination URL of shortcut.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#destination_uri GoogleDiscoveryEngineWidgetConfig#destination_uri}
	DestinationUri *string `field:"optional" json:"destinationUri" yaml:"destinationUri"`
	// icon block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#icon GoogleDiscoveryEngineWidgetConfig#icon}
	Icon *GoogleDiscoveryEngineWidgetConfigHomepageSettingShortcutsIcon `field:"optional" json:"icon" yaml:"icon"`
	// Title of the shortcut.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#title GoogleDiscoveryEngineWidgetConfig#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

