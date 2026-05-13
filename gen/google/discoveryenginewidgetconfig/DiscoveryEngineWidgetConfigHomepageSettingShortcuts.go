package discoveryenginewidgetconfig


type DiscoveryEngineWidgetConfigHomepageSettingShortcuts struct {
	// Destination URL of shortcut.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#destination_uri DiscoveryEngineWidgetConfig#destination_uri}
	DestinationUri *string `field:"optional" json:"destinationUri" yaml:"destinationUri"`
	// icon block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#icon DiscoveryEngineWidgetConfig#icon}
	Icon *DiscoveryEngineWidgetConfigHomepageSettingShortcutsIcon `field:"optional" json:"icon" yaml:"icon"`
	// Title of the shortcut.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_widget_config#title DiscoveryEngineWidgetConfig#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

